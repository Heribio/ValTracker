package tui

import (
	"fmt"

	"github.com/Heribio/ValTracker/internal/valorantapi"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type selectedMatchState struct {
    list list.Model
}

type selectedMatch valorantapi.Match

func (m model) selectMatchUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
			return m, tea.Quit
        case "enter":
            if selectedItem, ok := m.state.selectedMatchPage.list.SelectedItem().(Player); ok {
                m.state.matchListPage = MatchList(selectedItem.Username, selectedItem.Tag, "competitive")
                m = m.SwitchPage(matchListPage)
            }
            return m, nil
        case "esc":
            m = m.SwitchPage(matchListPage)
            return m, nil
        }
    }
    var cmd tea.Cmd
	m.state.selectedMatchPage.list, cmd= m.state.selectedMatchPage.list.Update(msg)
	return m, cmd
}

func (m Player) Title() string {
    return fmt.Sprintf("%s - %s %s", m.Username, m.CharacterName, m.Rank)
}

func (m Player) Description() string {
    return fmt.Sprintf("Kills: %d | Deaths: %d | Assists: %d          K/D: %.1f | ACS: %d",
        m.Kills, m.Deaths, m.Assists, float32(m.Kills)/float32(m.Deaths), m.Score/m.Rounds)
}

func (i Player) FilterValue() string {return i.Username}


func (m model) selectedMatchView() string {
    return docStyle.Render(m.state.selectedMatchPage.list.View())
}

type Player valorantapi.Player

func SelectedMatchList(selectedMatch string) selectedMatchState{
    match := valorantapi.GetMatch(selectedMatch)
    players := valorantapi.GetPlayers(match)

    items := make([]list.Item, len(players))
    for i := range len(players){
        items[i] = Player(players[i])
    }

    windowWidth, windowHeight := 100, 40 
    list := list.New(items, selectedmatchDelegate{match: match}, windowWidth, windowHeight)

    list.Title = fmt.Sprintf("%s", match.Data.Metadata.GameStartPatched)
    sm := selectedMatchState{list: list}
    return sm
}
