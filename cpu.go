package collectd

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
)

// CpuInfoData cpu info data
type CpuInfoData struct {
	Count     float64
	Used      float64
	Percent   float64
	Container int
}

// CpuInfo read the cpu info for Container or not, with the special pid
func CpuInfo(container bool, pid int32) CpuInfoData {
	cid := CpuInfoData{}

	if container {
		cid.Container = 1
		cpuPeriod, _ := readUint("/sys/fs/cgroup/cpu/cpu.cfs_period_us")
		cpuQuota, _ := readUint("/sys/fs/cgroup/cpu/cpu.cfs_quota_us")
		if cpuPeriod > 0 && cpuQuota > 0 {
			cid.Count = float64(cpuQuota) / float64(cpuPeriod)
		}
		if cid.Count > 0 {
			if p, err := process.NewProcess(pid); err == nil {
				if cpuPercent, err := p.Percent(time.Second); err == nil {
					cid.Percent = cpuPercent / cid.Count
					cid.Used = cpuPercent / 100 // unit Count
				}
			}
		}
	} else {
		count, _ := cpu.Counts(true)
		cid.Count = float64(count)

		percent, _ := cpu.Percent(time.Second, false)
		cid.Percent = percent[0]
		cid.Used = cid.Count * cid.Percent / 100 // unit Count
	}
	return cid
}
