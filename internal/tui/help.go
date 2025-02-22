package tui

import "github.com/charmbracelet/bubbles/key"

func (d matchlistDelegate) ShortHelp() []key.Binding {
    return []key.Binding{keys.MatchPageBinding, keys.LoginPageBinding}
}

func (d matchlistDelegate) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{keys.LoginPageBinding}, 
		{keys.MatchPageBinding},  
    }
}
