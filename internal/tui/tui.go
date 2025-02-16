package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func Run() {
    model, _:= NewModel(
        lipgloss.DefaultRenderer(),
	)

    if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}

type page = int

const(
    loginPage = iota
    matchListPage
)

type state struct {
    loginPage   loginState
    matchListPage  matchListState
}

type model struct {
    page         page
    renderer     *lipgloss.Renderer
    accountPages []page
    switched     bool
    state        state
    name         string
    tag          string
}

func NewModel(renderer *lipgloss.Renderer) (tea.Model, error) {
    result := model{
        page: loginPage,
        renderer: renderer,
        accountPages: []page{
            loginPage,
            matchListPage,
        },
        state: state{
            loginPage: InitialModel(),
            matchListPage: MatchList("", ""),
        },
    }
    return result, nil
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) SwitchPage(page page) model {
    m.page = page
    m.switched = true
    return m
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "esc":
            return m, tea.Quit
        }
    }

    var cmd tea.Cmd
    switch m.page{
    case loginPage:
        var updatedModel tea.Model
        updatedModel, cmd = m.loginUpdate(msg)
        if newModel, ok := updatedModel.(model); ok {
            m = newModel
        }
    case matchListPage:
        var updatedModel tea.Model
        updatedModel, cmd = m.matchListUpdate(msg)
        if newModel, ok := updatedModel.(model); ok {
            m = newModel
        }
    }

    if m.switched {
		m.switched = false
	}

    var headerCmd tea.Cmd

    cmds := []tea.Cmd{headerCmd}
    cmds = append(cmds, cmd)
    return m, tea.Batch(cmds...)
}

func (m model) View() string {
    switch m.page{
    case loginPage:
        return m.loginView()
    case matchListPage:
        return m.matchListView()
    default:
        return m.matchListView()
    }
}
