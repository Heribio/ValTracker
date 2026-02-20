package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Heribio/ValTracker/internal/jsonthings"
	"github.com/Heribio/ValTracker/internal/valorantapi"
)

type (
	errMsg error
)

type searchState struct {
	input     textinput.Model
	favorites []jsonthings.Username
	err       error
}

func InitialModel() searchState {
	favorites := jsonthings.GetFavoriteData().Favorites

	input := textinput.New()
	input.Placeholder = "Name#Tag  (e.g. PlayerName#1234)"
	input.Focus()
	input.CharLimit = 163
	input.Width = 36

	return searchState{
		input:     input,
		favorites: favorites,
		err:       nil,
	}
}

// loadFavorite switches to the match list for the nth favorite (0-indexed).
// Returns false if the favorite index is out of range.
func (m *model) loadFavorite(index int) bool {
	favoriteData := jsonthings.GetFavoriteData()
	if index < 0 || index >= len(favoriteData.Favorites) {
		return false
	}
	player := favoriteData.Favorites[index]
	username := valorantapi.Username{Name: player.Name, Tag: player.Tag}
	jsonthings.WriteData("data.json", username)
	m.state.matchListPage = MatchList(player.Name, player.Tag, m.mode)
	m.resizeMatchList()
	return true
}

func (m model) searchUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			raw := m.state.searchPage.input.Value()
			parts := strings.SplitN(raw, "#", 2)
			if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
				m.state.searchPage.err = fmt.Errorf("enter player as Name#Tag")
				return m, nil
			}
			playerName, playerTag := parts[0], parts[1]
			username := valorantapi.Username{Name: playerName, Tag: playerTag}
			jsonthings.WriteData("data.json", username)
			m.state.matchListPage = MatchList(playerName, playerTag, m.mode)
			m.resizeMatchList()
			return m.SwitchPage(matchListPage), nil
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return m.SwitchPage(matchListPage), nil
		case "alt+1":
			if m.loadFavorite(0) {
				return m.SwitchPage(matchListPage), nil
			}
		case "alt+2":
			if m.loadFavorite(1) {
				return m.SwitchPage(matchListPage), nil
			}
		case "alt+3":
			if m.loadFavorite(2) {
				return m.SwitchPage(matchListPage), nil
			}
		case "alt+4":
			if m.loadFavorite(3) {
				return m.SwitchPage(matchListPage), nil
			}
		}
	}

	m.state.searchPage.input, cmd = m.state.searchPage.input.Update(msg)
	return m, cmd
}

var searchShortcutKeys = []string{"alt+1", "alt+2", "alt+3", "alt+4"}

func (m model) searchView() string {
	favorites := m.state.searchPage.favorites
	input := m.state.searchPage.input

	// Build the input box
	inputBox := searchInputStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			"Search for a Valorant player",
			"",
			input.View(),
		),
	)

	// Build error line (empty if no error)
	errLine := ""
	if m.state.searchPage.err != nil {
		errLine = lipgloss.NewStyle().Foreground(lipgloss.Color("#e4485d")).Render(
			"  ✗ "+m.state.searchPage.err.Error(),
		) + "\n"
	}

	// Build favorites list
	favoriteView := "Favorites:\n"
	if len(favorites) == 0 {
		favoriteView += "  (none saved)\n"
	} else {
		for i, fav := range favorites {
			if i >= len(searchShortcutKeys) {
				break
			}
			shortcut := helpKeyStyle("[" + searchShortcutKeys[i] + "]")
			entry := helpDescStyle(fmt.Sprintf("  %s#%s", fav.Name, fav.Tag))
			favoriteView += shortcut + entry + "\n"
		}
	}

	help := shortHelpView([]key.Binding{
		keys.ConfirmBinding,
		keys.MatchListBinding,
		keys.QuickSwitchBinding,
	})

	return docStyle.Render(lipgloss.JoinVertical(lipgloss.Top,
		inputBox,
		errLine,
		favoriteView,
		"",
		help,
	))
}
