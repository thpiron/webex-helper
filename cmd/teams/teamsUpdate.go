package teams

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
	webexteams "github.com/thpiron/webex-teams/sdk"
)

var (
	updatedTeamName string
)

func updateTeam(teamID string) error {
	wc := utils.NewWebexTeamsClient()

	team, resp, err := wc.Teams.UpdateTeam(teamID, &webexteams.TeamUpdateRequest{
		Name: updatedTeamName,
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
	utils.PrintStructAsTable(*team, fields)
	return nil
}

// teamsUpdateCmd represents the teamsUpdate command
var teamsUpdateCmd = &cobra.Command{
	Use:     "teams",
	Short:   "Update a team",
	Long:    `Update a team.`,
	Aliases: []string{"team"},
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("teamsFields", cmd.Flags().Lookup("teams-fields"))
	},
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := updateTeam(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.UpdateCmd.AddCommand(teamsUpdateCmd)

	teamsUpdateCmd.Flags().StringVar(&updatedTeamName, "name", "", "Name of the team")
	teamsUpdateCmd.Flags().StringSliceVar(&teamsFields, "teams-fields", defaultTeamsFields, "Teams fields to display")

	teamsUpdateCmd.MarkFlagRequired("name")
}
