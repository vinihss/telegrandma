package system_checker

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"strings"
	"telegrandma/pkg/notify"
	"telegrandma/pkg/utils"
	"time"
)

// Limites críticos
const (
	CPU_THRESHOLD = 90.0 // % de uso da CPU
)

// Obtém os 5 princip

// Verifica o uso da CPU
func CheckCPU() {
	usages, err := cpu.Percent(time.Second, true)
	if err != nil {

		LogEvent(EventFactory("", "Erro ao obter uso da CPU: %v", ""))

		return
	}
	messages := make([]string, 0)
	cpu_count := len(usages)
	cpu_bussy := 0
	for i, usage := range usages {
		if usage > CPU_THRESHOLD {
			cpu_bussy++
			processInfo := utils.GetTopProcesses()
			alertMessage := fmt.Sprintf("Uso de CPU crítico no núcleo %d: %.2f%%\n", i, usage)
			for _, p := range processInfo {
				alertMessage += fmt.Sprintf("%s (%.2f%% CPU, %.2fMB RAM)\n", p.Name, p.CPUPercent, p.MemUsage)
			}
			messages = append(messages, alertMessage)
		}
	}
	if cpu_bussy >= cpu_count {

		notify.SendTelegram(strings.Join(messages, "\n"))

	}
}
