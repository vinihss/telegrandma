package system_checker

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"log"
	"telegrandma/pkg/utils"
	"telegrandma/pkg/utils/logger"
)

// Limites críticos
const (
	MEMORY_THRESHOLD = 90.0 // % de uso da memória
)

// Verifica o uso da memória
func CheckMemory() {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Erro ao obter uso da memória: %v", err)
		return
	}
	if vmStat.UsedPercent > MEMORY_THRESHOLD {
		processInfo := utils.GetTopProcesses()
		alertMessage := fmt.Sprintf("Uso de memória crítico: %.2f%%\n", vmStat.UsedPercent)
		for _, p := range processInfo {
			alertMessage += fmt.Sprintf("%s (%.2f%% CPU, %.2fMB RAM)\n", p.Name, p.CPUPercent, p.MemUsage)
		}
		logger.CreateLog(alertMessage, "CRITICAL")
	}
}
