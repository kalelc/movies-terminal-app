package views

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type Content struct {
	viewport viewport.Model
	title    string
	desc     string
}

func NewContent() Content {
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))
	width = int(float64(width)*0.70) - 5
	height = height - 2

	vp := viewport.New(width, height)
	vp.SetContent("Selecciona una película para ver los detalles")
	return Content{viewport: vp}
}

func (c *Content) SetData(title, desc string) {
	c.title = title
	c.desc = desc

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")) // violeta

	descStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")) // blanco

	content := fmt.Sprintf(
		"%s\n\n%s",
		titleStyle.Render(c.title),
		descStyle.Render(c.desc),
	)

	c.viewport.SetContent(content)
}

func (c *Content) Resize(width, height int) {
	c.viewport.Width = width
	c.viewport.Height = height
}

func (c *Content) Update(msg tea.Msg) (Content, tea.Cmd) {
	var cmd tea.Cmd
	c.viewport, cmd = c.viewport.Update(msg)
	return *c, cmd
}

func (c Content) View() string {
	return c.viewport.View()
}
