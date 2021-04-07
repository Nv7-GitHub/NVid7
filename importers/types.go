package importers

import (
	"image"

	"github.com/andlabs/ui"
)

type Importer interface {
	MakeUI(win *ui.Window, recalcFunc func()) (ui.Control, error)
	GetFrame(time int) (image.Image, error)
	Cleanup() error
	Length() (int, error) // -1 for infinite
	Name() string
}
