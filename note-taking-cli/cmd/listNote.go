package cmd

import (
	"github.com/spf13/cobra"
)

var listNoteCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "notes"},
	Short:   "List the Saved Notes",
	Long:    "List the saved notes",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		listNote()
	},
}

func init() {
	rootCmd.AddCommand(listNoteCmd)
}
