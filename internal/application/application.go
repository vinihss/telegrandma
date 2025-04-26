package application

import (
	"fmt"

	"github.com/vinihss/telegrandma/pkg/cmd/cli"
)

type Application struct {
	Name        string
	Description string
	Version     string
	AppConfig   AppConfig
}

// PrintASCIIArt prints an ASCII art during application startup
func (app *Application) PrintASCIIArt() {
	fmt.Println(``)
}

// Run sets up the root command and displays ASCII art
func (app *Application) Run() error {
	app.PrintASCIIArt()

	if cli.GetCli().Execute() != nil {
		return fmt.Errorf("Error executing root command")
	}
	return nil
}

func (app *Application) SetProcess(process func()) {
	process()
}
