package system_checker

import (
	"fmt"
	"log"
	"os"
	"telegrandma/pkg/utils/logger"
	"time"
)

const LOG_FILE = "events.log"

type Event struct {
	Time        time.Time
	Title       string
	Description string
	Severity    string
}

func LogEvent(event Event) {
	logger.CreateLog(event.Severity, event.Description)

	//LogEventToFile(event)
}

func LogEventToFile(event Event) {
	// Abre o arquivo em modo append, cria se não existir
	file, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Erro ao abrir o arquivo de log: %v", err)
		return
	}
	defer file.Close()
	LogEvent(event)

	// Formata a mensagem do evento
	logMessage := fmt.Sprintf("%s - %s: %s (%s)\n", event.Time.Format(time.RFC3339), event.Title, event.Description, event.Severity)

	// Escreve a mensagem no arquivo
	if _, err := file.WriteString(logMessage); err != nil {
		log.Printf("Erro ao escrever no arquivo de log: %v", err)
	}
}

func EventFactory(eventType, description, severity string) Event {
	var title string
	switch eventType {
	case "CPU_USAGE":
		title = "CPU Usage"
	case "MEMORY_USAGE":
		title = "Memory Usage"
	case "DISK_USAGE":
		title = "Disk Usage"
	case "NETWORK_USAGE":
		title = "Network Usage"
	default:
		title = "Unknown Event"
	}
	return Event{
		Time:        time.Now(),
		Title:       title,
		Description: description,
		Severity:    severity,
	}
}
