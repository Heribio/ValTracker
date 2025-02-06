package tui

import (
	"fmt"

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
        case "ctrl+c":
			return m, tea.Quit
        case "<":
            m = m.SwitchPage(loginPage)
            return m, nil
       }
    }

    var cmd tea.Cmd
	m.state.matchListPage.list, cmd= m.state.matchListPage.list.Update(msg)
	return m, cmd
}

//Settings for margin of list
var docStyle = lipgloss.NewStyle().Margin(2, 3)

func (m model) matchListView() string {
    return docStyle.Render(m.state.matchListPage.list.View())
}

//Display of informations in list items
func (m Match) Title() string {
    return fmt.Sprintf("%s - %s", m.MapName, m.CharacterName)
}

func (m Match) Description() string {
    return fmt.Sprintf("Kills: %d | Deaths: %d | Assists: %d", m.Kills, m.Deaths, m.Assists)
}
func (i Match) FilterValue() string { return i.Id}

func getMatch(match valorantapi.Match) Match {
    i := Match{
        Id: match.Id,
        MapName: match.MapName,
        Mode: match.Mode,
        Kills: match.Kills,
        Deaths: match.Deaths,
        Assists: match.Assists,
        CharacterName: match.CharacterName,
    }
    return i
}

//Get matches
func MatchList() matchListState {
    //TODO Get name and tag from loginui
    puuid := valorantapi.GetAccountPUUID("Name", "Tag")
    resp := valorantapi.GetAccountMatches(puuid)
    formattedMatches := valorantapi.FormatMatches(resp)

    items := make([]list.Item, len(formattedMatches))
    for i := range len(formattedMatches){
        items[i] = getMatch(formattedMatches[i])
    }

    m := matchListState{list: list.New(items, list.NewDefaultDelegate(), 50, 40)}
    return m
}
