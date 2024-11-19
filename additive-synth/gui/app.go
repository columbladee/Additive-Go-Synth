package gui

import "fyne.io/fyne/v2/app"

func StartApp() {
	application := app.New()
	window := application.NewWindow("Chapstick's Additive Synth 0.01")
	window.ShowAndRun()
}


