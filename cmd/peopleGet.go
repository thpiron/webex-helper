/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/utils"
)

func GetPeople(peopleID string) error {
	wc := utils.NewWebexTeamsClient()
	var (
		person *webexteams.Person
		resp   *resty.Response
		err    error
	)
	if peopleID == "me" {
		person, resp, err = wc.People.GetMe()
	} else {
		person, resp, err = wc.People.GetPerson(peopleID)
	}
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

// peopleGetCmd represents the peopleGet command
var peopleGetCmd = &cobra.Command{
	Use:   "people",
	Short: "Get a people details",
	Long: `Retrieve a person details from its id
You can look at your own detail by using me instead of an ID.
You can set the fields to display in table mode in your config file ($HOME/.config/webex-helper/config.yaml):
peopleFields:
	- ID
	- Emails
	- SIPAddresses
	- PhoneNumbers
	- DisplayName
	- NickName
	- FirstName
	- LastName
	- Avatar
	- OrgID
	- Roles
	- Licenses
	- Created
	- LastModified
	- TimeZone
	- LastActivity
	- Status
	- InvitePending
	- LoginEnabled
	- PersonType
`,
	Args:    cobra.ExactArgs(1),
	Example: "webex-helper get people [flags] <peopleID>",
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("peopleFields", cmd.Flags().Lookup("people-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := GetPeople(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	getCmd.AddCommand(peopleGetCmd)

	peopleGetCmd.Flags().StringSliceVar(&peopleFields, "people-fields", defaultPeopleFields, "People fields to display")
}
