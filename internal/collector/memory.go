package collector

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/MiraWuka/Mirkafetch/pkg/utils"
)

func (c *SystemCollector) getMemory(ctx context.Context) string {
	switch runtime.GOOS {
	case "linux":
		return c.getLinuxMemory()
	case "darwin":
		return c.getDarwinMemory(ctx)
	case "windows":
		return c.getWindowsMemory(ctx)
	default:
		return "Unknown"
	}
}

func (c *SystemCollector) getLinuxMemory() string {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return "Unknown"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalKB, availableKB int64

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		switch fields[0] {
		case "MemTotal:":
			if val, err := strconv.ParseInt(fields[1], 10, 64); err == nil {
				totalKB = val
			}
		case "MemAvailable:":
			if val, err := strconv.ParseInt(fields[1], 10, 64); err == nil {
				availableKB = val
			}
		}
	}

	if totalKB > 0 && availableKB > 0 {
		usedKB := totalKB - availableKB
		usedBytes := usedKB * 1024
		totalBytes := totalKB * 1024
		return fmt.Sprintf("%s / %s", utils.FormatBytes(usedBytes), utils.FormatBytes(totalBytes))
	}

	return "Unknown"
}

func (c *SystemCollector) getDarwinMemory(ctx context.Context) string {
	out, err := utils.ExecCommand(ctx, "sysctl", "-n", "hw.memsize")
	if err != nil {
		return "Unknown"
	}

	total, err := strconv.ParseInt(strings.TrimSpace(string(out)), 10, 64)
	if err != nil {
		return "Unknown"
	}

	return fmt.Sprintf("? / %s", utils.FormatBytes(total))
}

func (c *SystemCollector) getWindowsMemory(ctx context.Context) string {
	return "Unknown"
}
