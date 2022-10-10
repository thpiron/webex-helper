package teamMemberships

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
	webexteams "github.com/thpiron/webex-teams/sdk"
)

var (
	updatedIsModerator bool
)

func updateTeamMembership(teamMembershipID string) error {
	wc := utils.NewWebexTeamsClient()
	teamMembership, resp, err := wc.TeamMemberships.UpdateTeamMembership(teamMembershipID, &webexteams.TeamMembershipUpdateRequest{
		IsModerator: updatedIsModerator,
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

	fields := viper.GetStringSlice("teamsFields")
	utils.PrintStructAsTable(*teamMembership, fields)
	return nil
}

// teamMembershipsUpdateCmd represents the teamMembershipsUpdate command
var teamMembershipsUpdateCmd = &cobra.Command{
	Use:     "teamMemberships",
	Short:   "Update a teamMembership",
	Long:    `Update a teamMembership`,
	Aliases: []string{"teamMembership"},
	Args:    cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("teamMembershipsFields", cmd.Flags().Lookup("teamMemberships-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := updateTeamMembership(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.UpdateCmd.AddCommand(teamMembershipsUpdateCmd)

	teamMembershipsUpdateCmd.Flags().BoolVar(&updatedIsModerator, "is-moderator", false, "make the user a team's moderator")
	teamMembershipsUpdateCmd.Flags().StringSliceVar(&teamMembershipsFields, "teamMemberships-fields", defaultTeamMembershipsFields, "Teams memberships fields to display")
}
