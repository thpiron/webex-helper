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
}
