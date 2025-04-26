package system_checker

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"telegrandma/pkg/notify"
	"telegrandma/pkg/utils/logger"
	"time"
)

const (
	LOG_MODE = 1
)

var checkupPipeline CheckupPipeline
var mode = 1
var checkHealthTime = 1 * time.Second
var checkHealthAnalyzer = 1 * time.Minute

func (pipeline CheckupPipeline) EnableSystemAlert() {
	checkupPipeline.systemAlert = true
}

func (pipeline CheckupPipeline) DisableSystemAlert() {
	log.Println("INFO: System alert disabled")
	checkupPipeline.systemAlert = false
}
func (pipeline CheckupPipeline) AddRoutine(routine func()) []func() {
	return append(pipeline.Routines, routine)
}

type CheckupPipeline struct {
	systemAlert bool
	status      bool
	Routines    []func()
}

func (pipeline CheckupPipeline) getStatus() bool {
	return pipeline.status
}
func (pipeline CheckupPipeline) Check() {
	log.Println("Checking system alert")
	for _, routine := range pipeline.Routines {
		routine()
	}
}

func ShowConfig() {

}

var (
	runningPipelines []CheckupPipeline
	mu               sync.Mutex
)

func AddPipeline(pipeline CheckupPipeline) {
	mu.Lock()
	defer mu.Unlock()
	runningPipelines = append(runningPipelines, pipeline)
}

func GetRunningPipelines() []CheckupPipeline {
	mu.Lock()
	defer mu.Unlock()
	return runningPipelines
}

func StartUnixSocketServer() {
	log.Println("Starting Unix Socket Server")
	socketPath := "/tmp/system_checker.sock"
	os.Remove(socketPath)

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Println("ERROR: error starting unix socket server")
		panic(err)
	}

	defer listener.Close()
	log.Println("INFO: Application started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("ERROR: error accepting connection")
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	mu.Lock()
	mu.Unlock()
	data, err := json.Marshal(runningPipelines)
	if err != nil {
		return
	}
	conn.Write(data)
}
func buildDefaultPipeline() CheckupPipeline {
	checkupPipeline = CheckupPipeline{}
	checkupPipeline.AddRoutine(CheckDisksUsage)
	checkupPipeline.AddRoutine(CheckMemory)
	checkupPipeline.AddRoutine(CheckCPU)
	checkupPipeline.AddRoutine(CheckNetwork)
	checkupPipeline.AddRoutine(CheckLogs)
	log.Println("INFO: Building default pipeline")

	return checkupPipeline
}

func GetPipeline() CheckupPipeline {
	log.Println("INFO: Getting pipeline")

	if checkupPipeline.Routines != nil {
		log.Println("INFO: Returning default pipeline existing")

		return checkupPipeline
	}

	return buildDefaultPipeline()

}
func (pipeline CheckupPipeline) Start() {
	log.Println("Monitor de recursos iniciado...")
	ticker := time.NewTicker(checkHealthTime)
	logTicker := time.NewTicker(checkHealthAnalyzer) // Ajuste o intervalo conforme necessário

	defer ticker.Stop()
	defer logTicker.Stop()

	for {
		select {
		case <-ticker.C:
			checkupPipeline.Check()
		case <-logTicker.C:
			ReadLogsAndSendAlert()
		}
	}
}

func (pipeline CheckupPipeline) IsRunning() bool {
	return checkupPipeline.getStatus() == true
}

func init() {
	checkupPipeline = GetPipeline()
}

func ReadLogsAndSendAlert() {
	log.Println("Reading logs from log table and sending alert")

	// Assuming you have a function to get logs from the log table
	recentLogs, err := logger.GetRecentLogs(10)
	if err != nil {
		log.Printf("Erro ao obter logs recentes: %v", err)
		return
	}

	var logs string
	for _, logEntry := range recentLogs {
		logs += fmt.Sprintf("%s - %s\n", logEntry.Level, logEntry.Message)
	}

	notifyEvent(fmt.Sprintf("Logs do sistema:\n%s", logs))

	for _, logEntry := range recentLogs {
		log.Printf("Log: %s - %s", logEntry.Level, logEntry.Message)
	}
}

func notifyEvent(log string) {

	if checkupPipeline.systemAlert {
		fmt.Println("system alert")

	}

	notify.SendTelegram(log)

}
