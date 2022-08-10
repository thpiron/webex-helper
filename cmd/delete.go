package cmd

import (
	"github.com/spf13/cobra"
)

// DeleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a resource",
	Long:  `Delete a resource.`,
}

func init() {
	rootCmd.AddCommand(DeleteCmd)
}
