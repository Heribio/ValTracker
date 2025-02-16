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
    return fmt.Sprintf("%d-%d %s - %s",m.BlueTeamScore, m.RedTeamScore, m.MapName, m.CharacterName)
}

func (m Match) Description() string {
    return fmt.Sprintf("Kills: %d | Deaths: %d | Assists: %d | ACS: %d", m.Kills, m.Deaths, m.Assists, m.Score/(m.BlueTeamScore+m.RedTeamScore))
}
func (i Match) FilterValue() string { return i.Id}

//Get matches
func MatchList(name string, tag string) matchListState {
    //TODO Get name and tag from loginui
    puuid := valorantapi.GetAccountPUUID(name, tag)
    resp := valorantapi.GetAccountMatches(puuid)
    matchList := valorantapi.FormatMatches(resp)

    items := make([]list.Item, len(matchList))
    for i := range len(matchList){
        items[i] = Match(matchList[i])
    }

    // winColor := lipgloss.Color("#4bff27")
    // lossColor = lipgloss.Color("f7659")

    windowWidth, windowHeight := 50, 45
    d := list.NewDefaultDelegate()
    // d.Styles.SelectedTitle = d.Styles.SelectedTitle.Foreground(winColor)

    m := matchListState{list: list.New(items, d, windowWidth, windowHeight)}
    return m
}
