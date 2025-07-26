package collector

import (
	"context"
	"runtime"
	"strings"

	"github.com/MiraWuka/Mirkafetch/pkg/utils"
)

func (c *SystemCollector) getGPU(ctx context.Context) string {
	switch runtime.GOOS {
	case "linux":
		return c.getLinuxGPU(ctx)
	case "darwin":
		return c.getDarwinGPU(ctx)
	case "windows":
		return c.getWindowsGPU(ctx)
	default:
		return "Unknown"
	}
}

func (c *SystemCollector) getLinuxGPU(ctx context.Context) string {
	out, err := utils.ExecCommand(ctx, "lspci")
	if err != nil {
		return "Unknown"
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		lowerLine := strings.ToLower(line)
		if strings.Contains(lowerLine, "vga") ||
			strings.Contains(lowerLine, "3d") ||
			strings.Contains(lowerLine, "display") {
			parts := strings.Split(line, ": ")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}

	return "Unknown"
}

func (c *SystemCollector) getDarwinGPU(ctx context.Context) string {
	out, err := utils.ExecCommand(ctx, "system_profiler", "SPDisplaysDataType")
	if err != nil {
		return "Unknown"
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Chipset Model:") ||
			strings.Contains(line, "Graphics:") {
			parts := strings.Split(line, ": ")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}

	return "Unknown"
}

func (c *SystemCollector) getWindowsGPU(ctx context.Context) string {
	return "Unknown"
}
