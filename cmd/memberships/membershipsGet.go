package memberships

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
)

func getMembership(membershipId string) error {
	wc := utils.NewWebexTeamsClient()

	membership, resp, err := wc.Memberships.GetMembership(membershipId)

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

	fields := viper.GetStringSlice("membershipsFields")
	utils.PrintStructAsTable(*membership, fields)
	return nil
}

// membershipsGetCmd represents the membershipsGet command
var membershipsGetCmd = &cobra.Command{
	Use:     "memberships",
	Short:   "Get memberships details",
	Long:    `Get memberships details.`,
	Aliases: []string{"membership"},
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("membershipsFields", cmd.Flags().Lookup("memberships-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := getMembership(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.GetCmd.AddCommand(membershipsGetCmd)
	membershipsGetCmd.Flags().StringSliceVar(&membershipsFields, "memberships-fields", defaultMembershipsFields, "Memberships fields to display")
}
