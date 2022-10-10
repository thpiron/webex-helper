package rooms

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
	webexteams "github.com/thpiron/webex-teams/sdk"
)

const lastActivity string = "lastactivity"

var (
	roomsMax           int
	roomsFields        []string
	sortBy             string
	defaultRoomsFields = []string{"Title", "RoomType", "IsLocked", "Created"}
)

func listRooms(max int) error {
	wc := utils.NewWebexTeamsClient()

	queryParams := &webexteams.ListRoomsQueryParams{
		Max:    max,
		SortBy: sortBy,
	}
	rooms, resp, err := wc.Rooms.ListRooms(queryParams)

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
	s := make([]interface{}, len(rooms.Items))
	for i, v := range rooms.Items {
		s[i] = v
	}

	fields := viper.GetStringSlice("roomsFields")
	utils.PrintStructSliceAsTable(s, fields)
	return nil
}

// roomsListCmd represents the rooms command
var roomsListCmd = &cobra.Command{
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
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("roomsFields", cmd.Flags().Lookup("rooms-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := listRooms(roomsMax)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.ListCmd.AddCommand(roomsListCmd)
	roomsListCmd.Flags().IntVarP(
		&roomsMax,
		"max",
		"m",
		20,
		"Number max of rooms to list",
	)
	roomsListCmd.Flags().StringVarP(&sortBy, "sort-by", "s", lastActivity, "Choose how to sort rooms (id, lastactivity, created)")
	roomsListCmd.Flags().StringSliceVar(&roomsFields, "rooms-fields", defaultRoomsFields, "Rooms fields to display")
}
