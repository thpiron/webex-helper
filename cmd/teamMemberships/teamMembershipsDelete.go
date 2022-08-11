package teamMemberships

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
)

func deleteTeamMembership(teamMembershipID string) error {
	wc := utils.NewWebexTeamsClient()

	resp, err := wc.TeamMemberships.DeleteTeamMembership(teamMembershipID)

	if err != nil {
		return err
	}
	if err := utils.CheckWebexError(*resp); err != nil {
		return err
	}
	fmt.Println("TeamMembership deleted.")
	return nil
}

// teamMembershipsDeleteCmd represents the teamMembershipsDelete command
var teamMembershipsDeleteCmd = &cobra.Command{
	Use:     "teamMemberships",
	Short:   "Delete a teamMemberships",
	Long:    `Delete a teamMemberships`,
	Aliases: []string{"teamMembership"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := deleteTeamMembership(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.DeleteCmd.AddCommand(teamMembershipsDeleteCmd)
}
