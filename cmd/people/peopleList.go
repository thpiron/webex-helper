package people

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
	webexteams "github.com/thpiron/webex-teams/sdk"
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
	wc := utils.NewWebexTeamsClient()

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
	if err := utils.CheckWebexError(*resp); err != nil {
		return err
	}

	if cmd.JsonOutput {
		fmt.Println(string(resp.Body()))
		return nil
	}

	s := make([]interface{}, len(peoples.Items))
	for i, v := range peoples.Items {
		s[i] = v
	}

	fields := viper.GetStringSlice("peopleFields")
	utils.PrintStructSliceAsTable(s, fields)
	return nil
}

// peopleListCmd represents the people command
var peopleListCmd = &cobra.Command{
	Use:   "people",
	Short: "List peoples",
	Long: `
People command let you search/list people in webex API.
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
	$ webex-helper list people -d John Doe (for non admin user -d display_name is necessary)
`,
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("peopleFields", cmd.Flags().Lookup("people-fields"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := ListPeople(peopleEmail, peopleDisplayName, peopleOrgID, peopleMax)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.ListCmd.AddCommand(peopleListCmd)
	peopleListCmd.Flags().StringVarP(
		&peopleEmail,
		"email",
		"e",
		"",
		"email address of the person to search (either this or display-name are required for non admin users)",
	)
	peopleListCmd.Flags().StringVarP(
		&peopleDisplayName,
		"display-name",
		"d",
		"",
		"display name of the person to search (either this or email are required for non admin users)",
	)
	peopleListCmd.Flags().StringVarP(
		&peopleOrgID,
		"orgID",
		"o",
		"",
		"orgID of the persons to list",
	)
	peopleListCmd.Flags().IntVarP(
		&peopleMax,
		"max",
		"m",
		20,
		"Number max of people to list",
	)
	peopleListCmd.Flags().StringSliceVar(&peopleFields, "people-fields", defaultPeopleFields, "People fields to display")
}
