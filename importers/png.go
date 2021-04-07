package importers

import (
	"errors"
	"image"
	"image/png"
	"os"

	"github.com/andlabs/ui"
)

type PNGImporter struct {
	Path     string
	imgCache image.Image
}

func NewPNGImporter() Importer {
	return &PNGImporter{}
}

func (i *PNGImporter) refreshCache() error {
	file, err := os.Open(i.Path)
	if err != nil {
		return err
	}
	i.imgCache, err = png.Decode(file)
	return err
}

func (i *PNGImporter) MakeUI(win *ui.Window, recalcFunc func()) (ui.Control, error) {
	form := ui.NewForm()
	filename := ""
	hbox := ui.NewHorizontalBox()

	fileShower := ui.NewEntry()
	saveBtn := ui.NewButton("Select")
	saveBtn.OnClicked(func(*ui.Button) {
		err := i.refreshCache()
		if err != nil {
			ui.MsgBoxError(win, "Error!", err.Error())
			return
		}

		filename = ui.SaveFile(win)
		fileShower.SetText(filename)
		recalcFunc()
	})

	hbox.Append(fileShower, true)
	hbox.Append(saveBtn, false)
	form.Append("Output File", hbox, false)
	return form, nil
}

func (i *PNGImporter) GetFrame(time int) (image.Image, error) {
	if i.imgCache == nil {
		return nil, errors.New("image hasn't been selected")
	}
	return i.imgCache, nil
}

func (i *PNGImporter) Cleanup() error       { return nil }
func (i *PNGImporter) Length() (int, error) { return -1, nil }
func (i *PNGImporter) Name() string         { return "PNG Clip: " + i.Path }
