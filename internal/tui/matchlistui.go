package tui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Heribio/ValTracker/internal/jsonthings"
	"github.com/Heribio/ValTracker/internal/valorantapi"
	_ "github.com/Heribio/ValTracker/internal/valorantapi"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type matchListState struct {
	list list.Model

	width  int
	height int
}

type Match valorantapi.Match

var matchpage = 2

var modes = []string{"competitive", "unrated", "deathmatch", "swiftplay"}

func (m *model) SwitchMode(forward bool) {
	for i, mode := range modes {
		if mode == m.mode {
			if forward {
				m.mode = modes[(i+1)%len(modes)]
			} else {
				m.mode = modes[(i-1+len(modes))%len(modes)]
			}
			return
		}
	}
}

func (m model) matchListUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.state.matchListPage.height = msg.Height
		m.state.matchListPage.width = msg.Width

		vMargin := 3 * 2
		hMargin := 4 * 2

		helpView := shortHelpView([]key.Binding{
			keys.NavigationBindings,
			keys.SearchPageBinding,
			keys.MatchPageBinding,
			keys.PreviousModeBinding,
			keys.NextModeBinding,
			keys.QuickSwitchBinding,
		})
		helpHeight := lipgloss.Height(helpView)

		listWidth := msg.Width - hMargin
		listHeight := msg.Height - vMargin - helpHeight

		m.state.matchListPage.list.SetSize(listWidth, listHeight)

		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.PreviousModeBinding):
			m.SwitchMode(false)

			var player valorantapi.Username
			jsonthings.ReadData("data.json", &player)
			name := player.Name
			tag := player.Tag
			puuid := valorantapi.GetAccountPUUID(name, tag)
			matchlist := valorantapi.GetAccountMatches(puuid, "1", "eu", m.mode)

			newItems := make([]list.Item, len(matchlist))
			for i, match := range matchlist {
				newItems[i] = Match(match)
			}

			m.state.matchListPage.list.SetItems(newItems)
			mode := strings.ToUpper(m.mode[:1]) + m.mode[1:]
			m.state.matchListPage.list.Title = fmt.Sprintf("%s: %s#%s", mode, name, tag)
			m.resizeMatchList()
			return m, nil

		case key.Matches(msg, keys.NextModeBinding):
			m.SwitchMode(true)

			var player valorantapi.Username
			jsonthings.ReadData("data.json", &player)
			name := player.Name
			tag := player.Tag
			puuid := valorantapi.GetAccountPUUID(name, tag)
			matchlist := valorantapi.GetAccountMatches(puuid, "1", "eu", m.mode)

			newItems := make([]list.Item, len(matchlist))
			for i, match := range matchlist {
				newItems[i] = Match(match)
			}

			m.state.matchListPage.list.SetItems(newItems)
			mode := strings.ToUpper(m.mode[:1]) + m.mode[1:]
			m.state.matchListPage.list.Title = fmt.Sprintf("%s: %s#%s", mode, name, tag)
			m.resizeMatchList()
			return m, nil

		case key.Matches(msg, keys.MatchPageBinding):
			var existingMatches []valorantapi.Match
			for _, item := range m.state.matchListPage.list.Items() {
				if match, ok := item.(Match); ok {
					existingMatches = append(existingMatches, valorantapi.Match(match))
				}
			}

			updatedMatches := valorantapi.AppendMatchList(existingMatches, strconv.Itoa(matchpage), "eu", m.mode)
			matchpage++

			newItems := make([]list.Item, len(updatedMatches))
			for i, match := range updatedMatches {
				newItems[i] = Match(match)
			}

			m.state.matchListPage.list.SetItems(newItems)
			return m, nil
		}
		switch msg.String() {
		case "ctrl+f":
			var username valorantapi.Username
			jsonthings.ReadData("data.json", &username)

			player := jsonthings.Username{Name: username.Name, Tag: username.Tag}
			params := jsonthings.WriteFavoriteParams{Player: player}
			jsonthings.WriteFavoriteData(params)
			m.state.searchPage = InitialModel()
		case "alt+1":
			m.loadFavorite(0)
			return m, nil
		case "alt+2":
			m.loadFavorite(1)
			return m, nil
		case "alt+3":
			m.loadFavorite(2)
			return m, nil
		case "alt+4":
			m.loadFavorite(3)
			return m, nil

		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if selectedItem, ok := m.state.matchListPage.list.SelectedItem().(Match); ok {
				m.selectedMatch = &selectedItem
				m.state.selectedMatchPage = SelectedMatchList(m.selectedMatch.Id)
				m.resizeSelectedMatch()
				m = m.SwitchPage(selectedMatchPage)
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.state.matchListPage.list, cmd = m.state.matchListPage.list.Update(msg)
	return m, cmd
}

var docStyle = lipgloss.NewStyle().
	Margin(3, 4)

func (m model) matchListView() string {
	help := shortHelpView([]key.Binding{
		keys.NavigationBindings,
		keys.SearchPageBinding,
		keys.MatchPageBinding,
		keys.PreviousModeBinding,
		keys.NextModeBinding,
		keys.QuickSwitchBinding,
	})

	base := docStyle.Render(lipgloss.JoinVertical(lipgloss.Top, m.state.matchListPage.list.View(), help))

	if m.searchOpen {
		return lipgloss.Place(
			m.state.width,
			m.state.height,
			lipgloss.Center,
			lipgloss.Center,
			m.searchPopupView(),
			lipgloss.WithWhitespaceChars(""),
			lipgloss.WithWhitespaceForeground(lipgloss.AdaptiveColor{Light: "0", Dark: "0"}),
		)
	}

	return base
}

func (m Match) Title() string {
	matchTime, err := time.Parse(time.RFC3339Nano, m.StartedAt)
	matchDay, matchMonth, matchYear, matchHour, matchMinute := matchTime.Day(), matchTime.Month(), matchTime.Year(), matchTime.Hour(), matchTime.Minute()
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}

	if m.Team == "Blue" {
		if m.BlueTeamScore > m.RedTeamScore {
			return fmt.Sprintf("Win - %d-%d %s - %s  %d %s %d %dh%d",
				m.BlueTeamScore, m.RedTeamScore, m.MapName, m.CharacterName, matchDay, matchMonth, matchYear, matchHour, matchMinute)
		}
		if m.BlueTeamScore < m.RedTeamScore {
			return fmt.Sprintf("Loss - %d-%d %s - %s  %d %s %d %dh%d",
				m.BlueTeamScore, m.RedTeamScore, m.MapName, m.CharacterName, matchDay, matchMonth, matchYear, matchHour, matchMinute)
		} else {
			return fmt.Sprintf("Draw - %d-%d %s - %s  %d %s %d %dh%d",
				m.BlueTeamScore, m.RedTeamScore, m.MapName, m.CharacterName, matchDay, matchMonth, matchYear, matchHour, matchMinute)
		}
	} else {
		if m.BlueTeamScore < m.RedTeamScore {
			return fmt.Sprintf("Win - %d-%d %s - %s  %d %s %d %dh%d",
				m.RedTeamScore, m.BlueTeamScore, m.MapName, m.CharacterName, matchDay, matchMonth, matchYear, matchHour, matchMinute)
		}
		if m.BlueTeamScore > m.RedTeamScore {
			return fmt.Sprintf("Loss- %d-%d %s - %s  %d %s %d %dh%d",
				m.RedTeamScore, m.BlueTeamScore, m.MapName, m.CharacterName, matchDay, matchMonth, matchYear, matchHour, matchMinute)
		} else {
			return fmt.Sprintf("Draw- %d-%d %s - %s  %d %s %d %dh%d",
				m.RedTeamScore, m.BlueTeamScore, m.MapName, m.CharacterName, matchDay, matchMonth, matchYear, matchHour, matchMinute)
		}
	}
}

func (m Match) Description() string {
	totalshots := m.Bodyshots + m.Headshots + m.Legshots
	headshotPercentage := float64(m.Headshots) / float64(totalshots) * 100
	return fmt.Sprintf("Kills: %d | Deaths: %d | Assists: %d |          K/D: %.1f | ACS: %d | HS: %.1f%%",
		m.Kills, m.Deaths, m.Assists, float32(m.Kills)/float32(m.Deaths), m.Score/(m.BlueTeamScore+m.RedTeamScore), headshotPercentage)
}
func (i Match) FilterValue() string { return i.StartedAt }

// resizeMatchList applies the correct list dimensions based on the stored window size.
func (m *model) resizeMatchList() {
	if m.state.width == 0 && m.state.height == 0 {
		return
	}
	vMargin := 3 * 2
	hMargin := 4 * 2
	helpView := shortHelpView([]key.Binding{
		keys.NavigationBindings,
		keys.SearchPageBinding,
		keys.MatchPageBinding,
		keys.PreviousModeBinding,
		keys.NextModeBinding,
		keys.FavoriteBinding,
	})
	helpHeight := lipgloss.Height(helpView)
	m.state.matchListPage.list.SetSize(m.state.width-hMargin, m.state.height-vMargin-helpHeight)
}

func MatchList(name string, tag string, mode string) matchListState {
	puuid := valorantapi.GetAccountPUUID(name, tag)
	matchList := valorantapi.GetAccountMatches(puuid, "1", "eu", mode)

	items := make([]list.Item, len(matchList))
	for i := range len(matchList) {
		items[i] = Match(matchList[i])
	}

	windowWidth, windowHeight := 0, 0
	
	l := list.New(items, matchlistDelegate{}, windowWidth, windowHeight)
	l.Title = fmt.Sprintf("%s: %s#%s", strings.ToUpper(mode[:1])+mode[1:], name, tag)

	l.SetShowHelp(false)
	m := matchListState{list: l}
	return m
}
