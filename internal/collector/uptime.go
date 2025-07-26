package collector

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/MiraWuka/Mirkafetch/pkg/utils"
)

func (c *SystemCollector) getUptime(ctx context.Context) string {
	switch runtime.GOOS {
	case "linux":
		return c.getLinuxUptime()
	case "darwin":
		return c.getDarwinUptime(ctx)
	case "windows":
		return c.getWindowsUptime(ctx)
	default:
		return "Unknown"
	}
}

func (c *SystemCollector) getLinuxUptime() string {
	content, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "Unknown"
	}

	parts := strings.Fields(string(content))
	if len(parts) == 0 {
		return "Unknown"
	}

	uptimeSeconds, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return "Unknown"
	}

	return c.formatUptime(int64(uptimeSeconds))
}

func (c *SystemCollector) getDarwinUptime(ctx context.Context) string {
	out, err := utils.ExecCommand(ctx, "uptime")
	if err != nil {
		return "Unknown"
	}

	output := strings.TrimSpace(string(out))
	if strings.Contains(output, "up") {
		return c.parseUptimeOutput(output)
	}

	return "Unknown"
}

func (c *SystemCollector) getWindowsUptime(ctx context.Context) string {
	return "Unknown"
}

func (c *SystemCollector) formatUptime(seconds int64) string {
	duration := time.Duration(seconds) * time.Second

	days := int(duration.Hours()) / 24
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	var parts []string
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%dm", minutes))
	}

	if len(parts) == 0 {
		return "< 1m"
	}

	return strings.Join(parts, " ")
}

func (c *SystemCollector) parseUptimeOutput(output string) string {
	return strings.TrimSpace(output)
}
