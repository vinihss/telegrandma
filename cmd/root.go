package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vinihss/telegrandma/internal/core"

	"os"
)

var errChan chan error

var rootCmd = &cobra.Command{
	Use:   "motherbot",
	Short: "Manage telegram bots",
	Long:  `manage telegram bots.`,
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		core.LogError(fmt.Errorf("error executing %s script: %v", err))
		os.Exit(1)
	}

}
