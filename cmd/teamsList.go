package cmd

import (
	"fmt"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/utils"
)

var (
	teamsMax           int
	teamsFields        []string
	defaultTeamsFields = []string{"ID", "Name", "CreatorID", "Created"}
)

func listTeams(max int) error {
	wc := utils.NewWebexTeamsClient()
	queryParams := &webexteams.ListTeamsQueryParams{
		Max: max,
	}
	teams, resp, err := wc.Teams.ListTeams(queryParams)

	if err != nil {
		return err
	}

	if err := utils.CheckWebexError(*resp); err != nil {
		return err
	}

	if jsonOutput {
		fmt.Println(string(resp.Body()))
		return nil
	}
	s := make([]interface{}, len(teams.Items))
	for i, v := range teams.Items {
		s[i] = v
	}

	fields := viper.GetStringSlice("teamsFields")
	utils.PrintStructSliceAsTable(s, fields)
	return nil
}

// teamsCmd represents the teams command
var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Retrieve information on teams",
	Long: `
Teams command let you list and get details on teams. You can only list teams you're into.
You can set the fields to display in table mode in your config file ($HOME/.config/webex-helper/config.yaml):
teamsFields:
- ID
- Name
- CreatorID
- Created
`,
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("teamsFields", cmd.Flags().Lookup("teams-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := listTeams(teamsMax)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	listCmd.AddCommand(teamsCmd)
	teamsCmd.Flags().IntVarP(
		&teamsMax,
		"max",
		"m",
		20,
		"Number max of teams to list",
	)
	teamsCmd.Flags().StringSliceVar(&teamsFields, "teams-fields", defaultTeamsFields, "Teams fields to display")
}
