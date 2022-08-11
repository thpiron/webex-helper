package teamMemberships

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
)

func getTeamMembership(membershipId string) error {
	wc := utils.NewWebexTeamsClient()

	teamMembership, resp, err := wc.TeamMemberships.GetTeamMembership(membershipId)

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

// teamMembershipsGetCmd represents the teamMembershipsGet command
var teamMembershipsGetCmd = &cobra.Command{
	Use:     "teamMemberships",
	Short:   "Get teamMemberships details",
	Long:    `Get teamMemberships details`,
	Aliases: []string{"teamMembership"},
	Args:    cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("teamMembershipsFields", cmd.Flags().Lookup("teamMemberships-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := getTeamMembership(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.GetCmd.AddCommand(teamMembershipsGetCmd)
	teamMembershipsGetCmd.Flags().StringSliceVar(&teamMembershipsFields, "teamMemberships-fields", defaultTeamMembershipsFields, "Teams memberships fields to display")
}
