package rooms

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
)

func roomsDelete(roomID string) error {
	wc := utils.NewWebexTeamsClient()

	resp, err := wc.Rooms.DeleteRoom(roomID)

	if err != nil {
		return err
	}
	if err := utils.CheckWebexError(*resp); err != nil {
		return err
	}
	fmt.Println("Room deleted.")
	return nil
}

// roomsDeleteCmd represents the roomsDelete command
var roomsDeleteCmd = &cobra.Command{
	Use:     "rooms",
	Short:   "Delete a room given its id",
	Long:    `Delete a room `,
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"room"},
	Run: func(cmd *cobra.Command, args []string) {
		err := roomsDelete(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.DeleteCmd.AddCommand(roomsDeleteCmd)
}
