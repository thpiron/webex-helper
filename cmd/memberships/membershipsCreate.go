package memberships

import (
	"fmt"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
)

func createMembership() error {
	wc := utils.NewWebexTeamsClient()
	membership, resp, err := wc.Memberships.CreateMembership(&webexteams.MembershipCreateRequest{
		RoomID:      roomID,
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

	fields := viper.GetStringSlice("membershipsFields")
	utils.PrintStructAsTable(*membership, fields)
	return nil
}

// membershipsCreateCmd represents the membershipsCreate command
var membershipsCreateCmd = &cobra.Command{
	Use:     "memberships",
	Short:   "Create a membership",
	Long:    `Create a membership.`,
	Aliases: []string{"membership"},
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("membershipsFields", cmd.Flags().Lookup("memberships-fields"))
		if personEmail == "" {
			cmd.MarkFlagRequired("person-id")
		}
		if personID == "" {
			cmd.MarkFlagRequired("person-email")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := createMembership()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.CreateCmd.AddCommand(membershipsCreateCmd)
	membershipsCreateCmd.Flags().StringVar(&roomID, "room-id", "", "Room ID of the membership")
	membershipsCreateCmd.Flags().StringVar(&personEmail, "person-email", "", "Email of the membership's user")
	membershipsCreateCmd.Flags().StringVar(&personID, "person-id", "", "ID of the membership's user")
	membershipsCreateCmd.Flags().BoolVar(&isModerator, "is-moderator", false, "make the user a room's moderator")

	membershipsCreateCmd.MarkFlagsMutuallyExclusive("person-id", "person-email")

	membershipsCreateCmd.Flags().StringSliceVar(&membershipsFields, "memberships-fields", defaultMembershipsFields, "Memberships fields to display")

}
