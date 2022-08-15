package messages

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
)

func getMessage(messageID string) error {
	wc := utils.NewWebexTeamsClient()

	message, resp, err := wc.Messages.GetMessage(messageID)

	if err != nil {
		return err
	}

	if err := utils.CheckWebexError(*resp); err != nil {
		return err
	}

	if cmd.JsonOutput {
		fmt.Println(string(resp.Body()))
		return nil
	}
	fields := viper.GetStringSlice("messagesFields")
	utils.PrintStructAsTable(*message, fields)
	return nil
}

// messagesGetCmd represents the messagesGet command
var messagesGetCmd = &cobra.Command{
	Use:     "messages",
	Short:   "Get a message detail",
	Long:    `Get a message detail.`,
	Aliases: []string{"message"},
	Args:    cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("messagesFields", cmd.Flags().Lookup("messages-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := getMessage(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.GetCmd.AddCommand(messagesGetCmd)
	messagesGetCmd.Flags().StringSliceVar(&messagesFields, "messages-fields", defaultMessagesFields, "Memberships fields to display")
}
