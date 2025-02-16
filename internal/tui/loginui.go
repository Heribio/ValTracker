package tui

import(
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"

	"github.com/Heribio/ValTracker/internal/valorantapi"
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
    var cmds []tea.Cmd = make([]tea.Cmd, len(m.state.loginPage.inputs))

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "enter":
            if m.state.loginPage.focused == len(m.state.loginPage.inputs)-1 {
                m.name = m.state.loginPage.inputs[name].Value()
                m.tag = m.state.loginPage.inputs[tag].Value()
                fmt.Printf("registered user %s#%s", m.name, m.tag)
                return m, nil
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

        for i := range m.state.loginPage.inputs {
            m.state.loginPage.inputs[i].Blur()
        }
        m.state.loginPage.inputs[m.state.loginPage.focused].Focus()
    }

    for i := range m.state.loginPage.inputs {
        m.state.loginPage.inputs[i], cmds[i] = m.state.loginPage.inputs[i].Update(msg)
    }

    return m, tea.Batch(cmds...)
}


func (m model) loginView() string {
	return fmt.Sprintf(
		"Insert the name and tag of the valorant player\n\n%s\n\n%s",
		m.state.loginPage.inputs[name].View(),
		m.state.loginPage.inputs[tag].View(),
	) + "\n"
}

func (m *model) nextInput() {
	m.state.loginPage.focused = (m.state.loginPage.focused + 1) % len(m.state.loginPage.inputs)
}

func (m *model) prevInput() {
	m.state.loginPage.focused--
	if m.state.loginPage.focused < 0 {
		m.state.loginPage.focused = len(m.state.loginPage.inputs) - 1
	}
}

func valapi(name string, tag string) {
	puuid := valorantapi.GetAccountPUUID(name, tag)	
	valorantapi.GetAccountMatches(puuid)
	valorantapi.GetAccountMMR(puuid)
}
