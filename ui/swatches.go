package ui

import (
	"image/color"
	"pixl/dvh/swatch"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func BuildSwatches(app *AppInit) *fyne.Container {
	canvasSwatch := make([]fyne.CanvasObject, 0, 64)

	for i := 0; i < cap(app.Swatch); i++ {
		initialColor := color.NRGBA{255, 255, 255, 255}
		s := swatch.NewSwatch(app.State, initialColor, i, func(s *swatch.Swatch) {
			for j := 0; i < len(app.Swatch); j++ {
				app.Swatch[j].Selected = false
				canvasSwatch[j].Refresh()
			}
			app.State.SwatchSelected = s.SwatchIndex
			app.State.BrushColor = s.Color
		})

		if i == 0 {
			s.Selected = true
			app.State.SwatchSelected = 0
			s.Refresh()
		}
		app.Swatch = append(app.Swatch, s)
		canvasSwatch = append(canvasSwatch, s)
	}

	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatch...)
}
