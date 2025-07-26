package display

import (
	"fmt"
	"io"
	"strings"

	"github.com/MiraWuka/Mirkafetch/internal/models"
	"github.com/MiraWuka/Mirkafetch/pkg/utils"
)

type ConsoleDisplay struct {
	writer io.Writer
}

func NewConsoleDisplay(writer io.Writer) *ConsoleDisplay {
	return &ConsoleDisplay{writer: writer}
}

func (d *ConsoleDisplay) Show(info *models.SystemInfo) error {
	logoWidth := d.calculateLogoWidth()
	padding := 2

	items := d.createInfoItems(info)

	header := fmt.Sprintf("%s%s%s@%s%s%s",
		utils.ColorBold+utils.ColorGreen, info.User, utils.ColorReset,
		utils.ColorBold+utils.ColorGreen, info.Hostname, utils.ColorReset)

	if err := d.writeLine(header); err != nil {
		return err
	}

	separator := strings.Repeat("-", len(info.User)+len(info.Hostname)+1)
	if err := d.writeLine(separator); err != nil {
		return err
	}

	maxLines := utils.Max(len(utils.ASCIILogo), len(items))
	for i := 0; i < maxLines; i++ {
		if err := d.displayLine(i, logoWidth, padding, items); err != nil {
			return err
		}
	}

	if err := d.displayColorPalette(); err != nil {
		return err
	}

	return nil
}

func (d *ConsoleDisplay) calculateLogoWidth() int {
	width := 0
	for _, line := range utils.ASCIILogo {
		if len(line) > width {
			width = len(line)
		}
	}
	return width
}

func (d *ConsoleDisplay) createInfoItems(info *models.SystemInfo) []models.InfoItem {
	return []models.InfoItem{
		{"OS", info.OS, utils.ColorBlue},
		{"Kernel", info.Kernel, utils.ColorYellow},
		{"Uptime", info.Uptime, utils.ColorMagenta},
		{"Shell", info.Shell, utils.ColorGreen},
		{"CPU", info.CPU, utils.ColorRed},
		{"GPU", info.GPU, utils.ColorCyan},
		{"Memory", info.Memory, utils.ColorYellow},
		{"Disk", info.Disk, utils.ColorGreen},
		{"Packages", info.Packages, utils.ColorBlue},
	}
}

func (d *ConsoleDisplay) displayLine(index, logoWidth, padding int, items []models.InfoItem) error {
	logoPart := ""
	if index < len(utils.ASCIILogo) {
		logoPart = utils.ASCIILogo[index]
	}

	line := fmt.Sprintf("%-*s", logoWidth+padding, logoPart)

	if index < len(items) {
		item := items[index]
		line += fmt.Sprintf("%s%-9s%s: %s",
			item.Color, item.Label, utils.ColorReset, item.Value)
	}

	return d.writeLine(line)
}

func (d *ConsoleDisplay) displayColorPalette() error {
	if err := d.writeLine(""); err != nil {
		return err
	}

	colors := []string{
		utils.ColorRed, utils.ColorGreen, utils.ColorYellow,
		utils.ColorBlue, utils.ColorMagenta, utils.ColorCyan, utils.ColorWhite,
	}

	palette := ""
	for _, color := range colors {
		palette += fmt.Sprintf("%s███%s", color, utils.ColorReset)
	}

	return d.writeLine(palette)
}

func (d *ConsoleDisplay) writeLine(line string) error {
	_, err := fmt.Fprintln(d.writer, line)
	return err
}
