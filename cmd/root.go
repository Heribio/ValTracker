package cmd

import(
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "valtracker",
    Short: "Find your valorant stats",
        Run: func(cmd *cobra.Command, args []string) {
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
