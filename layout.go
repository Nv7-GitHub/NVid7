package main

import (
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

	inspector = ui.NewGroup("Inspector")
	inspector.SetMargined(true)

	hbox.Append(inspector, false)
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
