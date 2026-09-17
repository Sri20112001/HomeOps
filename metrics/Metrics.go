package metrics

type SystemInfo struct {
	Host      HostInfo      `json:"host"`
	CPU       CPUInfo       `json:"cpu"`
	Memory    MemoryInfo    `json:"memory"`
	Disks     []DiskInfo    `json:"disks"`
	Network   []NetInfo     `json:"network"`
	Processes []ProcessInfo `json:"processes"`
}

func GetSystemInfo() (SystemInfo, error) {
	host, err := GetHostInfo()
	if err != nil {
		return SystemInfo{}, err
	}

	cpu, err := GetCPUInfo()
	if err != nil {
		return SystemInfo{}, err
	}

	memory, err := GetMemoryInfo()
	if err != nil {
		return SystemInfo{}, err
	}

	disks, err := GetDiskInfo()
	if err != nil {
		return SystemInfo{}, err
	}

	network, err := GetNetworkInfo()
	if err != nil {
		return SystemInfo{}, err
	}

	processes, err := GetProcessInfo()
	if err != nil {
		return SystemInfo{}, err
	}

	return SystemInfo{
		Host:      host,
		CPU:       cpu,
		Memory:    memory,
		Disks:     disks,
		Network:   network,
		Processes: processes,
	}, nil
}
