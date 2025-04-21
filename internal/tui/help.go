package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

func (d matchlistDelegate) ShortHelp() []key.Binding {
    return []key.Binding{keys.MatchPageBinding, keys.SearchPageBinding, keys.PreviousModeBinding, keys.NextModeBinding, keys.QuickSwitchBinding}
}

func (d matchlistDelegate) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{keys.SearchPageBinding}, 
		{keys.MatchPageBinding},  
		{keys.PreviousModeBinding},  
		{keys.NextModeBinding},
        {keys.QuickSwitchBinding},
    }
}


var helpKeyStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
var helpDescStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render

func HelpView(bindings []key.Binding) string {
    var b strings.Builder

    for _, kb := range bindings {
        help := kb.Help()
        if help.Key != "" || help.Desc != "" {
            line := fmt.Sprintf("%s  %s",
                helpKeyStyle(help.Key),
                helpDescStyle(help.Desc),
            )
            b.WriteString(line)
            b.WriteString("\n\n")
        }
    }

    return b.String()
}

func shortHelpView(bindings []key.Binding) string {
    var b strings.Builder

    var parts []string
    for _, kb := range bindings {
        help := kb.Help()
        if help.Key != "" || help.Desc != "" {
            part := fmt.Sprintf("%s %s",
                helpKeyStyle("[" + help.Key + "]"),
                helpDescStyle(help.Desc),
            )
            parts = append(parts, part)
        }
    }

    // Join them with spaces (horizontally aligned)
    b.WriteString(strings.Join(parts, "  "))

    return b.String()
}
