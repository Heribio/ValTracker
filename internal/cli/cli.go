package cli

import(
    "log"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"

	"github.com/Heribio/ValTracker/internal/valorantapi"
)

func UserInput() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

const (
	name = iota
	tag
)

type model struct {
    inputs []textinput.Model
    focused   int
    err       error
}

func initialModel() model {
	var inputs []textinput.Model = make([]textinput.Model, 2)

	inputs[name] = textinput.New()
	inputs[name].Placeholder = "Name"
	inputs[name].Focus()
	inputs[name].CharLimit = 156
	inputs[name].Width = 20

	inputs[tag] = textinput.New()
	inputs[tag].Placeholder = "Tag"
	inputs[tag].CharLimit = 6
	inputs[tag].Width = 20

	return model{
		inputs: inputs,
		focused: 0,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				valapi(m.inputs[name].Value(), m.inputs[tag].Value())
				return m, tea.Quit
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()


	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(
		"Insert the name and tag of the valorant player\n\n%s\n\n%s",
		m.inputs[name].View(),
		m.inputs[tag].View(),
	) + "\n"
}

func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}

func valapi(name string, tag string) {
    vapi := valorantapi.Authorization()

	puuid := valorantapi.GetAccountPUUID(name, tag, vapi)	
	valorantapi.GetAccountMatches(puuid, vapi)
	valorantapi.GetAccountMMR(puuid, vapi)
}
