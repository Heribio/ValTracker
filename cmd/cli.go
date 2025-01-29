package cmd

import (
	"github.com/Heribio/ValTracker/internal/cli"

	"github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(cliCommand)
}

var cliCommand= &cobra.Command{
    Use: "cli",
    Short: "Launches the cli version of valtracker",
    Run: func(cmd *cobra.Command, args []string) {
        cli.UserInput()
    },
}
