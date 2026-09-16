package metrics

import (
	"encoding/json"
	"fmt"
	"log"
)

type SystemInfo struct {
	Host      HostInfo      `json:"host"`
	CPU       CPUInfo       `json:"cpu"`
	Memory    MemoryInfo    `json:"memory"`
	Disks     []DiskInfo    `json:"disks"`
	Network   []NetInfo     `json:"network"`
	Processes []ProcessInfo `json:"processes"`
}

func ShowMetrics() {

	host, err := GetHostInfo()
	if err != nil {
		log.Fatal(err)
	}

	cpu, err := GetCPUInfo()
	if err != nil {
		log.Fatal(err)
	}

	memory, err := GetMemoryInfo()
	if err != nil {
		log.Fatal(err)
	}

	disks, err := GetDiskInfo()
	if err != nil {
		log.Fatal(err)
	}

	network, err := GetNetworkInfo()
	if err != nil {
		log.Fatal(err)
	}

	processes, err := GetProcessInfo()
	if err != nil {
		log.Fatal(err)
	}

	systemInfo := SystemInfo{
		Host:      host,
		CPU:       cpu,
		Memory:    memory,
		Disks:     disks,
		Network:   network,
		Processes: processes,
	}

	data, err := json.MarshalIndent(systemInfo, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
