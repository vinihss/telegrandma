package system_checker

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"log"
	"telegrandma/pkg/utils/logger"
)

var disks = Disks{}

const DISK_THRESHOLD = 10.0 // % of free disk space
type DiskUsage struct {
	MountPoint  string  `json:"mountPoint"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}

type Disks struct {
	Disks []DiskUsage `json:"disks"`
}

func Diagnostic(handler func()) {
	CheckDisksUsage()
	handler()
}
func CheckDisksUsage() {
	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Printf("Error getting disk partitions: %v", err)
		return
	}

	for _, part := range partitions {
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			log.Printf("Error getting disk usage (%s): %v", part.Mountpoint, err)
			continue
		}
		diskUsage := DiskUsage{part.Mountpoint, usage.Total, usage.Used, usage.Free, usage.UsedPercent}
		disks.Disks = append(disks.Disks, diskUsage)
		Diagnostic(func() {

			if diskUsage.UsedPercent > (100 - DISK_THRESHOLD) {
				logger.CreateLog(fmt.Sprintf("Critical disk space on %s: %.2f%% used", part.Mountpoint, usage.UsedPercent), "CRITICAL")
			}
		})
	}
}

func GetDisks() Disks {
	return disks
}
