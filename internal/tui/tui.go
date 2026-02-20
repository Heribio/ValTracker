package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Heribio/ValTracker/internal/jsonthings"
	"github.com/Heribio/ValTracker/internal/valorantapi"
)

func Run() {
	model, _ := NewModel(
		lipgloss.DefaultRenderer(),
	)
	// size := tea.WindowSize()

	if _, err := tea.NewProgram(model, tea.WithAltScreen(), tea.WithFPS(120)).Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}

type page = int

const (
	overviewPage = iota
	searchPage
	matchListPage
	selectedMatchPage
)

type state struct {
	overviewPage      overViewState
	searchPage        searchState
	matchListPage     matchListState
	selectedMatchPage selectedMatchState

	width  int
	height int
}

type model struct {
	page          page
	renderer      *lipgloss.Renderer
	accountPages  []page
	state         state
	name          string
	tag           string
	selectedMatch *Match
	mode          string
	searchOpen    bool
}

func NewModel(renderer *lipgloss.Renderer) (tea.Model, error) {
	var username valorantapi.Username
	jsonthings.ReadData("data.json", &username)

	result := model{
		page:     matchListPage,
		renderer: renderer,
		accountPages: []page{
			overviewPage,
			matchListPage,
			selectedMatchPage,
		},
		state: state{
			searchPage:        InitialModel(),
			matchListPage:     MatchList(username.Name, username.Tag, "competitive"),
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
	case tea.WindowSizeMsg:
		m.state.width = msg.Width
		m.state.height = msg.Height
		m.resizeMatchList()
		m.resizeSelectedMatch()
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	// When the search popup is open, route all input there first.
	if m.searchOpen {
		var updatedModel tea.Model
		var cmd tea.Cmd
		updatedModel, cmd = m.searchUpdate(msg)
		if newModel, ok := updatedModel.(model); ok {
			m = newModel
		}
		return m, cmd
	}

	// Open the search popup with 's' from the match list page.
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		if keyMsg.String() == "s" && m.page == matchListPage {
			m.searchOpen = true
			m.state.searchPage = InitialModel()
			return m, nil
		}
	}

	var cmd tea.Cmd
	switch m.page {
	case overviewPage:
		var updatedModel tea.Model
		updatedModel, cmd = m.overViewUpdate(msg)
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
	switch m.page {
	case overviewPage:
		return m.overViewView()
	case matchListPage:
		return m.matchListView()
	case selectedMatchPage:
		return m.selectedMatchView()
	default:
		return ""
	}
}
