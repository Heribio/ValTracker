package tui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
    NavigationBindings key.Binding
	MatchPageBinding key.Binding
    SearchPageBinding key.Binding
    NextModeBinding key.Binding
    PreviousModeBinding key.Binding
    FavoriteBinding key.Binding
    QuickSwitchBinding key.Binding
    NextInputBinding key.Binding
    PreviousInputBinding key.Binding
    ConfirmBinding key.Binding
    MatchListBinding key.Binding
}

var keys = keyMap{
    NavigationBindings: key.NewBinding(
        key.WithKeys("h/j/k/l"),
        key.WithHelp("h/j/k/l", "Move left/down/up/right"),
        ),
    MatchPageBinding: key.NewBinding(
        key.WithKeys("m"),
        key.WithHelp("m", "Display more matches"),
        ),
    SearchPageBinding: key.NewBinding(
        key.WithKeys("s"),
        key.WithHelp("s", "Search"),
        ),
    NextModeBinding: key.NewBinding(
        key.WithKeys("right"),
        key.WithHelp("→", "Next mode"),
        ),
    PreviousModeBinding: key.NewBinding(
        key.WithKeys("left"),
        key.WithHelp("←", "Previous mode"),
        ),
    FavoriteBinding: key.NewBinding(
        key.WithKeys("f"),
        key.WithHelp("f", "Favorite a player"),
        ),
    QuickSwitchBinding: key.NewBinding(
        key.WithKeys("ctrl+h/j/k/l"),
        key.WithHelp("ctrl+h/j/k/l", "Switch between favorite profiles"),
        ),
    NextInputBinding: key.NewBinding(
        key.WithKeys("tab"),
        key.WithHelp("tab", "Next"),
        ),
     PreviousInputBinding: key.NewBinding(
        key.WithKeys("shift+tab"),
        key.WithHelp("shift+tab", "Previous"),
       ),
    ConfirmBinding:  key.NewBinding(
        key.WithKeys("Enter"),
        key.WithHelp("Enter", "Confirm"),
        ),
    MatchListBinding: key.NewBinding(
        key.WithKeys("esc"),
        key.WithHelp("esc", "Matchlist"),
        ),
}
