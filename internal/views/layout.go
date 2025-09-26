package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kalelc/movies/internal/domain"
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
		l.content.SetData(selected)
	}

	return l, cmd
}

func (l Layout) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, l.sidebar.View(), l.content.View())
}
