package main

import (
	"github.com/vinihss/telegrandma/cmd"
	"github.com/vinihss/telegrandma/internal/application"
)

var app = application.NewApplication(
	"telegrandma",
	"telegrandma - Remote Server Manager CLI",
)

func main() {

	err := app.Run()
	if err != nil {
		return
	}
	app.SetProcess(func() {

		cmd.Execute()
	})

}
