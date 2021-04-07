package importers

import (
	"errors"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/andlabs/ui"
)

type PNGImporter struct {
	Path      string
	imgCache  image.Image
	BlendMode BlendMode
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

func (i *PNGImporter) MakeUI(win *ui.Window, recalcFunc func(int), index int) (ui.Control, error) {
	form := ui.NewForm()
	form.SetPadded(true)
	filename := ""
	hbox := ui.NewHorizontalBox()

	form.Append("Blend Mode", makeBlendModeCb(&i.BlendMode), false)

	fileShower := ui.NewEntry()
	saveBtn := ui.NewButton("Select")
	saveBtn.OnClicked(func(*ui.Button) {
		filename = ui.OpenFile(win)
		fileShower.SetText(filename)
		i.Path = filename

		err := i.refreshCache()
		if err != nil {
			ui.MsgBoxError(win, "Error!", err.Error())
			return
		}
		recalcFunc(index)
	})

	hbox.Append(fileShower, true)
	hbox.Append(saveBtn, false)
	form.Append("File", hbox, false)
	return form, nil
}

func (i *PNGImporter) GetFrame(time int) (Frame, error) {
	if i.imgCache == nil {
		return Frame{}, errors.New("image hasn't been selected")
	}
	return Frame{
		Image:     i.imgCache,
		BlendMode: i.BlendMode,
	}, nil
}

func (i *PNGImporter) Cleanup() error       { return nil }
func (i *PNGImporter) Length() (int, error) { return -1, nil }
func (i *PNGImporter) Name() string {
	if i.Path == "" {
		return "PNG Clip"
	}
	return filepath.Base(i.Path)
}
