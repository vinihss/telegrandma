package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"telegrandma/internal/core"
)

var errChan chan error

var rootCmd = &cobra.Command{
	Use:   "motherbot",
	Short: "Manage telegram bots",
	Long:  `manage telegram bots.`,
}

func Execute() {

	//errChan = make(chan error)
	//InitializeErrorHandler()
	if err := rootCmd.Execute(); err != nil {
		core.LogError(fmt.Errorf("error executing %s script: %v", err))
		os.Exit(1)
	}

}

func InitializeErrorHandler() {
	// Inicie a gorotina e passe o canal de erros
	//errChan <- fmt.Errorf("ocorreu um erro na gorotina")

	// Leia os erros do canal
	select {
	case err := <-errChan:
		if err != nil {
			fmt.Printf("Erro recebido da gorotina: %v\n", err)
		}
	default:
		fmt.Printf("Erro nao: ")

	}
}
