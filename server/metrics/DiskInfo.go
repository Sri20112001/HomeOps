package metrics

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/disk"
)

type DiskInfo struct {
	Path        string  `json:"path"`
	Filesystem  string  `json:"filesystem"`
	Total       uint64  `json:"total_bytes"`
	Used        uint64  `json:"used_bytes"`
	Free        uint64  `json:"free_bytes"`
	UsedPercent float64 `json:"used_percent"`
}

func GetDiskInfo() ([]DiskInfo, error) {

	disks := make([]DiskInfo, 0)

	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, fmt.Errorf("get disk partitions: %w", err)
	}

	for _, partition := range partitions {

		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			// One disk failing shouldn't prevent us
			// from collecting the other disks.
			continue
		}

		disks = append(disks, DiskInfo{
			Path:        usage.Path,
			Filesystem:  usage.Fstype,
			Total:       usage.Total,
			Used:        usage.Used,
			Free:        usage.Free,
			UsedPercent: usage.UsedPercent,
		})
	}

	return disks, nil
}
