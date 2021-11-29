package collectd

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const defaultContainerPid int32 = 1

// devices, memory,pids or name=systemd
var cgroupKeyParse = [...]string{"devices", "memory", "pids"}

// RunningInDockerContainer determine if it is running in the Container environment.
// Simply determine by the ways: 1: .dockerenv; 2: cgroups(normal pid=1 in docker)
func RunningInDockerContainer() bool {
	return runningInDockerContainer(defaultContainerPid)
}

// RunningInDockerContainerPid determine if the process with pid is running in the Container environment.
// Simply determine by the ways: 1: .dockerenv; 2: cgroups(normal pid=1 in docker)
func RunningInDockerContainerPid(pid int32) bool {
	return runningInDockerContainer(pid)
}

func runningInDockerContainer(pid int32) bool {
	// docker creates a .dockerenv file at the root of the directory tree inside the Container.
	// if this file exists then the viewer is running from inside a Container so return true
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	// by cgroups
	cgs := parseCgroups(pid)

	if cgs == nil {
		return false
	}

	for _, cgk := range cgroupKeyParse {
		if cg, ok := cgs[cgk]; ok {
			if containsContainerKeyword(cg) {
				return true
			}
		}
	}
	return false
}

func containsContainerKeyword(cg string) bool {
	cg = strings.TrimPrefix(cg, "/")
	if strings.Contains(cg, "docker") || strings.Contains(cg, "kubepods") || strings.Contains(cg, "machind-rkt") || strings.Contains(cg, "sandbox") {
		return true
	}
	return false
}

func parseCgroups(pid int32) map[string]string {
	cgs, err := ParseCgroupFile(fmt.Sprintf("/proc/%d/cgroup", pid))
	if err != nil {
		return nil
	}
	return cgs
}

// ParseCgroupFile parse cgroup file
func ParseCgroupFile(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	return parseCgroupsFromReader(f)
}

func parseCgroupsFromReader(r io.Reader) (map[string]string, error) {
	var (
		cgroups = make(map[string]string)
		s       = bufio.NewScanner(r)
	)
	for s.Scan() {
		var (
			text  = s.Text()
			parts = strings.SplitN(text, ":", 3)
		)
		if len(parts) < 3 {
			return nil, fmt.Errorf("invalid cgroup entry: %q", text)
		}
		for _, subs := range strings.Split(parts[1], ",") {
			if subs != "" {
				cgroups[subs] = parts[2]
			}
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return cgroups, nil
}

func readUint(path string) (uint64, error) {
	v, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}
	return parseUint(strings.TrimSpace(string(v)), 10, 64)
}

func parseUint(s string, base, bitSize int) (uint64, error) {
	v, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		intValue, intErr := strconv.ParseInt(s, base, bitSize)
		// 1. Handle negative values greater than MinInt64 (and)
		// 2. Handle negative values lesser than MinInt64
		if intErr == nil && intValue < 0 {
			return 0, nil
		} else if intErr != nil &&
			intErr.(*strconv.NumError).Err == strconv.ErrRange &&
			intValue < 0 {
			return 0, nil
		}
		return 0, err
	}
	return v, nil
}
