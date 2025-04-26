package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"telegrandma/internal/conversation"
	"telegrandma/pkg/cmd/cli"
)

type CommandSettings struct {
	Name             string
	Use              string
	ShortDescription string
	RunAction        func()
}

var AvailableAgents = []string{"gpt3"}
var chatCmd = cli.BaseCommand("chat", "Manage chatbots")

var start = &cobra.Command{
	Use:   "start",
	Short: "Run a bot",
	Run: func(cmd *cobra.Command, args []string) {
		conversation.InitializeBot()
	},
}
var chatList = &cobra.Command{
	Use:   "list",
	Short: "List all available agents",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("implements lost")

	},
}

var setup = &cobra.Command{
	Use:   "init",
	Short: "Init chatbot settings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("implements lost")
	},
}

func init() {

	chatCmd.AddCommand(setup)
	chatCmd.AddCommand(start)
	rootCmd.AddCommand(chatCmd)
}
