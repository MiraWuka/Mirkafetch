package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/MiraWuka/Mirkafetch/internal/app"
	"github.com/MiraWuka/Mirkafetch/internal/collector"
	"github.com/MiraWuka/Mirkafetch/internal/display"
)

const (
	defaultTimeout = 30 * time.Second
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	systemCollector := collector.NewSystemCollector()
	consoleDisplay := display.NewConsoleDisplay(os.Stdout)

	application := app.New(systemCollector, consoleDisplay)

	if err := application.Run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
