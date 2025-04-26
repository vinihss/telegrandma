package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vinihss/telegrandma/internal/conversation"
	"github.com/vinihss/telegrandma/pkg/cmd/cli"
)

type CommandSettings struct {
	Name             string
	Use              string
	ShortDescription string
	RunAction        func()
}

var chatCmd = cli.BaseCommand("chat", "Manage chatbots")

var start = &cobra.Command{
	Use:   "start",
	Short: "Run a bot",
	Run: func(cmd *cobra.Command, args []string) {
		conversation.InitializeBot()
	},
}

func init() {

	chatCmd.AddCommand(start)
	rootCmd.AddCommand(chatCmd)
}
