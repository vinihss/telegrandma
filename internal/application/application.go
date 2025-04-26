package application

import (
	"fmt"

	"telegrandma/pkg/cmd/cli"
)

type Application struct {
	Name        string
	Description string
	Version     string
	AppConfig   AppConfig
}

// PrintASCIIArt prints an ASCII art during application startup
func (app *Application) PrintASCIIArt() {
	fmt.Println(`
   _____  _	 _  _____
  / ____|| |   (_)/ ____|
 | (___  | |__  _| |  __
  \___ \ | '_ \| | | |_ |
  ____) || | | | | |__| |
 |_____/ |_| |_|_|\_____|

--------------------------------
ShipX - Remote Server Manager CLI
--------------------------------
	`)
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
