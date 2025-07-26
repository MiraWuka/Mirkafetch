// Package collector handles system information gathering.
package collector

import (
	"context"
	"fmt"

	"github.com/MiraWuka/Mirkafetch/internal/models"
)

type SystemCollector struct{}

func NewSystemCollector() *SystemCollector {
	return &SystemCollector{}
}

func (c *SystemCollector) Collect(ctx context.Context) (*models.SystemInfo, error) {
	info := &models.SystemInfo{}

	var err error
	if info.User, err = c.getUser(); err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}

	if info.Hostname, err = c.getHostname(); err != nil {
		return nil, fmt.Errorf("failed to get hostname: %w", err)
	}

	info.OS = c.getOS(ctx)
	info.Kernel = c.getKernel(ctx)
	info.Uptime = c.getUptime(ctx)
	info.Shell = c.getShell()
	info.CPU = c.getCPU(ctx)
	info.Memory = c.getMemory(ctx)
	info.Disk = c.getDisk(ctx)
	info.Packages = c.getPackages(ctx)
	info.GPU = c.getGPU(ctx)

	return info, nil
}
