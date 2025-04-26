package system_checker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Caminhos dos principais logs do Arch Linux
var logFiles = map[string]string{
	"pacman": "/var/log/pacman.log",
	"syslog": "/var/log/syslog",
	"auth":   "/var/log/auth.log",
	"boot":   "/var/log/boot.log",
	"xorg":   "/var/log/Xorg.0.log",
}

// Estrutura para representar uma entrada de log
type LogEntry struct {
	Timestamp string
	Message   string
}

func CheckLogs() { // Verifica os logs
	for logName, logPath := range logFiles {
		fmt.Printf("=== Verificando %s (%s) ===\n", logName, logPath)
		_, err := readRecentLogs(logPath, 5) // Lê as últimas 5 entradas
		if err != nil {
			//	log.Printf("Erro ao ler %s: %v\n", logPath, err)
			continue
		}

		//fmt.Println()
	}

}

// Lê as últimas entradas de um arquivo de log
func readRecentLogs(path string, limit int) ([]LogEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []LogEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Extrai o timestamp e a mensagem (depende do formato do log)
		timestamp, message := parseLogLine(line)
		entries = append(entries, LogEntry{
			Timestamp: timestamp,
			Message:   message,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Retorna apenas as últimas `limit` entradas
	if len(entries) > limit {
		entries = entries[len(entries)-limit:]
	}
	return entries, nil
}

// Extrai o timestamp e a mensagem de uma linha de log
func parseLogLine(line string) (string, string) {
	// Formato de log comum: "Mês Dia Hora:Minuto:Segundo Hostname Mensagem"
	parts := strings.SplitN(line, " ", 4)
	if len(parts) < 4 {
		return time.Now().Format("2006-01-02 15:04:05"), line
	}

	timestamp := fmt.Sprintf("%s %s %s", parts[0], parts[1], parts[2])
	message := parts[3]
	return timestamp, message
}
