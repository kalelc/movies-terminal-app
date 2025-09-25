package views

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kalelc/movies/internal/domain"
	"github.com/kalelc/movies/internal/services"
	"golang.org/x/term"
)

type Sidebar struct {
	List list.Model
}

func NewSidebar(s *services.TmdbService) Sidebar {
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))
	width = int(float64(width)*0.30) - 10
	height = height - 2

	movies := s.GetMovies()
	items := make([]list.Item, len(movies))
	for i, m := range movies {
		items[i] = m
	}

	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = "Menú"

	return Sidebar{List: l}
}

func (s Sidebar) Init() tea.Cmd {
	return nil
}

func (s Sidebar) Update(msg tea.Msg) (Sidebar, tea.Cmd, *domain.Movie) {
	var cmd tea.Cmd
	var selected *domain.Movie

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			if sel, ok := s.List.SelectedItem().(domain.Movie); ok {
				selected = &sel
			}
		}
	}

	s.List, cmd = s.List.Update(msg)
	return s, cmd, selected
}

func (s Sidebar) View() string {
	return s.List.View()
}
