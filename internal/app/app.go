// Package app contains the main application logic.
package app

import (
	"context"
	"fmt"

	"github.com/MiraWuka/Mirkafetch/internal/models"
)

// Collector defines the interface for collecting system information.
type Collector interface {
	Collect(ctx context.Context) (*models.SystemInfo, error)
}

// Display defines the interface for displaying system information.
type Display interface {
	Show(info *models.SystemInfo) error
}

// App represents the main application.
type App struct {
	collector Collector
	display   Display
}

// New creates a new application instance.
func New(collector Collector, display Display) *App {
	return &App{
		collector: collector,
		display:   display,
	}
}

// Run executes the application.
func (a *App) Run(ctx context.Context) error {
	info, err := a.collector.Collect(ctx)
	if err != nil {
		return fmt.Errorf("failed to collect system information: %w", err)
	}

	if err := a.display.Show(info); err != nil {
		return fmt.Errorf("failed to display system information: %w", err)
	}

	return nil
}
