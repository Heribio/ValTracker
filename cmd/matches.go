package cmd

import(
    "errors"
    "fmt"

    "github.com/spf13/cobra"
    "github.com/Heribio/ValTracker/internal/valorantapi"
)

func init() {
    rootCmd.AddCommand(matchesCmd)
}

var matchesCmd = &cobra.Command{
    Use: "matches",
    Short: "Get you latest matches",
    Args: func(cmd *cobra.Command, args []string) error {
        if len(args) < 2 {
            return errors.New(`Not enough arguments specified. "Expected valtracker matches [name] [tag]"`)
        }
        if len(args) > 2 {
            return errors.New(`Too many arguments specified. Expected "valtracker matches [name] [tag]"`)
        }
        return nil
    },
    Run: func(cmd *cobra.Command, args []string) {
        puuid := valorantapi.GetAccountPUUID(args [0] , args[1])
        matches := valorantapi.FormatMatches(valorantapi.GetAccountMatches(puuid))
        for _, match   := range matches{
            fmt.Println(match.MapName)
            fmt.Println(match.Mode)
            fmt.Println(match.Kills)
            fmt.Println(match.Deaths)
            fmt.Println(match.CharacterName)
            fmt.Println("---")
        }
    },
}

