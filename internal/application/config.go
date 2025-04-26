package application

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// AppConfig mantém as configurações de caminhos da aplicação
type AppConfig struct {
	ConfigDir  string
	ConfigFile string
	LogDir     string
	LogFile    string
}

// getAppName retorna o nome da aplicação para usar nos caminhos
func getAppName() string {
	return "myapp" // Substitua pelo nome real da sua aplicação
}

// getDefaultConfig retorna a configuração padrão baseada no SO
func getDefaultConfig() (*AppConfig, error) {
	appName := getAppName()
	var configDir, logDir string

	// Determina os diretórios base conforme o sistema operacional
	switch runtime.GOOS {
	case "windows":
		configDir = filepath.Join(os.Getenv("APPDATA"), appName)
		logDir = filepath.Join(os.Getenv("LOCALAPPDATA"), appName, "logs")
	case "darwin": // macOS
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("falha ao obter diretório home: %w", err)
		}
		configDir = filepath.Join(home, "Library", "Application Support", appName)
		logDir = filepath.Join(home, "Library", "Logs", appName)
	default: // Linux e outros Unix-like
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("falha ao obter diretório home: %w", err)
		}
		configDir = filepath.Join(home, ".config", appName)
		logDir = filepath.Join(home, ".local", "share", appName, "logs")
	}

	return &AppConfig{
		ConfigDir:  configDir,
		ConfigFile: filepath.Join(configDir, "config.json"),
		LogDir:     logDir,
		LogFile:    filepath.Join(logDir, appName+".log"),
	}, nil
}

// setupDirectories cria os diretórios necessários
func setupDirectories(cfg *AppConfig) error {
	// Cria diretório de configuração
	if err := os.MkdirAll(cfg.ConfigDir, 0755); err != nil {
		return fmt.Errorf("falha ao criar diretório de configuração: %w", err)
	}

	// Cria diretório de logs
	if err := os.MkdirAll(cfg.LogDir, 0755); err != nil {
		return fmt.Errorf("falha ao criar diretório de logs: %w", err)
	}

	return nil
}

// initLogger configura o sistema de logging
func initLogger(cfg *AppConfig) (*os.File, error) {
	logFile, err := os.OpenFile(cfg.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir arquivo de log: %w", err)
	}

	// Configura saída padrão do logger
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	return logFile, nil
}

func init() {
	// Obtém configuração padrão baseada no SO
	cfg, err := getDefaultConfig()
	if err != nil {
		log.Fatalf("Erro ao obter configuração: %v", err)
	}

	// Cria diretórios necessários
	if err := setupDirectories(cfg); err != nil {
		log.Fatalf("Erro ao configurar diretórios: %v", err)
	}

	// Inicializa logger
	logFile, err := initLogger(cfg)
	if err != nil {
		log.Fatalf("Erro ao configurar logger: %v", err)
	}
	defer logFile.Close()

}
