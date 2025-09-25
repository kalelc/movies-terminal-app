package views

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kalelc/movies/internal/domain"
	"golang.org/x/term"
)

type Layout struct {
	sidebar Sidebar
	content Content
}

func NewLayout(sidebar Sidebar, content Content) Layout {
	return Layout{
		sidebar: sidebar,
		content: content,
	}
}

func (l Layout) Init() tea.Cmd {
	return nil
}

func (l Layout) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var selected *domain.Movie

	l.sidebar, cmd, selected = l.sidebar.Update(msg)
	l.content, _ = l.content.Update(msg)

	if selected != nil {
		l.content.SetData(selected.Name, selected.Overview)
	}

	return l, cmd
}

func (l Layout) View() string {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80
		height = 40
	}

	col1Width := int(float64(width) * 0.30)
	col2Width := (width - 5) - col1Width
	rowHeight := height - 2

	col1Style := lipgloss.NewStyle()
	renderedCol1 := col1Style.Render(l.sidebar.View())

	col2Style := lipgloss.NewStyle().
		Width(col2Width).
		Height(rowHeight).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(0, 1)

	renderedCol2 := col2Style.Render(l.content.View())

	return lipgloss.JoinHorizontal(lipgloss.Top, renderedCol1, renderedCol2)
}
