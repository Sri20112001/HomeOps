package metrics

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/net"
)

type NetInfo struct {
	Name        string `json:"name"`
	BytesSent   uint64 `json:"bytes_sent"`
	BytesRecv   uint64 `json:"bytes_received"`
	PacketsSent uint64 `json:"packets_sent"`
	PacketsRecv uint64 `json:"packets_received"`
}

func GetNetworkInfo() ([]NetInfo, error) {
	networks := make([]NetInfo, 0)

	networkInfo, err := net.IOCounters(true)
	if err != nil {
		return nil, fmt.Errorf("network error: %w", err)
	}

	for _, network := range networkInfo {

		networks = append(networks, NetInfo{
			Name:        network.Name,
			BytesSent:   network.BytesSent,
			BytesRecv:   network.BytesRecv,
			PacketsSent: network.PacketsSent,
			PacketsRecv: network.PacketsRecv,
		})
	}
	return networks, nil
}
