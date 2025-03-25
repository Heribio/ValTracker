package tui

import "github.com/charmbracelet/bubbles/key"

func (d matchlistDelegate) ShortHelp() []key.Binding {
    return []key.Binding{keys.MatchPageBinding, keys.SearchPageBinding, keys.PreviousModeBinding, keys.NextModeBinding}
}

func (d matchlistDelegate) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{keys.SearchPageBinding}, 
		{keys.MatchPageBinding},  
		{keys.PreviousModeBinding},  
		{keys.NextModeBinding},  
    }
}
