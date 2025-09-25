package views

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kalelc/movies/internal/domain"
	"github.com/kalelc/movies/internal/services"
	"golang.org/x/term"
)

type Model struct {
	list   list.Model
	detail string
}

func NewModel(s *services.TmdbService) Model {

	movies := s.GetMovies()
	items := make([]list.Item, len(movies))
	for i, m := range movies {
		items[i] = m
	}

	const defaultWidth = 40
	const defaultHeight = 40
	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, defaultHeight)
	l.Title = "Menú"

	return Model{list: l}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if sel, ok := m.list.SelectedItem().(domain.Movie); ok {
				m.detail = fmt.Sprintf("%s\n%s\n%s", sel.Name, sel.Overview, sel.ReleaseDate)
			}
		}
	}
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80
		height = 40
	}

	col1Width := int(float64(width) * 0.30)
	col2Width := (width - 5) - col1Width

	rowHeight := height - 2

	col1Style := lipgloss.NewStyle().
		Width(col1Width).
		Height(rowHeight).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(0, 1)

	col2Style := lipgloss.NewStyle().
		Width(col2Width).
		Height(rowHeight).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(0, 1)

	renderedCol1 := col1Style.Render(m.list.View())
	renderedCol2 := col2Style.Render(m.detail)

	return lipgloss.JoinHorizontal(lipgloss.Top, renderedCol1, renderedCol2)
}
