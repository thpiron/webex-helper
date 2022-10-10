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
	updateRoomTitle string
	// waiting for this PR to be merge on the cisco sdk https://github.com/jbogarin/go-cisco-webex-teams/pull/26
	// updateClassificationID   string
	// updateIsLocked           bool
	// updateIsAnnouncementOnly bool
)

func UpdateRoom(roomID string) error {
	wc := utils.NewWebexTeamsClient()

	room, resp, err := wc.Rooms.UpdateRoom(roomID, &webexteams.RoomUpdateRequest{
		Title: updateRoomTitle,
		// ClassificationId:   updateClassificationID,
		// IsLocked:           updateIsLocked,
		// IsAnnouncementOnly: updateIsAnnouncementOnly,
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

// roomsCreateCmd represents the roomsCreate command
var roomsUpdateCmd = &cobra.Command{
	Use:     "rooms",
	Short:   "Update a new room",
	Long:    `Update a new room.`,
	Aliases: []string{"room"},
	Args:    cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("roomsFields", cmd.Flags().Lookup("rooms-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := UpdateRoom(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.UpdateCmd.AddCommand(roomsUpdateCmd)
	roomsUpdateCmd.Flags().StringVar(&updateRoomTitle, "title", "", "Title of the room.")
	// roomsUpdateCmd.Flags().StringVar(&teamID, "team-id", "", "The ID for the team with which this room is associated.")
	// roomsUpdateCmd.Flags().StringVar(&classificationID, "classification-id", "", "The classificationId for the room.")
	// roomsUpdateCmd.Flags().BoolVar(&isLocked, "is-locked", false, "Set the space as locked/moderated and the creator becomes a moderator")
	// roomsUpdateCmd.Flags().BoolVar(&isAnnouncementOnly, "is-announcement-only", false, "Sets the space into Announcement Mode.")
	roomsUpdateCmd.MarkFlagRequired("title")
	roomsUpdateCmd.Flags().StringSliceVar(&roomsFields, "rooms-fields", defaultRoomsFields, "Rooms fields to display")
}
