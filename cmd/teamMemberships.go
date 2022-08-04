package cmd

import (
	"fmt"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	teamMembershipsTeamID        string
	teamMembershipsMax           int
	teamMembershipsFields        []string
	defaultTeamMembershipsFields = []string{"PersonEmail", "PersonDisplayName", "IsModerator"}
)

func listTeamMemberships(max int) error {
	wc := NewWebexTeamsClient()
	queryParams := &webexteams.ListTeamMemberhipsQueryParams{
		TeamID: teamMembershipsTeamID,
		Max:    max,
	}
	teamMemberships, resp, err := wc.TeamMemberships.ListTeamMemberhips(queryParams)

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
	s := make([]interface{}, len(teamMemberships.Items))
	for i, v := range teamMemberships.Items {
		s[i] = v
	}

	fields := viper.GetStringSlice("teamMembershipsFields")
	printStructSliceAsTable(s, fields)
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
	Run: func(cmd *cobra.Command, args []string) {
		err := listTeamMemberships(teamMembershipsMax)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(teamMembershipsCmd)
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
	viper.BindPFlag("teamMembershipsFields", teamMembershipsCmd.Flags().Lookup("teamMemberships-fields"))
}
