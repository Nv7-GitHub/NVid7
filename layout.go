package main

import (
	"image"
	"math"

	"github.com/andlabs/ui"
)

var win *ui.Window
var inspector *ui.Group
var seq *Sequencer

func setupUI() {
	win = ui.NewWindow("NVideo", 800, 600, true)
	win.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		win.Destroy()
		return true
	})
	win.SetMargined(true)

	outerbox := ui.NewVerticalBox()
	outerbox.Append(createMenu(), false)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	imGroup := ui.NewGroup("Image")
	imGroup.SetMargined(true)
	hbox.Append(imGroup, true)

	ibox := ui.NewVerticalBox()

	settings := ui.NewGroup("Settings")
	settings.SetMargined(true)
	sForm := ui.NewForm()
	sForm.SetPadded(true)

	outWidth := ui.NewSpinbox(1, math.MaxInt32)
	outHeight := ui.NewSpinbox(1, math.MaxInt32)
	outWidth.SetValue(1920)
	outHeight.SetValue(1080)

	outWidth.OnChanged(func(s *ui.Spinbox) {
		sz := seq.Bounds()
		seq.SetOutputDimensions(image.Rect(0, 0, s.Value(), sz.Dy()))
		seq.Update()
	})
	outHeight.OnChanged(func(s *ui.Spinbox) {
		sz := seq.Bounds()
		seq.SetOutputDimensions(image.Rect(0, 0, sz.Dx(), s.Value()))
		seq.Update()
	})

	sForm.Append("Output Width", outWidth, false)
	sForm.Append("Output Height", outHeight, false)
	settings.SetChild(sForm)

	ibox.Append(settings, false)

	inspector = ui.NewGroup("Inspector")
	inspector.SetMargined(true)
	ibox.Append(inspector, false)

	hbox.Append(ibox, false)
	vbox.Append(hbox, true)

	sequ := ui.NewGroup("Sequencer")
	sequ.SetMargined(true)

	seq = NewSequencer()
	// TODO: Make slider or spinbox to change length
	seq.SetAnimationLength(250)
	area := ui.NewArea(seq)
	seq.SetArea(area)

	sequ.SetChild(area)
	vbox.Append(sequ, true)
	outerbox.Append(vbox, true)

	win.SetChild(outerbox)

	win.Show()
}
