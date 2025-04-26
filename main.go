package main

import (
	"telegrandma/cmd"
	"telegrandma/internal/application"
)

var app = application.NewApplication("telegrandma", "telegrandma - Remote Server Manager CLI")

func main() {
	app.Run()
	app.SetProcess(func() {
		cmd.cmd.Execute()
	})

}
