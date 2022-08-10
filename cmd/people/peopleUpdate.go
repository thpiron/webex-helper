package people

import (
	"fmt"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
)

var (
	updatedEmails      []string
	updatedDisplayName string
	updatedFirstName   string
	updatedLastName    string
	updatedAvatar      string
	updatedOrgID       string
	updatedRoles       []string
	updatedLicenses    []string
)

func UpdatePerson(personID string) error {
	wc := utils.NewWebexTeamsClient()

	person, resp, err := wc.People.Update(personID, &webexteams.PersonRequest{
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

	if cmd.JsonOutput {
		fmt.Println(string(resp.Body()))
		return nil
	}

	fields := viper.GetStringSlice("peopleFields")
	utils.PrintStructAsTable(*person, fields)
	return nil
}

// peopleUpdateCmd represents the peopleUpdate command
var peopleUpdateCmd = &cobra.Command{
	Use:     "people",
	Short:   "Update a person",
	Long:    `Update a person.`,
	Example: `webex-helper update people --display-name "My new name" <peopleID>`,
	Args:    cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("peopleFields", cmd.Flags().Lookup("people-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := UpdatePerson(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.UpdateCmd.AddCommand(peopleUpdateCmd)
	peopleUpdateCmd.Flags().StringSliceVar(&peopleFields, "people-fields", defaultPeopleFields, "People fields to display")

	peopleUpdateCmd.Flags().StringSliceVar(&updatedEmails, "emails", nil, "List of user's emails")
	peopleUpdateCmd.Flags().StringVar(&updatedDisplayName, "display-name", "", "User's display name")
	peopleUpdateCmd.Flags().StringVar(&updatedFirstName, "first-name", "", "User's first name")
	peopleUpdateCmd.Flags().StringVar(&updatedLastName, "last-name", "", "User's last name")
	peopleUpdateCmd.Flags().StringVar(&updatedAvatar, "avatar", "", "User's avatar")
	peopleUpdateCmd.Flags().StringVar(&updatedOrgID, "org-id", "", "User's orgID")
	peopleUpdateCmd.Flags().StringSliceVar(&updatedRoles, "roles", nil, "List of user's roles")
	peopleUpdateCmd.Flags().StringSliceVar(&updatedLicenses, "licenses", nil, "List of user's licenses")

	peopleUpdateCmd.MarkFlagRequired("display-name")
}
