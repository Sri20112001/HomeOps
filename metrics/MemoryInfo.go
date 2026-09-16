package metrics

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

type MemoryInfo struct {
	Total       uint64  `json:"total_bytes"`
	Available   uint64  `json:"available_bytes"`
	Used        uint64  `json:"used_bytes"`
	UsedPercent float64 `json:"used_percent"`
}

func GetMemoryInfo() (MemoryInfo, error) {

	memoryInfo, err := mem.VirtualMemory()
	if err != nil {
		return MemoryInfo{},
			fmt.Errorf("get memory info: %w", err)
	}

	return MemoryInfo{
		Total:       memoryInfo.Total,
		Available:   memoryInfo.Available,
		Used:        memoryInfo.Used,
		UsedPercent: memoryInfo.UsedPercent,
	}, nil
}
