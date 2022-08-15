package memberships

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
)

func deleteMembership(membershipID string) error {
	wc := utils.NewWebexTeamsClient()

	resp, err := wc.Memberships.DeleteMembership(membershipID)

	if err != nil {
		return err
	}
	if err := utils.CheckWebexError(*resp); err != nil {
		return err
	}
	fmt.Println("Membership deleted.")
	return nil
}

// membershipsDeleteCmd represents the membershipsDelete command
var membershipsDeleteCmd = &cobra.Command{
	Use:     "memberships",
	Short:   "Delete a membership",
	Long:    `Delete a membership.`,
	Aliases: []string{"membership"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := deleteMembership(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.DeleteCmd.AddCommand(membershipsDeleteCmd)
}
