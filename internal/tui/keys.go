package tui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	MatchPageBinding key.Binding
    LoginPageBinding key.Binding
}

var keys = keyMap{
    MatchPageBinding: key.NewBinding(
        key.WithKeys("m"),
        key.WithHelp("m", "Display more matches"),
    ),
    LoginPageBinding: key.NewBinding(
        key.WithKeys("l"),
        key.WithHelp("l", "Login"),
    ),
}
