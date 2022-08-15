package memberships

import (
	"fmt"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
)

var (
	max int
)

func listTeamMemberships(max int) error {
	wc := utils.NewWebexTeamsClient()
	queryParams := &webexteams.ListMembershipsQueryParams{
		RoomID:      roomID,
		PersonID:    personID,
		PersonEmail: personEmail,
		Max:         max,
	}
	teamMemberships, resp, err := wc.Memberships.ListMemberships(queryParams)

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
	s := make([]interface{}, len(teamMemberships.Items))
	for i, v := range teamMemberships.Items {
		s[i] = v
	}

	fields := viper.GetStringSlice("membershipsFields")
	utils.PrintStructSliceAsTable(s, fields)
	return nil
}

// membershipsListCmd represents the membershipsList command
var membershipsListCmd = &cobra.Command{
	Use:   "memberships",
	Short: "List memberships",
	Long:  `List memberships.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("membershipsFields", cmd.Flags().Lookup("memberships-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := listTeamMemberships(max)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.ListCmd.AddCommand(membershipsListCmd)
	membershipsListCmd.Flags().StringVar(&roomID, "room-id", "", "Filter memberships by room ID")
	membershipsListCmd.Flags().StringVar(&personID, "person-id", "", "Filter memberships by person ID")
	membershipsListCmd.Flags().StringVar(&personEmail, "person-email", "", "Filter memberships by person email")
	membershipsListCmd.Flags().IntVarP(
		&max,
		"max",
		"m",
		20,
		"Number max of teams to list",
	)
	membershipsListCmd.Flags().StringSliceVar(&membershipsFields, "memberships-fields", defaultMembershipsFields, "Memberships fields to display")

}
