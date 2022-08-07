/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/utils"
)

func GetRoom(roomID string) error {
	wc := utils.NewWebexTeamsClient()
	room, resp, err := wc.Rooms.GetRoom(roomID)
	if err != nil {
		return err
	}
	if err := utils.CheckWebexError(*resp); err != nil {
		return err
	}

	if jsonOutput {
		fmt.Println(string(resp.Body()))
		return nil
	}

	fields := viper.GetStringSlice("roomsFields")
	fmt.Println(fields)
	utils.PrintStructAsTable(*room, fields)
	return nil
}

// roomsGetCmd represents the peopleGet command
var roomsGetCmd = &cobra.Command{
	Use:   "rooms",
	Short: "Get a room details",
	Long: `Retrieve a room details from its id
You can set the fields to display in table mode in your config file ($HOME/.config/webex-helper/config.yaml):
roomsFields: 
	- ID
	- Title
	- RoomType
	- IsLocked
	- TeamID
	- CreatorID
	- LastActivity
	- Created
`,
	Args:    cobra.ExactArgs(1),
	Example: "webex-helper get people [flags] <peopleID>",
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("roomsFields", cmd.Flags().Lookup("rooms-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := GetRoom(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	getCmd.AddCommand(roomsGetCmd)

	roomsGetCmd.Flags().StringSliceVar(&roomsFields, "rooms-fields", defaultRoomsFields, "Rooms fields to display")
}
