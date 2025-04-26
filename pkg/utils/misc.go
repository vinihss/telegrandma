package utils

import (
	"github.com/shirou/gopsutil/process"
	"sort"
)

// Processo com consumo alto
type HighUsageProcess struct {
	Name       string
	Path       string
	CPUPercent float64
	MemUsage   float32
	NetUsage   float64
}

func GetTopProcesses() []HighUsageProcess {
	procs, _ := process.Processes()
	var highUsage []HighUsageProcess

	for _, p := range procs {
		name, _ := p.Name()
		path, _ := p.Exe()
		cpuPercent, _ := p.CPUPercent()
		memInfo, _ := p.MemoryInfo()

		if cpuPercent > 0 {
			highUsage = append(highUsage, HighUsageProcess{
				Name:       name,
				Path:       path,
				CPUPercent: cpuPercent,
				MemUsage:   float32(memInfo.RSS) / (1024 * 1024), // Em MB
			})
		}
	}

	sort.Slice(highUsage, func(i, j int) bool {
		return highUsage[i].CPUPercent > highUsage[j].CPUPercent
	})

	if len(highUsage) > 5 {
		return highUsage[:5]
	}
	return highUsage
}
