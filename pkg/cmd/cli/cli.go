package cli

import "github.com/spf13/cobra"

var cli *Cli

type Cli struct {
	RootCommand *cobra.Command
}

func (c *Cli) Execute() error {
	return c.RootCommand.Execute()
}

func NewCli() *Cli {
	cli := &Cli{}
	cli.RootCommand = &cobra.Command{
		Use:   "motherbot",
		Short: "Manage telegram bots",
		Long:  `manage telegram bots.`,
	}
	return cli
}

func AddCommand(c *cobra.Command) {
	cli.RootCommand.AddCommand(c)

}

func init() {
	cli = NewCli()
}

func GetCli() *Cli {
	return cli
}
