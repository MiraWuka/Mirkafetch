package display

import "github.com/MiraWuka/Mirkafetch/internal/models"

type Display interface {
	Show(info *models.SystemInfo) error
}
