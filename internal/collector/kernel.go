package collector

import (
	"context"
	"os"
	"runtime"
	"strings"

	"github.com/MiraWuka/Mirkafetch/pkg/utils"
)

func (c *SystemCollector) getKernel(ctx context.Context) string {
	switch runtime.GOOS {
	case "linux":
		return c.getLinuxKernel()
	case "darwin":
		return c.getDarwinKernel(ctx)
	case "windows":
		return c.getWindowsKernel(ctx)
	default:
		return "Unknown"
	}
}

func (c *SystemCollector) getLinuxKernel() string {
	content, err := os.ReadFile("/proc/version")
	if err != nil {
		return "Unknown"
	}

	versionStr := string(content)
	if strings.Contains(versionStr, "Linux version") {
		parts := strings.Fields(versionStr)
		if len(parts) >= 3 {
			return parts[2]
		}
	}
	return "Unknown"
}

func (c *SystemCollector) getDarwinKernel(ctx context.Context) string {
	out, err := utils.ExecCommand(ctx, "uname", "-r")
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSpace(string(out))
}

func (c *SystemCollector) getWindowsKernel(ctx context.Context) string {
	out, err := utils.ExecCommand(ctx, "cmd", "/c", "ver")
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSpace(string(out))
}
