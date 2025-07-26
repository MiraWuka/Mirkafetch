package collector

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/MiraWuka/Mirkafetch/pkg/utils"
)

func (c *SystemCollector) getDisk(ctx context.Context) string {
	switch runtime.GOOS {
	case "linux", "darwin":
		return c.getUnixDisk(ctx)
	case "windows":
		return c.getWindowsDisk(ctx)
	default:
		return "Unknown"
	}
}

func (c *SystemCollector) getUnixDisk(ctx context.Context) string {
	out, err := utils.ExecCommand(ctx, "df", "-h", "/")
	if err != nil {
		return "Unknown"
	}

	lines := strings.Split(string(out), "\n")
	if len(lines) >= 2 {
		fields := strings.Fields(lines[1])
		if len(fields) >= 5 {
			return fmt.Sprintf("%s / %s (%s)", fields[2], fields[1], fields[4])
		}
	}

	return "Unknown"
}

func (c *SystemCollector) getWindowsDisk(ctx context.Context) string {
	if diskInfo := c.getWindowsDiskPowerShell(ctx); diskInfo != "Unknown" {
		return diskInfo
	}

	return c.getWindowsDiskWmic(ctx)
}

func (c *SystemCollector) getWindowsDiskPowerShell(ctx context.Context) string {
	// PowerShell command to get C: drive information
	psCmd := `Get-WmiObject -Class Win32_LogicalDisk -Filter "DeviceID='C:'" | Select-Object @{Name="Used";Expression={[math]::Round(($_.Size-$_.FreeSpace)/1GB,1)}}, @{Name="Total";Expression={[math]::Round($_.Size/1GB,1)}}, @{Name="Percent";Expression={[math]::Round((($_.Size-$_.FreeSpace)/$_.Size)*100,0)}} | Format-List`

	out, err := utils.ExecCommand(ctx, "powershell", "-Command", psCmd)
	if err != nil {
		return "Unknown"
	}

	return c.parseWindowsPowerShellOutput(string(out))
}

func (c *SystemCollector) getWindowsDiskWmic(ctx context.Context) string {
	// Get total size
	sizeOut, err := utils.ExecCommand(ctx, "wmic", "logicaldisk", "where", "caption=\"C:\"", "get", "size", "/value")
	if err != nil {
		return "Unknown"
	}

	freeOut, err := utils.ExecCommand(ctx, "wmic", "logicaldisk", "where", "caption=\"C:\"", "get", "freespace", "/value")
	if err != nil {
		return "Unknown"
	}

	return c.parseWindowsWmicOutput(string(sizeOut), string(freeOut))
}

func (c *SystemCollector) parseWindowsPowerShellOutput(output string) string {
	lines := strings.Split(output, "\n")
	var used, total, percent string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Used") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				used = strings.TrimSpace(parts[1]) + "GB"
			}
		} else if strings.HasPrefix(line, "Total") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				total = strings.TrimSpace(parts[1]) + "GB"
			}
		} else if strings.HasPrefix(line, "Percent") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				percent = strings.TrimSpace(parts[1]) + "%"
			}
		}
	}

	if used != "" && total != "" && percent != "" {
		return fmt.Sprintf("%s / %s (%s)", used, total, percent)
	}

	return "Unknown"
}

func (c *SystemCollector) parseWindowsWmicOutput(sizeOutput, freeOutput string) string {
	var totalBytes, freeBytes int64

	sizeLines := strings.Split(sizeOutput, "\n")
	for _, line := range sizeLines {
		if strings.Contains(line, "Size=") {
			sizeStr := strings.TrimPrefix(strings.TrimSpace(line), "Size=")
			if size := c.parseWindowsBytes(sizeStr); size > 0 {
				totalBytes = size
			}
		}
	}

	freeLines := strings.Split(freeOutput, "\n")
	for _, line := range freeLines {
		if strings.Contains(line, "FreeSpace=") {
			freeStr := strings.TrimPrefix(strings.TrimSpace(line), "FreeSpace=")
			if free := c.parseWindowsBytes(freeStr); free > 0 {
				freeBytes = free
			}
		}
	}

	if totalBytes > 0 && freeBytes > 0 {
		usedBytes := totalBytes - freeBytes
		usedPercent := float64(usedBytes) / float64(totalBytes) * 100

		return fmt.Sprintf("%s / %s (%.0f%%)",
			utils.FormatBytes(usedBytes),
			utils.FormatBytes(totalBytes),
			usedPercent)
	}

	return "Unknown"
}

func (c *SystemCollector) parseWindowsBytes(byteStr string) int64 {
	byteStr = strings.TrimSpace(byteStr)
	if byteStr == "" {
		return 0
	}

	var bytes int64
	if _, err := fmt.Sscanf(byteStr, "%d", &bytes); err == nil {
		return bytes
	}

	return 0
}
