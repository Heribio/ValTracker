package cmd

import(
    "fmt"

    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(helpCommand)
}

var helpCommand= &cobra.Command{
    Use: "help",
    Short: "Help command",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(`
amazing
        dwadwad
            `)
    },
}
