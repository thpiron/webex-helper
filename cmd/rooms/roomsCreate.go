package rooms

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
	webexteams "github.com/thpiron/webex-teams/sdk"
)

var (
	roomTitle string
	teamID    string
	// waiting for this PR to be merge on the cisco sdk https://github.com/jbogarin/go-cisco-webex-teams/pull/26
	// classificationID   string
	// isLocked           bool
	// isAnnouncementOnly bool
)

func CreateRoom() error {
	wc := utils.NewWebexTeamsClient()

	room, resp, err := wc.Rooms.CreateRoom(&webexteams.RoomCreateRequest{
		Title:  roomTitle,
		TeamID: teamID,
		// ClassificationId:   classificationID,
		// IsLocked:           isLocked,
		// IsAnnouncementOnly: isAnnouncementOnly,
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

	fields := viper.GetStringSlice("roomsFields")
	utils.PrintStructAsTable(*room, fields)
	return nil
}

// roomCreateCmd represents the roomsCreate command
var roomCreateCmd = &cobra.Command{
	Use:     "rooms",
	Short:   "Creates a new room",
	Long:    `Create a new room.`,
	Aliases: []string{"room"},
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("roomsFields", cmd.Flags().Lookup("rooms-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := CreateRoom()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.CreateCmd.AddCommand(roomCreateCmd)
	roomCreateCmd.Flags().StringVar(&roomTitle, "title", "", "Title of the room.")
	roomCreateCmd.Flags().StringVar(&teamID, "team-id", "", "The ID for the team with which this room is associated.")
	// createCmd.Flags().StringVar(&classificationID, "classification-id", "", "The classificationId for the room.")
	// createCmd.Flags().BoolVar(&isLocked, "is-locked", false, "Set the space as locked/moderated and the creator becomes a moderator")
	// createCmd.Flags().BoolVar(&isAnnouncementOnly, "is-announcement-only", false, "Sets the space into Announcement Mode.")
	roomCreateCmd.MarkFlagRequired("title")
	roomCreateCmd.Flags().StringSlice("rooms-fields", defaultRoomsFields, "Rooms fields to display")
}
