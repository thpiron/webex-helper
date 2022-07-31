package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// saveTokenCmd represents the saveToken command
var saveTokenCmd = &cobra.Command{
	Use:   "saveToken",
	Short: "Save your temporary token in a config file",
	Long: `This command allow you to save your temporary token in a config file.
This config file can be found at $HOME/.config/webex-helper.yml
`,
	Run: func(cmd *cobra.Command, args []string) {
		token := ""
		prompt := &survey.Password{
			Message: "Please paste your token here:",
		}
		survey.AskOne(prompt, &token)
		viper.Set("token", token)
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(saveTokenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveTokenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveTokenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
