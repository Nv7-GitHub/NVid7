package main

import (
	"github.com/Nv7-Github/NVid7/importers"
	"github.com/andlabs/ui"
)

var imps = []func() importers.Importer{importers.NewPNGImporter}
var impNames = []string{"PNG"}

func createMenu() *ui.Grid {
	menu := ui.NewGrid()

	impsBox := ui.NewCombobox()
	impsBox.Append("Import")
	for _, name := range impNames {
		impsBox.Append(name)
	}
	impsBox.SetSelected(0)
	impsBox.OnSelected(func(i *ui.Combobox) {
		if i.Selected() == 0 {
			return
		}
		index := seq.ClipCount()
		seq.AddClip(imps[i.Selected()-1]())
		seq.SetSelected(index)
		i.SetSelected(0)
		seq.Update()
	})

	menu.Append(impsBox, 0, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
	return menu
}
