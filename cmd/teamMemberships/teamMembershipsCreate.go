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
	teamID      string
	personID    string
	personEmail string
	isModerator bool
)

func createTeamMembership() error {
	wc := utils.NewWebexTeamsClient()
	teamMembership, resp, err := wc.TeamMemberships.CreateTeamMembership(&webexteams.TeamMembershipCreateRequest{
		TeamID:      teamID,
		PersonID:    personID,
		PersonEmail: personEmail,
		IsModerator: isModerator,
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

	fields := viper.GetStringSlice("teamMembershipsFields")
	utils.PrintStructAsTable(*teamMembership, fields)
	return nil
}

// teamMembershipsCreateCmd represents the teamMembershipsCreate command
var teamMembershipsCreateCmd = &cobra.Command{
	Use:     "teamMemberships",
	Short:   "Create a teamMembership",
	Long:    "Create a teamMembership",
	Aliases: []string{"teamMembership"},
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("teamMembershipsFields", cmd.Flags().Lookup("teamMemberships-fields"))
		if personEmail == "" {
			cmd.MarkFlagRequired("person-id")
		}
		if personID == "" {
			cmd.MarkFlagRequired("person-email")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := createTeamMembership()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.CreateCmd.AddCommand(teamMembershipsCreateCmd)

	teamMembershipsCreateCmd.Flags().StringVar(&teamID, "team-id", "", "Team ID of the membership")
	teamMembershipsCreateCmd.Flags().StringVar(&personEmail, "person-email", "", "Email of the membership's user")
	teamMembershipsCreateCmd.Flags().StringVar(&personID, "person-id", "", "ID of the membership's user")
	teamMembershipsCreateCmd.Flags().BoolVar(&isModerator, "is-moderator", false, "make the user a team's moderator")

	teamMembershipsCreateCmd.MarkFlagsMutuallyExclusive("person-id", "person-email")

	teamMembershipsCreateCmd.Flags().StringSliceVar(&teamMembershipsFields, "teamMemberships-fields", defaultTeamMembershipsFields, "Teams memberships fields to display")

}
