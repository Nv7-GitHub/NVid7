package main

import (
	"github.com/andlabs/ui"
)

var win *ui.Window

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

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	imGroup := ui.NewGroup("Image")
	imGroup.SetMargined(true)
	hbox.Append(imGroup, true)

	inspectorGroup := ui.NewGroup("Inspector")
	inspectorGroup.SetMargined(true)

	inspectorForm := ui.NewForm()
	inspectorForm.SetPadded(true)
	inspectorForm.Append("Input", ui.NewEntry(), false)
	inspectorForm.Append("Button", ui.NewButton("Button"), false)

	inspectorGroup.SetChild(inspectorForm)
	hbox.Append(inspectorGroup, false)
	vbox.Append(hbox, true)

	area := ui.NewGroup("Sequencer")
	area.SetMargined(true)
	vbox.Append(area, true)

	win.SetChild(vbox)

	win.Show()
}
