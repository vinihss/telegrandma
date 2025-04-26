package system_checker

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"log"
	"telegrandma/pkg/utils"
	"time"
)

// Limites críticos
const (
	NETWORK_THRESHOLD = 1000 // KB/s (limite arbitrário)
)

// Verifica o tráfego de rede
func CheckNetwork() {
	counters1, err := net.IOCounters(true)
	if err != nil {
		log.Printf("Erro ao obter estatísticas de rede: %v", err)
		return
	}
	time.Sleep(time.Second)

	counters2, err := net.IOCounters(true)
	if err != nil {
		log.Printf("Erro ao obter estatísticas de rede: %v", err)
		return
	}

	for i := range counters1 {
		recvRate := float64(counters2[i].BytesRecv-counters1[i].BytesRecv) / 1024
		sendRate := float64(counters2[i].BytesSent-counters1[i].BytesSent) / 1024

		if recvRate > NETWORK_THRESHOLD || sendRate > NETWORK_THRESHOLD {
			processInfo := utils.GetTopProcesses()
			alertMessage := fmt.Sprintf("Tráfego de rede crítico: Interface %s - Download: %.2f KB/s, Upload: %.2f KB/s\n",
				counters2[i].Name, recvRate, sendRate)
			for _, p := range processInfo {
				alertMessage += fmt.Sprintf("%s (%.2f%% CPU, %.2fMB RAM)\n", p.Name, p.CPUPercent, p.MemUsage)
			}
			LogEvent(EventFactory("NETWORK_USAGE", alertMessage, ""))
		}
	}
}
