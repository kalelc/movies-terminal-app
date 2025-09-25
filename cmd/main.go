package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kalelc/movies/internal/services"
	"github.com/kalelc/movies/internal/views"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			services.NewTmdbService,
			views.NewModel,
		),
		fx.Invoke(runUI),
	)
	app.Run()
}

func runUI(m views.Model) {
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
