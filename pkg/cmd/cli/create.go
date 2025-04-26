package cli

import (
	"github.com/spf13/cobra"
)

type CommandSettings struct {
	Name             string
	Use              string
	ShortDescription string
	RunAction        func()
}

func BaseCommand(use string, shortDescription string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: shortDescription,
	}
}
func NewCommand(commandSettings CommandSettings) *cobra.Command {

	if commandSettings.RunAction == nil {
		commandSettings.RunAction = func() {}
	}

	return &cobra.Command{
		Use:   commandSettings.Use,
		Short: commandSettings.ShortDescription,
		Run: func(cmd *cobra.Command, args []string) {
			commandSettings.RunAction()
		},
	}

}
