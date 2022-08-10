/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/utils"
)

var (
	emails      []string
	displayName string
	firstName   string
	lastName    string
	avatar      string
	orgID       string
	roles       []string
	licenses    []string
)

func CreatePerson() error {
	wc := utils.NewWebexTeamsClient()

	person, resp, err := wc.People.CreatePerson(&webexteams.PersonRequest{
		Emails:      updatedEmails,
		DisplayName: updatedDisplayName,
		FirstName:   updatedFirstName,
		LastName:    updatedLastName,
		Avatar:      updatedAvatar,
		OrgID:       updatedOrgID,
		Roles:       updatedRoles,
		Licenses:    updatedLicenses,
	})
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

	fields := viper.GetStringSlice("peopleFields")
	utils.PrintStructAsTable(*person, fields)
	return nil
}

// peopleCreateCmd represents the peopleCreate command
var peopleCreateCmd = &cobra.Command{
	Use:   "people",
	Short: "Create a person",
	Long:  `Create a person`,
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("peopleFields", cmd.Flags().Lookup("people-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := CreatePerson()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	createCmd.AddCommand(peopleCreateCmd)
	peopleCreateCmd.Flags().StringSliceVar(&peopleFields, "people-fields", defaultPeopleFields, "People fields to display")

	peopleCreateCmd.Flags().StringSliceVar(&emails, "emails", nil, "List of user's emails")
	peopleCreateCmd.Flags().StringVar(&displayName, "display-name", "", "User's display name")
	peopleCreateCmd.Flags().StringVar(&firstName, "first-name", "", "User's first name")
	peopleCreateCmd.Flags().StringVar(&lastName, "last-name", "", "User's last name")
	peopleCreateCmd.Flags().StringVar(&avatar, "avatar", "", "User's avatar")
	peopleCreateCmd.Flags().StringVar(&orgID, "org-id", "", "User's orgID")
	peopleCreateCmd.Flags().StringSliceVar(&roles, "roles", nil, "List of user's roles")
	peopleCreateCmd.Flags().StringSliceVar(&licenses, "licenses", nil, "List of user's licenses")

	peopleCreateCmd.MarkFlagRequired("display-name")
}
