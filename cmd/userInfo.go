package cmd

import (
	"log"
	"os"

	"github.com/khase/leaseplanabocarexporter/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	userInfoCmd = &cobra.Command{
		Use:   "userInfo",
		Short: "gets the userInfo fot the provided user",
		Long:  `gets the userInfo fot the provided user`,
		Run: func(cmd *cobra.Command, args []string) {
			token := viper.GetString("AddressToken")

			if token == "" {
				log.Fatalln("no token present, use 'login' to log into the service.")
				os.Exit(1)
			}

			userInfo, err := pkg.GetUserInfo(token)

			if err != nil {
				log.Fatalln(err)
				os.Exit(1)
			}

			if output != "" {
				_, err := dumpToFile(output, userInfo)
				if err != nil {
					log.Fatalln(err)
					os.Exit(1)
				}
			} else {
				log.Println("Assigned role for " + userInfo.Firstname + " " + userInfo.Name1 + " is " + userInfo.AddressRole.RoleName)
			}
		},
	}
)

func init() {
	userInfoCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "filename where the json should be written to.")
}
