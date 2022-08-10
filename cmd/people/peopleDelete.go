package people

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thpiron/webex-helper/cmd"
	"github.com/thpiron/webex-helper/utils"
)

func personDelete(personID string) error {
	wc := utils.NewWebexTeamsClient()

	resp, err := wc.People.DeletePerson(personID)

	if err != nil {
		return err
	}
	if err := utils.CheckWebexError(*resp); err != nil {
		return err
	}
	fmt.Println("Person deleted.")
	return nil
}

// peopleDeleteCmd represents the peopleDelete command
var peopleDeleteCmd = &cobra.Command{
	Use:   "people",
	Short: "Delete a person",
	Long:  `Delete a person`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := personDelete(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	cmd.DeleteCmd.AddCommand(peopleDeleteCmd)
}
