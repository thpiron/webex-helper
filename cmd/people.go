package cmd

import (
	"fmt"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	peopleEmail         string
	peopleDisplayName   string
	peopleOrgID         string
	peopleMax           int
	peopleFields        []string
	defaultPeopleFields = []string{"ID", "Emails", "DisplayName"}
)

func ListPeople(email, displayName, orgID string, max int) error {
	wc := NewWebexTeamsClient()
	queryParams := &webexteams.ListPeopleQueryParams{
		Email:       email,
		DisplayName: displayName,
		OrgID:       orgID,
		Max:         max,
	}
	peoples, resp, err := wc.People.ListPeople(queryParams)
	if err != nil {
		return err
	}
	if err := checkWebexError(*resp); err != nil {
		return err
	}

	if jsonOutput {
		fmt.Println(string(resp.Body()))
		return nil
	}

	s := make([]interface{}, len(peoples.Items))
	for i, v := range peoples.Items {
		s[i] = v
	}

	fields := viper.GetStringSlice("peopleFields")
	printStructSliceAsTable(s, fields)
	return nil
}

// peopleCmd represents the people command
var peopleCmd = &cobra.Command{
	Use:   "people",
	Short: "Interact with people endpoints",
	Long: `
People command let you get details and list people in webex API.
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
Example:
	$ webex-helper people list -d John Doe (for non admin user -d display_name is necessary)
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ListPeople(peopleEmail, peopleDisplayName, peopleOrgID, peopleMax)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(peopleCmd)
	peopleCmd.Flags().StringVarP(
		&peopleEmail,
		"email",
		"e",
		"",
		"email address of the person to search (either this or display-name are required for non admin users)",
	)
	peopleCmd.Flags().StringVarP(
		&peopleDisplayName,
		"display-name",
		"d",
		"",
		"display name of the person to search (either this or email are required for non admin users)",
	)
	peopleCmd.Flags().StringVarP(
		&peopleOrgID,
		"orgID",
		"o",
		"",
		"orgID of the persons to list",
	)
	peopleCmd.Flags().IntVarP(
		&peopleMax,
		"max",
		"m",
		20,
		"Number max of people to list",
	)
	peopleCmd.Flags().StringSliceVar(&peopleFields, "people-fields", defaultPeopleFields, "People fields to display")
	viper.BindPFlag("peopleFields", peopleCmd.Flags().Lookup("people-fields"))
}
