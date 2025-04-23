package tui

import (
	"fmt"
	_ "fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Heribio/ValTracker/internal/jsonthings"
	_ "github.com/Heribio/ValTracker/internal/jsonthings"
	"github.com/Heribio/ValTracker/internal/valorantapi"
)

type (
	errMsg error
)

const (
	name = iota
	tag
)

type searchState struct {
    inputs []textinput.Model
    favorites []jsonthings.Username
    focused   int
    err       error
}

func InitialModel() searchState {
    favorites := jsonthings.GetFavoriteData().Favorites
    inputs := make([]textinput.Model, 2)

	inputs[name] = textinput.New()
	inputs[name].Placeholder = "Name"
	inputs[name].Focus()
	inputs[name].CharLimit = 156
	inputs[name].Width = 20

	inputs[tag] = textinput.New()
	inputs[tag].Placeholder = "Tag"
	inputs[tag].CharLimit = 6
	inputs[tag].Width = 20

	return searchState{
		inputs: inputs,
        favorites: favorites,
		focused: 0,
		err:       nil,
	}
}

func (m model) searchUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
    inputs := m.state.searchPage.inputs
    focusedEntry := m.state.searchPage.focused
    var cmds []tea.Cmd = make([]tea.Cmd, len(inputs))

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "enter":
            if focusedEntry == len(inputs)-1 {
                name := inputs[name].Value()
                tag := inputs[tag].Value()
                username := valorantapi.Username{
                    Name: name,
                    Tag: tag,
                }
                jsonthings.WriteData("data.json", username)

                m.state.matchListPage = MatchList(name, tag, m.mode)
                return m.SwitchPage(matchListPage), nil
            }
            m.nextInput()
            return m, nil
        case "ctrl+c":
            return m, tea.Quit
        case "tab", "down":
            m.nextInput()
        case "shift+tab", "up":
            m.prevInput()
        case "esc":
            m = m.SwitchPage(matchListPage)
            return m, nil
        case "ctrl+h":
            favoriteData := jsonthings.GetFavoriteData()
            player := favoriteData.Favorites[0]
            username := valorantapi.Username{
                    Name: player.Name,
                    Tag: player.Tag,
                }
            jsonthings.WriteData("data.json", username)

            m.state.matchListPage = MatchList(player.Name, player.Tag, m.mode)
            return m.SwitchPage(matchListPage), nil
        case "ctrl+j":
            favoriteData := jsonthings.GetFavoriteData()
            player := favoriteData.Favorites[1]
            
            username := valorantapi.Username{
                    Name: player.Name,
                    Tag: player.Tag,
                }
            jsonthings.WriteData("data.json", username)

            m.state.matchListPage = MatchList(player.Name, player.Tag, m.mode)
            return m.SwitchPage(matchListPage), nil
        case "ctrl+k":
            favoriteData := jsonthings.GetFavoriteData()
            player := favoriteData.Favorites[2]
            username := valorantapi.Username{
                    Name: player.Name,
                    Tag: player.Tag,
                }
            jsonthings.WriteData("data.json", username)

            m.state.matchListPage = MatchList(player.Name, player.Tag, m.mode)
            return m.SwitchPage(matchListPage), nil
        case "ctrl+l":
            favoriteData := jsonthings.GetFavoriteData()
            player := favoriteData.Favorites[3]
            username := valorantapi.Username{
                    Name: player.Name,
                    Tag: player.Tag,
                }
            jsonthings.WriteData("data.json", username)

            m.state.matchListPage = MatchList(player.Name, player.Tag, m.mode)
            return m.SwitchPage(matchListPage), nil

        }


        for i := range inputs {
            inputs[i].Blur()
        }
        inputs[m.state.searchPage.focused].Focus()
    }

    for i := range inputs {
        inputs[i], cmds[i] = inputs[i].Update(msg)
    }
    return m, tea.Batch(cmds...)
}

func (m model) searchView() string {
    inputs := m.state.searchPage.inputs
    favorites := m.state.searchPage.favorites
    favoriteView  := "Favorites: \n"
    for i := range favorites {
        favoriteView = favoriteView + fmt.Sprintf("%d: %s#%s\n", i, favorites[i].Name, favorites[i].Tag)
    }
    help := shortHelpView([]key.Binding{
        keys.PreviousInputBinding,
        keys.NextInputBinding,
        keys.ConfirmBinding,
        keys.MatchListBinding,
    })

    inputView := fmt.Sprintf(
		"Insert the name and tag of the valorant player\n\n%s\n\n%s",
		inputs[name].View(),
		inputs[tag].View(),
	) + "\n\n"

    return docStyle.Render(lipgloss.JoinVertical(lipgloss.Top,
        inputView,
        favoriteView,
        help))
}

func (m *model) nextInput() {
    focusedEntry := m.state.searchPage.focused
    inputs := m.state.searchPage.inputs

	m.state.searchPage.focused = (focusedEntry + 1) % len(inputs)
}

func (m *model) prevInput() {
    inputs := m.state.searchPage.inputs

	m.state.searchPage.focused--
	if m.state.searchPage.focused < 0 {
		m.state.searchPage.focused = len(inputs) - 1
	}
}
