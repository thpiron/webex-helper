package messages

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
)

var (
	direct          bool
	mentionedPeople string
	// before related is commented for now, waiting for the sdk to be fixed
	// PR: https://github.com/jbogarin/go-cisco-webex-teams/pull/27
	// before          time.Time
	// beforeRaw     string
	beforeMessage string
)

func listMessages() error {
	wc := utils.NewWebexTeamsClient()

	var (
		messages *webexteams.Messages
		resp     *resty.Response
		err      error
	)
	if direct {
		messages, resp, err = wc.Messages.GetDirectMessages(&webexteams.DirectMessagesQueryParams{
			ParentID:    parentID,
			PersonID:    personID,
			PersonEmail: personEmail,
		})
	} else {
		messages, resp, err = wc.Messages.ListMessages(&webexteams.ListMessagesQueryParams{
			RoomID: roomID,
			// ParentID:        parentID, Not yet implemented in SDK
			// Before:          before,
			BeforeMessage:   beforeMessage,
			MentionedPeople: mentionedPeople,
		})
	}

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
	s := make([]interface{}, len(messages.Items))
	for i, v := range messages.Items {
		s[i] = v
	}

	fields := viper.GetStringSlice("messagesFields")
	utils.PrintStructSliceAsTable(s, fields)
	return nil
}

// messagesListCmd represents the messagesList command
var messagesListCmd = &cobra.Command{
	Use:   "messages",
	Short: "List messages",
	Long:  `List messages.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		viper.BindPFlag("messagesFields", cmd.Flags().Lookup("messages-fields"))
		// if beforeRaw != "" {
		// 	var err error
		// 	before, err = time.Parse("2006-01-02 15:04:05", beforeRaw)
		// 	if err != nil {
		// 		return errors.New("error when parsing 'before' date, please use the 'yyyy-mm-dd HH:MM:SS' format")
		// 	}
		// }
		if !direct {
			cmd.MarkFlagRequired("room-id")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := listMessages()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.ListCmd.AddCommand(messagesListCmd)
	messagesListCmd.Flags().BoolVar(&direct, "direct", false, "List only direct messages. Allows you to not set room-id")
	messagesListCmd.Flags().StringVar(&roomID, "room-id", "", "Room's ID of the messages to list. Required when direct flags not set.")
	messagesListCmd.Flags().StringVar(&parentID, "parent-id", "", "ID of the parent message of the messages to list (thread)")
	messagesListCmd.Flags().StringVar(&personID, "person-id", "", "Person ID of the creator of the messages to list")
	messagesListCmd.Flags().StringVar(&personEmail, "person-email", "", "Person email of the creator of the messages to list.")
	messagesListCmd.Flags().StringVar(&beforeMessage, "before-message", "", "Message ID of the last message to list (Only for not direct messages).")
	messagesListCmd.Flags().StringVar(&mentionedPeople, "mentioned-people", "", "List messages where this people is mentionned") // For now sdk don't accept a list of mentionned people
	// messagesListCmd.Flags().StringVar(&beforeRaw, "before", "", "List message before the given date (format: yyyy-mm-dd HH:MM:SS). This must be UTC time")

	messagesListCmd.Flags().StringSliceVar(&messagesFields, "messages-fields", defaultMessagesFields, "Memberships fields to display")
}
