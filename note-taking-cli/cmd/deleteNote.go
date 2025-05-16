package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteNoteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "remove"},
	Short:   "Delete a note",
	Long:    "Delete a note though its Id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(deleteNote(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(deleteNoteCmd)
}
