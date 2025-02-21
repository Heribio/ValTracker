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
        case "ctrl+c", "esc":
			return m, tea.Quit
        case "<":
            m = m.SwitchPage(loginPage)
            return m, nil
        case "m":
            var existingMatches []valorantapi.Match
            for _, item := range m.state.matchListPage.list.Items() {
                if match, ok := item.(Match); ok {
                    existingMatches = append(existingMatches, valorantapi.Match(match))
                }
            }

            updatedMatches := valorantapi.AppendMatchList(existingMatches, "1", "eu", "competitive")

            newItems := make([]list.Item, len(updatedMatches))
            for i, match := range updatedMatches {
                newItems[i] = Match(match)
            }

            m.state.matchListPage.list.SetItems(newItems)
            return m, nil
       }
    }

    var cmd tea.Cmd
	m.state.matchListPage.list, cmd= m.state.matchListPage.list.Update(msg)
	return m, cmd
}

//Settings for margin of list
var docStyle = lipgloss.NewStyle().
    Margin(3, 4)

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
    return fmt.Sprintf("Kills: %d | Deaths: %d | Assists: %d |          K/D: %.1f | ACS: %d", m.Kills, m.Deaths, m.Assists, float32(m.Kills)/float32(m.Deaths), m.Score/(m.BlueTeamScore+m.RedTeamScore))
}
func (i Match) FilterValue() string { return i.StartedAt }

//Get matches
func MatchList(name string, tag string) matchListState {
    puuid := valorantapi.GetAccountPUUID(name, tag)
    matchList := valorantapi.GetAccountMatches(puuid, "1",  "eu", "competitive")

    items := make([]list.Item, len(matchList))
    for i := range len(matchList){
        items[i] = Match(matchList[i])
    }

    windowWidth, windowHeight := 100, 40 
    l := list.New(items, matchlistDelegate{}, windowWidth, windowHeight)
    l.Title = "Matchlist"

    m := matchListState{list: l}
    return m
}
