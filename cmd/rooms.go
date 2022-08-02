package cmd

import (
	"fmt"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	roomsMax           int
	roomsFields        []string
	defaultRoomsFields = []string{"Title", "RoomType", "IsLocked", "Created"}
)

func listRooms(max int) error {
	wc := NewWebexTeamsClient()
	queryParams := &webexteams.ListRoomsQueryParams{
		Max: max,
	}
	rooms, resp, err := wc.Rooms.ListRooms(queryParams)

	if err != nil {
		return err
	}

	if err := checkWebexError(*resp); err != nil {
		return err
	}

	if jsonOutput {
		fmt.Println(string(resp.Body()))
		return nil
	}
	s := make([]interface{}, len(rooms.Items))
	for i, v := range rooms.Items {
		s[i] = v
	}

	fields := viper.GetStringSlice("roomsFields")
	printStructSliceAsTable(s, fields)
	return nil
}

// roomsCmd represents the rooms command
var roomsCmd = &cobra.Command{
	Use:   "rooms",
	Short: "Retrieve information on rooms",
	Long: `
Rooms command let you list and get details on rooms. You can only list rooms you're into.
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
	Run: func(cmd *cobra.Command, args []string) {
		err := listRooms(roomsMax)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(roomsCmd)
	roomsCmd.Flags().IntVarP(
		&roomsMax,
		"max",
		"m",
		20,
		"Number max of rooms to list",
	)
	roomsCmd.Flags().StringSliceVar(&roomsFields, "rooms-fields", defaultRoomsFields, "Rooms fields to display")
	viper.BindPFlag("roomsFields", roomsCmd.Flags().Lookup("rooms-fields"))
}
