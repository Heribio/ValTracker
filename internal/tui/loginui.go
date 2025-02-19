package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Heribio/ValTracker/internal/jsonthings"
)

type (
	errMsg error
)

const (
	name = iota
	tag
)


type loginState struct {
    inputs []textinput.Model
    focused   int
    err       error
}

func InitialModel() loginState {
    inputs := make([]textinput.Model, 2)

	inputs[name] = textinput.New()
	inputs[name].Placeholder = "Name"
	inputs[name].Focus()
	inputs[name].CharLimit = 156
	inputs[name].Width = 20

	inputs[tag] = textinput.New()
	inputs[tag].Placeholder = "Tag"
	inputs[tag].CharLimit = 6
	inputs[tag].Width = 20

	return loginState{
		inputs: inputs,
		focused: 0,
		err:       nil,
	}
}

func (m model) loginUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
    inputs := m.state.loginPage.inputs
    focusedEntry := m.state.loginPage.focused
    var cmds []tea.Cmd = make([]tea.Cmd, len(inputs))

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "enter":
            if focusedEntry == len(inputs)-1 {
                name := inputs[name].Value()
                tag := inputs[tag].Value()
                jsonthings.WriteFileData(name, tag)
                fmt.Printf("Switched to %s#%s", name, tag)

                m.state.matchListPage = MatchList(name, tag)
                return m.SwitchPage(matchListPage), nil
            }
            m.nextInput()
            return m, nil
        case "ctrl+c", "esc":
            return m, tea.Quit
        case "tab", "down":
            m.nextInput()
        case "shift+tab", "up":
            m.prevInput()
        case "<":
            m = m.SwitchPage(matchListPage)
            return m, nil
        }

        for i := range inputs {
            inputs[i].Blur()
        }
        inputs[m.state.loginPage.focused].Focus()
    }

    for i := range inputs {
        inputs[i], cmds[i] = inputs[i].Update(msg)
    }
    return m, tea.Batch(cmds...)
}


func (m model) loginView() string {
    inputs := m.state.loginPage.inputs

	return fmt.Sprintf(
		"Insert the name and tag of the valorant player\n\n%s\n\n%s",
		inputs[name].View(),
		inputs[tag].View(),
	) + "\n"
}

func (m *model) nextInput() {
    focusedEntry := m.state.loginPage.focused
    inputs := m.state.loginPage.inputs

	m.state.loginPage.focused = (focusedEntry + 1) % len(inputs)
}

func (m *model) prevInput() {
    inputs := m.state.loginPage.inputs

	m.state.loginPage.focused--
	if m.state.loginPage.focused < 0 {
		m.state.loginPage.focused = len(inputs) - 1
	}
}
