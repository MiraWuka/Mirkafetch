package app

import (
	"context"
	"fmt"

	"github.com/MiraWuka/Mirkafetch/internal/models"
)

type Collector interface {
	Collect(ctx context.Context) (*models.SystemInfo, error)
}

type Display interface {
	Show(info *models.SystemInfo) error
}

type App struct {
	collector Collector
	display   Display
}

func New(collector Collector, display Display) *App {
	return &App{
		collector: collector,
		display:   display,
	}
}

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
