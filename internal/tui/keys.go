package tui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	MatchPageBinding key.Binding
    SearchPageBinding key.Binding
    NextModeBinding key.Binding
    PreviousModeBinding key.Binding
}

var keys = keyMap{
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
}
