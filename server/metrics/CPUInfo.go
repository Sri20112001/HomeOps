package metrics

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUInfo struct {
	ModelName     string  `json:"model_name"`
	PhysicalCores int     `json:"physical_cores"`
	LogicalCores  int     `json:"logical_cores"`
	Mhz           float64 `json:"mhz"`
	Usage         float64 `json:"usage_percent"`
}

func GetCPUInfo() (CPUInfo, error) {

	cpuInfo, err := cpu.Info()
	if err != nil {
		return CPUInfo{}, fmt.Errorf("get CPU info: %w", err)
	}

	// Physical cores
	physicalCores, err := cpu.Counts(false)
	if err != nil {
		return CPUInfo{}, fmt.Errorf("get physical cores: %w", err)
	}

	// Logical cores
	logicalCores, err := cpu.Counts(true)
	if err != nil {
		return CPUInfo{}, fmt.Errorf("get logical cores: %w", err)
	}

	// CPU usage
	cpuUsage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return CPUInfo{}, fmt.Errorf("get CPU usage: %w", err)
	}

	var cpuModel string
	var cpuMhz float64
	var cpuPercent float64

	if len(cpuInfo) > 0 {
		cpuModel = cpuInfo[0].ModelName
		cpuMhz = cpuInfo[0].Mhz
	}

	if len(cpuUsage) > 0 {
		cpuPercent = cpuUsage[0]
	}

	return CPUInfo{
		ModelName:     cpuModel,
		PhysicalCores: physicalCores,
		LogicalCores:  logicalCores,
		Mhz:           cpuMhz,
		Usage:         cpuPercent,
	}, nil
}
