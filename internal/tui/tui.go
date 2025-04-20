package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Heribio/ValTracker/internal/jsonthings"
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
    overviewPage = iota
    searchPage
    matchListPage
    selectedMatchPage
)

type state struct {
    overviewPage overViewState
    searchPage   searchState
    matchListPage  matchListState
    selectedMatchPage selectedMatchState
}

type model struct {
    page         page
    renderer     *lipgloss.Renderer
    accountPages []page
    state        state
    name         string
    tag          string
    selectedMatch *Match
    mode         string
}


func NewModel(renderer *lipgloss.Renderer) (tea.Model, error) {
    result := model{
        page: matchListPage,
        renderer: renderer,
        accountPages: []page{
            overviewPage,
            searchPage,
            matchListPage,
            selectedMatchPage,
        },
        state: state{
            searchPage: InitialModel(),
            matchListPage: MatchList(jsonthings.GetFileData("data.json").Name, jsonthings.GetFileData("data.json").Tag, "competitive"),
            selectedMatchPage: SelectedMatchList(""),
        },
        mode: "competitive",
    }
    return result, nil
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) SwitchPage(page page) model {
    m.page = page
    return m
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
        }
    }

    var cmd tea.Cmd
    switch m.page{
    case overviewPage:
        var updatedModel tea.Model
        updatedModel, cmd = m.overViewUpdate(msg)
        if newModel, ok := updatedModel.(model); ok {
            m = newModel
        }
    case searchPage:
        var updatedModel tea.Model
        updatedModel, cmd = m.searchUpdate(msg)
        if newModel, ok := updatedModel.(model); ok {
            m = newModel
        }
    case matchListPage:
        var updatedModel tea.Model
        updatedModel, cmd = m.matchListUpdate(msg)
        if newModel, ok := updatedModel.(model); ok {
            m = newModel
        }
    case selectedMatchPage:
        var updatedModel tea.Model
        updatedModel, cmd = m.selectMatchUpdate(msg)
        if newModel, ok := updatedModel.(model); ok {
            m = newModel
        }
     }
   
    var headerCmd tea.Cmd

    cmds := []tea.Cmd{headerCmd}
    cmds = append(cmds, cmd)
    return m, tea.Batch(cmds...)
}

func (m model) View() string {
    switch m.page{
    case overviewPage:
        return m.overViewView()
    case searchPage:
        return m.searchView()
    case matchListPage:
        return m.matchListView()
    case selectedMatchPage:
        return m.selectedMatchView()
    default:
        return ""
    }
}
