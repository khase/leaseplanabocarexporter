package cmd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/khase/leaseplanabocarexporter/dto"
	"github.com/khase/leaseplanabocarexporter/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	page                   int
	count                  int
	all                    bool
	output                 string
	pollDuration           int
	executeOnChangeCommand string

	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get list of cars",
		Long:  `get list of cars`,
		Run: func(cmd *cobra.Command, args []string) {
			token := viper.GetString("AddressToken")

			if token == "" {
				log.Fatalln("no token present, use 'login' to log into the service.")
				os.Exit(1)
			}

			for {
				var cars []dto.Item
				var err error

				if all {
					cars, err = pkg.GetAllCars(token, page, count)
				} else {
					cars, err = pkg.GetPageItems(token, page, count)
				}

				log.Println("Loading car list...")

				if err != nil {
					log.Fatalln(err)
					os.Exit(1)
				}

				log.Println("got " + strconv.Itoa(len(cars)) + " cars")

				if output != "" {
					changed, err := dumpToFile(output, cars)
					if err != nil {
						log.Fatalln(err)
						os.Exit(1)
					}
					if changed && executeOnChangeCommand != "" {
						log.Println("Output changed, executing post command")
						cmd := exec.Command(executeOnChangeCommand)
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr

						err := cmd.Run()
						if err != nil {
							log.Fatalln(err)
							os.Exit(1)
						}
					}
				} else {
					for index, car := range cars {
						log.Println(strconv.Itoa(index+1) + ") \t" + strconv.Itoa(car.VehiclesCount) + "x \t" + car.OfferTypeName + " - " + strconv.Itoa(int(car.RentalObject.PowerHP)) + "HP (" + strconv.Itoa(int(car.RentalObject.PriceProducer1)) + "€)")
					}
				}

				const sleepReminderInMinutes = 5
				if pollDuration > 0 {
					for i := 0; i < pollDuration; i += sleepReminderInMinutes {
						log.Println("Sleeping for " + strconv.Itoa(pollDuration-i) + "min...")
						time.Sleep(sleepReminderInMinutes * time.Minute)
					}
				} else {
					break
				}
			}
		},
	}
)

func init() {
	getCmd.PersistentFlags().IntVarP(&page, "page", "p", 1, "page to get.")
	getCmd.PersistentFlags().IntVarP(&count, "count", "n", 10, "number of cars to get for a single page.")
	getCmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "indicates if the programm should page through all cars.")
	getCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "filename where the json should be written to.")
	getCmd.PersistentFlags().StringVar(&executeOnChangeCommand, "exec", "", "command to execute when output file changed (requires output filename to be set)")
	getCmd.PersistentFlags().IntVar(&pollDuration, "poll", 0, "polling cycle in minutes.")
}

func dumpToFile(name string, data interface{}) (cahnged bool, err error) {
	byteData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return false, err
	}
	stringData := string(byteData)

	if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
		dir := filepath.Dir(name)
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return false, err
		}

		err = writeFile(name, stringData)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	currentContent, err := readFile(name)
	if err != nil {
		return false, err
	}

	if currentContent == stringData {
		return false, nil
	}

	err = writeFile(name, stringData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func writeFile(name string, content string) error {
	f, err := os.Create(name)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)

	if err != nil {
		return err
	}

	return nil
}

func readFile(name string) (string, error) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
