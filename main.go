package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func handle(err error) {
	if err != nil {
		ui.QueueMain(func() {
			ui.MsgBoxError(win, "Error!", err.Error())
		})
	}
}

func main() {
	ui.Main(setupUI)
}
