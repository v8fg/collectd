package collectd

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
)

// CpuInfoData cpu info data
type CpuInfoData struct {
	count     float64
	used      float64
	percent   float64
	container int
}

// CpuInfo read the cpu info for container or not, with the special pid
func CpuInfo(container bool, pid int32) CpuInfoData {
	cid := CpuInfoData{}

	if container {
		cid.container = 1
		cpuPeriod, _ := readUint("/sys/fs/cgroup/cpu/cpu.cfs_period_us")
		cpuQuota, _ := readUint("/sys/fs/cgroup/cpu/cpu.cfs_quota_us")
		if cpuPeriod > 0 && cpuQuota > 0 {
			cid.count = float64(cpuQuota) / float64(cpuPeriod)
		}
		if cid.count > 0 {
			if p, err := process.NewProcess(pid); err == nil {
				if cpuPercent, err := p.Percent(time.Second); err == nil {
					cid.percent = cpuPercent / cid.count
					cid.used = cpuPercent / 100 // unit count
				}
			}
		}
	} else {
		count, _ := cpu.Counts(true)
		cid.count = float64(count)

		percent, _ := cpu.Percent(time.Second, false)
		cid.percent = percent[0]
		cid.used = cid.count * cid.percent / 100 // unit count
	}
	return cid
}
