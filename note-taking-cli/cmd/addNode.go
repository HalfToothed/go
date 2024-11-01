package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addNoteCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"input", "take"},
	Short:   "Add a note",
	Long:    "Save a note",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Added Note: %s\n", addNote(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(addNoteCmd)
}
