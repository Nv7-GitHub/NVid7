package importers

import "github.com/andlabs/ui"

func makeBlendModeCb(blendModeVal *BlendMode) *ui.Combobox {
	cb := ui.NewCombobox()
	blendModes := make([]BlendMode, len(BlendModeNames))
	i := 0
	for k := range BlendModeNames {
		cb.Append(k)
		blendModes[i] = BlendModeNames[k]
	}
	cb.SetSelected(0)
	cb.OnSelected(func(c *ui.Combobox) {
		*blendModeVal = blendModes[c.Selected()]
	})
	return cb
}
