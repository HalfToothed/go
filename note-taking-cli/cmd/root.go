package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "note",
	Short: "note is a cli tool for taking and saving notes",
	Long:  "note is a cli tool for basic note taking functionality",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing note '%s' \n", err)
		os.Exit(1)
	}
}
