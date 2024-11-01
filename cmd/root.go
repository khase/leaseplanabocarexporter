package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
	token   string

	rootCmd = &cobra.Command{
		Use:   "LeaseplanAbocarExporter",
		Short: "LeaseplanAbocarExporter is a go tool for exporting information from the leaseplan abocar platform",
		Long:  `LeaseplanAbocarExporter is a go tool for exporting information from the leaseplan abocar platform`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.leaseplanabocar.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "token to be used for auth (can be retrieved using LeaseplanAbocarExporter login)")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.BindPFlag("AddressToken", rootCmd.PersistentFlags().Lookup("token"))

	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(userInfoCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".leaseplanabocar".
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".leaseplanabocar")

		viper.SafeWriteConfigAs(home + "/.leaseplanabocar")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}
