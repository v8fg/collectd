package collectd

import (
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

// MemInfoData memory info data
type MemInfoData struct {
	total     float64
	used      float64
	percent   float64
	container int
}

// MemInfo read the memory info for container or not, with the special pid
func MemInfo(container bool, pid int32) MemInfoData {
	mid := MemInfoData{}

	if container {
		mid.container = 1
		memLimit, _ := readUint("/sys/fs/cgroup/memory/memory.limit_in_bytes")
		mid.total = float64(memLimit)

		p, err := process.NewProcess(pid)
		if err == nil {
			mif, err := p.MemoryInfo()
			if err == nil {
				mid.percent = float64(mif.RSS) * 100 / mid.total
				mid.used = float64(mif.RSS)
			}
		}
	} else {
		p, err := process.NewProcess(pid)
		if err == nil {
			vm, err := mem.VirtualMemory()
			if err == nil {
				mid.total = float64(vm.Total)
			}

			mp, _ := p.MemoryPercent()
			mid.percent = float64(mp)
			mid.used = mid.total * float64(mp) / 100
		}
	}
	return mid
}
