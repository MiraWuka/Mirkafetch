package collector

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/MiraWuka/Mirkafetch/pkg/utils"
)

func (c *SystemCollector) getOS(ctx context.Context) string {
	switch runtime.GOOS {
	case "linux":
		return c.getLinuxDistro()
	case "darwin":
		return c.getMacOSVersion(ctx)
	case "windows":
		return c.getWindowsVersion(ctx)
	default:
		return runtime.GOOS
	}
}

func (c *SystemCollector) getLinuxDistro() string {
	content, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "Linux"
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			return strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), "\"")
		}
	}
	return "Linux"
}

func (c *SystemCollector) getMacOSVersion(ctx context.Context) string {
	productOut, err := utils.ExecCommand(ctx, "sw_vers", "-productName")
	if err != nil {
		return "macOS"
	}

	productName := strings.TrimSpace(string(productOut))

	versionOut, err := utils.ExecCommand(ctx, "sw_vers", "-productVersion")
	if err != nil {
		return productName
	}

	version := strings.TrimSpace(string(versionOut))
	return fmt.Sprintf("%s %s", productName, version)
}

func (c *SystemCollector) getWindowsVersion(ctx context.Context) string {
	out, err := utils.ExecCommand(ctx, "cmd", "/c", "ver")
	if err != nil {
		return "Windows"
	}
	return strings.TrimSpace(string(out))
}
