package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	token      string
	jsonOutput bool
	rootCmd    = &cobra.Command{
		Use:   "webex-helper",
		Short: "Little CLI around the webex API to easily read API data",
		Long: `CLI wrapper around the webex API, for now only a few GET endpoints are supported
Use "webex-helper help" to get more information on the different commands.
		`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.SetDefault("token", "")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/webex-helper.yaml)")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Webex token, you can save it in a config file using saveToken")
	rootCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "", false, "Return a json formated output")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
}

func InitConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		configDir, err := os.UserConfigDir()
		cobra.CheckErr(err)
		// creating webex-helper folder if not exists
		_ = os.Mkdir(configDir+"/webex-helper", os.ModePerm)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(configDir + "/webex-helper/")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := viper.SafeWriteConfig(); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error when reading config file:", err)
		}
	}
}
