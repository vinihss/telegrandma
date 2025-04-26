package core

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

var chatChannels = make(map[int64]chan tgbotapi.Update)
var mu sync.Mutex
var wg sync.WaitGroup

const numWorkers = 5 // Número de goroutines simultâneas
type task struct {
	name   string
	action func(update tgbotapi.Update)
}

// Logger instance for the application
var logger *log.Logger

func init() {
	// Initialize the logger
	file, err := os.OpenFile("motherbot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error creating log file: %v\n", err)
		return
	}
	logger = log.New(file, "motherbot: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(message string) {
	if logger != nil {
		logger.Println("INFO: " + message)
	}
}

func LogError(err error) {
	if logger != nil && err != nil {
		logger.Println("ERROR: " + err.Error())
	}
}

func LogDebug(message string) {
	if logger != nil {
		logger.Println("DEBUG: " + message)
	}
}

func worker(task *task) {
	defer wg.Done()
	for {
		mu.Lock()
		for _, ch := range chatChannels {
			select {
			case update := <-ch:
				mu.Unlock()
				task.action(update)

				mu.Lock()
			default:
				continue
			}
		}
		mu.Unlock()
		time.Sleep(100 * time.Millisecond) // Evita loop ocupado
	}
}

func ExecCommand(text string) (string, error) {
	var answer string

	cmd := exec.Command("/usr/bin/tgpt", "--provider", "phind", "-w", fmt.Sprintf("%v", text), "")
	output, err := cmd.CombinedOutput()
	if err != nil {
		answer = "Desculpe, estou com alguns problemas. Tente novamente mais tarde"
		return answer, err
	}
	answer = string(output)

	return answer, nil
}
