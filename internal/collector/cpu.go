package collector

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/MiraWuka/Mirkafetch/pkg/utils"
)

func (c *SystemCollector) getCPU(ctx context.Context) string {
	switch runtime.GOOS {
	case "linux":
		return c.getLinuxCPU()
	case "darwin":
		return c.getDarwinCPU(ctx)
	case "windows":
		return c.getWindowsCPU(ctx)
	default:
		return fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
	}
}

func (c *SystemCollector) getLinuxCPU() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cores := 0
	var modelName string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "model name") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				modelName = strings.TrimSpace(parts[1])
			}
		}
		if strings.HasPrefix(line, "processor") {
			cores++
		}
	}

	if modelName != "" {
		return fmt.Sprintf("%s (%d cores)", modelName, cores)
	}

	return fmt.Sprintf("%s %s (%d cores)", runtime.GOOS, runtime.GOARCH, cores)
}

func (c *SystemCollector) getDarwinCPU(ctx context.Context) string {
	brandOut, err := utils.ExecCommand(ctx, "sysctl", "-n", "machdep.cpu.brand_string")
	if err != nil {
		return fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
	}

	coreOut, err := utils.ExecCommand(ctx, "sysctl", "-n", "hw.ncpu")
	if err != nil {
		return strings.TrimSpace(string(brandOut))
	}

	brand := strings.TrimSpace(string(brandOut))
	cores := strings.TrimSpace(string(coreOut))

	return fmt.Sprintf("%s (%s cores)", brand, cores)
}

func (c *SystemCollector) getWindowsCPU(ctx context.Context) string {
	return fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
}
