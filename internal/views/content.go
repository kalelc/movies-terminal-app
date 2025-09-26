package views

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kalelc/movies/internal/domain"
	"golang.org/x/term"
)

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b)
	}()
)

type Content struct {
	viewport viewport.Model
	title    string
	body     string
}

func NewContent() Content {
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))
	width = int(float64(width)*0.70) - 5
	height = height - 2

	vp := viewport.New(width, height)
	vp.SetContent("Selecciona una película para ver los detalles")
	return Content{viewport: vp}
}

func (c *Content) SetData(movie *domain.Movie) {
	c.title = movie.Name
	c.body = movie.Overview

	bodyStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#7D56F4")).
		Width(c.viewport.Width).Render(c.body)

	content := fmt.Sprintf(
		"%s\n%s\n%s",
		c.headerView(),
		bodyStyle,
		c.footerView(),
	)

	c.viewport.SetContent(content)
}

func (c *Content) Update(msg tea.Msg) (Content, tea.Cmd) {
	var cmd tea.Cmd
	c.viewport, cmd = c.viewport.Update(msg)
	return *c, cmd
}

func (c Content) View() string {
	return c.viewport.View()
}

func (c *Content) headerView() string {
	title := titleStyle.Render(c.title)
	line := strings.Repeat("─", max(0, c.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (c *Content) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", c.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, c.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}
