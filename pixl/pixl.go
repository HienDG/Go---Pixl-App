package main

import (
	"image/color"
	"pixl/dvh/appType"
	"pixl/dvh/pxCanvas"
	"pixl/dvh/swatch"
	"pixl/dvh/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("Pixl App")

	pixlCanvasConfig := appType.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       10,
		PxCols:       10,
		PxSize:       30,
	}

	state := appType.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	pixlCanvas := pxCanvas.NewPxCanvas(&state, pixlCanvasConfig)

	appInit := ui.AppInit{
		PixlCanvas: pixlCanvas,
		PixlWindow: pixlWindow,
		State:      &state,
		Swatch:     make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)
	appInit.PixlWindow.ShowAndRun()

}
