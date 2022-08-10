/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thpiron/webex-helper/utils"
)

func deleteTeam(teamID string) error {
	wc := utils.NewWebexTeamsClient()

	resp, err := wc.Teams.DeleteTeam(teamID)

	if err != nil {
		return err
	}
	if err := utils.CheckWebexError(*resp); err != nil {
		return err
	}
	fmt.Println("Team deleted.")
	return nil
}

// teamsDeleteCmd represents the teamsDelete command
var teamsDeleteCmd = &cobra.Command{
	Use:     "teams",
	Short:   "Delete a team",
	Long:    `Delete a team`,
	Aliases: []string{"team"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := deleteTeam(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	deleteCmd.AddCommand(teamsDeleteCmd)
}
