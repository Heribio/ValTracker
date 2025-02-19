package tui

import (
	"fmt"
	"io"
	"strings"

	"github.com/Heribio/ValTracker/internal/valorantapi"
	_ "github.com/Heribio/ValTracker/internal/valorantapi"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


type matchListState struct {
    list list.Model
}

type Match valorantapi.Match

func (m model) matchListUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "esc":
			return m, tea.Quit
        case "<":
            m = m.SwitchPage(loginPage)
            return m, nil
        case "m":
            return m, nil
       }
    }

    var cmd tea.Cmd
	m.state.matchListPage.list, cmd= m.state.matchListPage.list.Update(msg)
	return m, cmd
}

//Settings for margin of list
var docStyle = lipgloss.NewStyle().
    Margin(2, 3)

func (m model) matchListView() string {
    return docStyle.Render(m.state.matchListPage.list.View())
}

//Display of informations in list items
func (m Match) Title() string {
    if (m.Team == "Blue") {
        if (m.BlueTeamScore > m.RedTeamScore){
            return fmt.Sprintf("Win - %d-%d %s - %s",m.BlueTeamScore, m.RedTeamScore, m.MapName, m.CharacterName)
        }
        if (m.BlueTeamScore < m.RedTeamScore){
            return fmt.Sprintf("Loss - %d-%d %s - %s",m.BlueTeamScore, m.RedTeamScore, m.MapName, m.CharacterName)
        } else {
            return fmt.Sprintf("Draw- %d-%d %s - %s",m.BlueTeamScore, m.RedTeamScore, m.MapName, m.CharacterName)
        }
    } else {
        if (m.BlueTeamScore < m.RedTeamScore){
            return fmt.Sprintf("Win - %d-%d %s - %s",m.RedTeamScore, m.BlueTeamScore, m.MapName, m.CharacterName)
        }
        if (m.BlueTeamScore > m.RedTeamScore){
            return fmt.Sprintf("Loss - %d-%d %s - %s",m.RedTeamScore, m.BlueTeamScore, m.MapName, m.CharacterName)
        } else {
            return fmt.Sprintf("Draw- %d-%d %s - %s",m.RedTeamScore, m.BlueTeamScore, m.MapName, m.CharacterName)
        }
    }
}

func (m Match) Description() string {
    return fmt.Sprintf("Kills: %d | Deaths: %d | Assists: %d | ACS: %d", m.Kills, m.Deaths, m.Assists, m.Score/(m.BlueTeamScore+m.RedTeamScore))
}
func (i Match) FilterValue() string { return i.StartedAt }

var(
    winStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, false, true).Foreground(lipgloss.Color("#4bff27")).Bold(true).Padding(0, 2)
    lossStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, false, true).Foreground(lipgloss.Color("#f7659")).Bold(true).Padding(0, 2)
    selectedStyle = lipgloss.NewStyle().Bold(true).Padding(0, 2)
	cursorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Bold(true)
	descStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Italic(true).PaddingLeft(4)
	separator      = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render("──────────────────────────")
)

type matchlistDelegate struct{}

func (d matchlistDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	it, ok := listItem.(Match)
	if !ok {
		return
	}

	title := it.Title()
    description := it.Description()
	var style lipgloss.Style

	if index == m.Index(){
        style = selectedStyle
	} else if strings.Contains(strings.ToLower(title), "win") {
		style = winStyle
    } else {
		style = lossStyle
	}

	cursor := "  " // No cursor for unselected items
	if index == m.Index() {
		cursor = cursorStyle.Render(">") // Highlight selected item
	}

	fmt.Fprintf(w, "%s %s\n%s\n", cursor, style.Render(title), descStyle.Render(description))
}


func (d matchlistDelegate) Height() int  { return 3 }
func (d matchlistDelegate) Spacing() int { return 1 }
func (d matchlistDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd{
	return nil
}

//Get matches
func MatchList(name string, tag string) matchListState {
    puuid := valorantapi.GetAccountPUUID(name, tag)
    matchList := valorantapi.GetAccountMatches(puuid, "1",  "eu", "competitive")

    items := make([]list.Item, len(matchList))
    for i := range len(matchList){
        items[i] = Match(matchList[i])
    }

    windowWidth, windowHeight := 50, 45
    l := list.New(items, matchlistDelegate{}, windowWidth, windowHeight)

    m := matchListState{list: l}
    return m
}
