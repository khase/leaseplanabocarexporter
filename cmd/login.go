package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/khase/leaseplanabocarexporter/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	email    string
	password string

	loginCmd = &cobra.Command{
		Use:   "login",
		Short: "login to the leaseplan abocar platform and safe the token to the current config",
		Long:  `login to the leaseplan abocar platform and safe the token to the current config`,
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			if email == "" {
				fmt.Print("E-Mail: ")
				tmp, _ := reader.ReadString('\n')
				email = strings.TrimSuffix(tmp, "\n")
			} else {
				fmt.Println("execute login for " + email)
			}

			if password == "" {
				fmt.Print("Password: ")
				bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
				password = string(bytePassword)
				fmt.Println("")
			}

			token, err := pkg.GetToken(email, password)

			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			log.Println("Token retrieved: " + token)

			viper.Set("AddressToken", token)
			viper.WriteConfig()

			log.Println("Token saved to: " + viper.ConfigFileUsed())
		},
	}
)

func init() {
	loginCmd.PersistentFlags().StringVar(&email, "email", "", "email to be used for login")
	loginCmd.PersistentFlags().StringVar(&password, "password", "", "password to be used for login")
}
