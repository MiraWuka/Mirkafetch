package collector

import (
	"context"
	"fmt"
	"strings"

	"github.com/MiraWuka/Mirkafetch/pkg/utils"
)

type PackageManager struct {
	Command string
	Args    []string
	Name    string
}

func (c *SystemCollector) getPackages(ctx context.Context) string {
	packageManagers := []PackageManager{
		{"dpkg", []string{"-l"}, "dpkg"},
		{"rpm", []string{"-qa"}, "rpm"},
		{"brew", []string{"list"}, "brew"},
		{"pacman", []string{"-Q"}, "pacman"},
		{"apk", []string{"list", "--installed"}, "apk"},
		{"pkg", []string{"info"}, "pkg"},
	}

	for _, pm := range packageManagers {
		if count := c.countPackages(ctx, pm); count >= 0 {
			return fmt.Sprintf("%d (%s)", count, pm.Name)
		}
	}

	return "Unknown"
}

func (c *SystemCollector) countPackages(ctx context.Context, pm PackageManager) int {
	out, err := utils.ExecCommand(ctx, pm.Command, pm.Args...)
	if err != nil {
		return -1
	}

	lines := strings.Split(string(out), "\n")
	count := 0

	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			count++
		}
	}

	switch pm.Name {
	case "dpkg":
		return utils.Max(0, count-6)
	case "apk":
		return utils.Max(0, count-1)
	default:
		return utils.Max(0, count-1)
	}
}
