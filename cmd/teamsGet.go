/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/utils"
)

func getTeams(teamID string) error {
	wc := utils.NewWebexTeamsClient()

	team, resp, err := wc.Teams.GetTeam(teamID)

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

	fields := viper.GetStringSlice("teamsFields")
	utils.PrintStructAsTable(*team, fields)
	return nil
}

// teamsGetCmd represents the teamsGet command
var teamsGetCmd = &cobra.Command{
	Use:   "teams",
	Short: "Get a team details",
	Long: `Get a team details
You can set the fields to display in table mode in your config file ($HOME/.config/webex-helper/config.yaml):
teamsFields:
- ID
- Name
- CreatorID
- Created
`,
	Aliases: []string{"team"},
	Args:    cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("teamsFields", cmd.Flags().Lookup("teams-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := getTeams(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	getCmd.AddCommand(teamsGetCmd)

	teamsGetCmd.Flags().StringSliceVar(&teamsFields, "teams-fields", defaultTeamsFields, "Teams fields to display")
}
