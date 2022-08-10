package teamMemberships

import (
	"fmt"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
)

var (
	teamMembershipsTeamID        string
	teamMembershipsMax           int
	teamMembershipsFields        []string
	defaultTeamMembershipsFields = []string{"PersonEmail", "PersonDisplayName", "IsModerator"}
)

func listTeamMemberships(max int) error {
	wc := utils.NewWebexTeamsClient()
	queryParams := &webexteams.ListTeamMemberhipsQueryParams{
		TeamID: teamMembershipsTeamID,
		Max:    max,
	}
	teamMemberships, resp, err := wc.TeamMemberships.ListTeamMemberhips(queryParams)

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

	fields := viper.GetStringSlice("teamMembershipsFields")
	utils.PrintStructSliceAsTable(s, fields)
	return nil
}

// teamsCmd represents the teams command
var teamMembershipsCmd = &cobra.Command{
	Use:   "teamMemberships",
	Short: "Retrieve information on a team's memberships",
	Long: `
Teams command let you list and get details on teams memberships.
You need to enter a team id.
You can set the fields to display in table mode in your config file ($HOME/.config/webex-helper/config.yaml):
teamsFields:
	- ID
	- TeamID
	- PersonID
	- PersonEmail
	- PersonDisplayName
	- IsModerator
	- Created
`,
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("teamMembershipsFields", cmd.Flags().Lookup("teamMemberships-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := listTeamMemberships(teamMembershipsMax)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.ListCmd.AddCommand(teamMembershipsCmd)
	teamMembershipsCmd.Flags().StringVar(
		&teamMembershipsTeamID,
		"team-id",
		"",
		"ID of the team",
	)
	teamMembershipsCmd.MarkFlagRequired("team-id")
	teamMembershipsCmd.Flags().IntVarP(
		&teamMembershipsMax,
		"max",
		"m",
		20,
		"Number max of teams to list",
	)
	teamMembershipsCmd.Flags().StringSliceVar(&teamMembershipsFields, "teamMemberships-fields", defaultTeamMembershipsFields, "Teams memberships fields to display")
}
