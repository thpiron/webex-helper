package cmd

import (
	"github.com/spf13/cobra"
)

// UpdateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a resouce",
	Long:  `Update a resource`,
}

func init() {
	rootCmd.AddCommand(UpdateCmd)
}
