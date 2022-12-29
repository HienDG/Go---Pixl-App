package ui

import (
	"pixl/dvh/appType"
	"pixl/dvh/pxCanvas"
	"pixl/dvh/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlCanvas *pxCanvas.PxCanvas
	PixlWindow fyne.Window
	State      *appType.State
	Swatch     []*swatch.Swatch
}
