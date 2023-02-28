package system

import (
	"fmt"
	"os"
	"runtime"
)

type Capabilities struct {
	Hostname   string `json:"hostname" binding:"required"`
	Cpus       int    `json:"cpus" binding:"required"`
	Os         string `json:"os" binding:"required"`
	Hypervisor string `json:"hypervisor" binding:"required"`
	Ram        struct {
		AllocMiB  uint64 `json:"alloc_MiB"`
		SystemMiB uint64 `json:"system_MiB"`
		Gc        uint32 `json:"gc"`
	} `json:"memory"`
}

func Info() Capabilities {

	var capabilities Capabilities
	// Memory
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// Hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	capabilities.Hostname = hostname
	capabilities.Cpus = runtime.NumCPU()

	capabilities.Ram.AllocMiB = bToMb(m.TotalAlloc)
	capabilities.Ram.SystemMiB = bToMb(bToMb(m.Sys))
	capabilities.Ram.Gc = m.NumGC

	return capabilities

}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
