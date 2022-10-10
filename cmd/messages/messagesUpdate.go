package messages

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
	webexteams "github.com/thpiron/webex-teams/sdk"
)

func updateMessage(messageID string) error {
	wc := utils.NewWebexTeamsClient()

	message, resp, err := wc.Messages.EditMessage(messageID, &webexteams.MessageEditRequest{
		RoomID:   roomID,
		Text:     text,
		Markdown: markdown,
	})

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

// messagesUpdateCmd represents the messagesUpdate command
var messagesUpdateCmd = &cobra.Command{
	Use:     "messages",
	Short:   "Update a message",
	Long:    `Update a message.`,
	Aliases: []string{"message"},
	Args:    cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("messagesFields", cmd.Flags().Lookup("messages-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := updateMessage(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.UpdateCmd.AddCommand(messagesUpdateCmd)
	messagesUpdateCmd.Flags().StringVar(&roomID, "room-id", "", "Room ID of the room where to send message")
	messagesUpdateCmd.Flags().StringVar(&text, "text", "", "Unformated text of the message")
	messagesUpdateCmd.Flags().StringVar(&markdown, "markdown", "", "Markdown formated text of the message")

	messagesUpdateCmd.Flags().StringSliceVar(&messagesFields, "messages-fields", defaultMessagesFields, "Memberships fields to display")

	messagesUpdateCmd.MarkFlagRequired("room-id")
}
