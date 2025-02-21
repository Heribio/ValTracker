package tui

import(
    "io"
    "fmt"
    "strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


var(
    winStyle = lipgloss.NewStyle().
        Border(lipgloss.ThickBorder(), false, false, false, true).
        BorderForeground(lipgloss.Color("#5ee790")).
        Bold(true).
        Padding(0, 2)
    lossStyle = lipgloss.NewStyle().
        Border(lipgloss.ThickBorder(), false, false, false, true).
        BorderForeground(lipgloss.Color("#e4485d")).
        Bold(true).
        Padding(0, 2)
    selectedStyle = lipgloss.NewStyle().
        Bold(true).
        Padding(0, 2).
        Foreground(lipgloss.Color("#2F6D90"))
    descStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("8")).
        PaddingLeft(4)
)

type matchlistDelegate struct{}

func (d matchlistDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
    //THAT IS TYPE ASSERTION
	it, ok := listItem.(Match)
	if !ok {
		return
	}

	title := it.Title()
    description := it.Description()
	var style lipgloss.Style

    if strings.Contains(strings.ToLower(title), "win") {
		style = winStyle
    } else {
		style = lossStyle
	}
	if index == m.Index(){
        style = selectedStyle
	} 
	fmt.Fprintf(w, "%s\n%s\n", style.Render(title), descStyle.Render(description))
}


func (d matchlistDelegate) Height() int  { return 3 }
func (d matchlistDelegate) Spacing() int { return 0 }
func (d matchlistDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd{
	return nil
}

