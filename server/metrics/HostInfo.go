package metrics

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/host"
)

type HostInfo struct {
	Hostname        string `json:"hostname"`
	OS              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platform_version"`
	KernelVersion   string `json:"kernel_version"`
	Uptime          uint64 `json:"uptime_seconds"`
}

func GetHostInfo() (HostInfo, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return HostInfo{}, fmt.Errorf("get host info: %w", err)
	}

	return HostInfo{
		Hostname:        hostInfo.Hostname,
		OS:              hostInfo.OS,
		Platform:        hostInfo.Platform,
		PlatformVersion: hostInfo.PlatformVersion,
		KernelVersion:   hostInfo.KernelVersion,
		Uptime:          hostInfo.Uptime,
	}, nil
}
