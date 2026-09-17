package metrics

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/v4/process"
)

type ProcessInfo struct {
	PID        int32   `json:"pid"`
	Name       string  `json:"name"`
	MemoryRSS  uint64  `json:"memory_rss_bytes"`
	CPUPercent float64 `json:"cpu_percent"`
}

func GetProcessInfo() ([]ProcessInfo, error) {

	processesInfo := make([]ProcessInfo, 0)

	processes, err := process.Processes()
	if err != nil {
		return nil, fmt.Errorf("get processes: %w", err)
	}

	for _, p := range processes {

		// Process name
		name, err := p.Name()
		if err != nil {
			log.Printf(
				"Name error for PID %d: %v",
				p.Pid,
				err,
			)
			continue
		}

		// Process memory
		memoryInfo, err := p.MemoryInfo()
		if err != nil {
			// log.Printf(
			// 	"Memory error for PID %d: %v",
			// 	p.Pid,
			// 	err,
			// )
			continue
		}

		// Process CPU
		processCPU, err := p.CPUPercent()
		if err != nil {
			processCPU = 0
		}

		processesInfo = append(processesInfo, ProcessInfo{
			PID:        p.Pid,
			Name:       name,
			MemoryRSS:  memoryInfo.RSS,
			CPUPercent: processCPU,
		})
	}

	return processesInfo, nil
}
