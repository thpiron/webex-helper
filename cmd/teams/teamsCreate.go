package teams

import (
	"fmt"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
)

var (
	teamName string
)

func createTeam() error {
	wc := utils.NewWebexTeamsClient()
	team, resp, err := wc.Teams.CreateTeam(&webexteams.TeamCreateRequest{
		Name: teamName,
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

// teamsCreateCmd represents the teamsCreate command
var teamsCreateCmd = &cobra.Command{
	Use:     "teams",
	Short:   "Create a team",
	Long:    `Create a team`,
	Aliases: []string{"team"},
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("teamsFields", cmd.Flags().Lookup("teams-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := createTeam()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.CreateCmd.AddCommand(teamsCreateCmd)

	teamsCreateCmd.Flags().StringVar(&teamName, "name", "", "Name of the team")
	teamsCreateCmd.Flags().StringSliceVar(&teamsFields, "teams-fields", defaultTeamsFields, "Teams fields to display")

	teamsCreateCmd.MarkFlagRequired("name")
}
