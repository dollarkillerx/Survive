package utils

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func GetCpuInfo() (string, error) {
	cpuInfos, err := cpu.Info()
	if err != nil {
		return "", err
	}

	return cpuInfos[0].ModelName, nil
}

func CpuUse() (string, error) {
	percent, err := cpu.Percent(time.Millisecond*500, false)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.0f", percent), nil
}

func MemUse() (string, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.0f", memory.UsedPercent), nil
}
