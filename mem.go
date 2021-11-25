package collectd

import (
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

// MemInfoData memory info data
type MemInfoData struct {
	Total     float64
	Used      float64
	Percent   float64
	Container int
}

// MemInfo read the memory info for Container or not, with the special pid
func MemInfo(container bool, pid int32) MemInfoData {
	mid := MemInfoData{}

	if container {
		mid.Container = 1
		memLimit, _ := readUint("/sys/fs/cgroup/memory/memory.limit_in_bytes")
		mid.Total = float64(memLimit)

		p, err := process.NewProcess(pid)
		if err == nil {
			mif, err := p.MemoryInfo()
			if err == nil {
				mid.Percent = float64(mif.RSS) * 100 / mid.Total
				mid.Used = float64(mif.RSS)
			}
		}
	} else {
		p, err := process.NewProcess(pid)
		if err == nil {
			vm, err := mem.VirtualMemory()
			if err == nil {
				mid.Total = float64(vm.Total)
			}

			mp, _ := p.MemoryPercent()
			mid.Percent = float64(mp)
			mid.Used = mid.Total * float64(mp) / 100
		}
	}
	return mid
}
