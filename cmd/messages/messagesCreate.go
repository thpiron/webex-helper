package messages

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
	webexteams "github.com/thpiron/webex-teams/sdk"
)

func createMessage() error {
	wc := utils.NewWebexTeamsClient()

	message, resp, err := wc.Messages.CreateMessage(&webexteams.MessageCreateRequest{
		RoomID:        roomID,
		ParentID:      parentID,
		ToPersonID:    personID,
		ToPersonEmail: personEmail,
		Text:          text,
		Markdown:      markdown,
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

// messagesCreateCmd represents the messagesCreate command
var messagesCreateCmd = &cobra.Command{
	Use:     "messages",
	Short:   "Create a message",
	Long:    `Create a message.`,
	Aliases: []string{"message"},
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("messagesFields", cmd.Flags().Lookup("messages-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := createMessage()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.CreateCmd.AddCommand(messagesCreateCmd)

	messagesCreateCmd.Flags().StringVar(&roomID, "room-id", "", "Room ID of the room where to send message")
	messagesCreateCmd.Flags().StringVar(&parentID, "parent-id", "", "ID of the message to respond (thread)")
	messagesCreateCmd.Flags().StringVar(&personID, "person-id", "", "ID of the person to send the message (direct message)")
	messagesCreateCmd.Flags().StringVar(&personEmail, "person-email", "", "Eamil of the person to send the message (direct message)")
	messagesCreateCmd.Flags().StringVar(&text, "text", "", "Unformated text of the message")
	messagesCreateCmd.Flags().StringVar(&markdown, "markdown", "", "Markdown formated text of the message")

	messagesCreateCmd.Flags().StringSliceVar(&messagesFields, "messages-fields", defaultMessagesFields, "Memberships fields to display")

	messagesCreateCmd.MarkFlagsMutuallyExclusive("room-id", "person-email", "person-id")
}
