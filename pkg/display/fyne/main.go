package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image"
)

type Display struct { //todo hide type
	window fyne.Window
	app    fyne.App
}

func (d *Display) DisplayImage(image image.Image) {
	fyimg := canvas.NewImageFromImage(image)
	d.window.Canvas().SetContent(fyimg)
}

//shows the display and blocks until it closes
func (d *Display) ShowAndRun() {
	d.window.ShowAndRun()
}

func InitDisplay(width, height int) *Display {
	a := app.New()
	w := a.NewWindow("OpenSong-remote-display")
	canvas.NewImageFromFile("")

	w.Resize(fyne.NewSize(float32(width), float32(height)))

	var d Display
	d.app = a
	d.window = w

	//w.SetFullScreen(true)
	return &d
}
