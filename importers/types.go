package importers

import (
	"image"

	"github.com/andlabs/ui"
)

type BlendMode int

const (
	BlendModeAlphaOver BlendMode = 1 << iota
	BlendModeMultiply
)

var BlendModeNames = map[string]BlendMode{
	"Alpha Over": BlendModeAlphaOver,
	"Multiply":   BlendModeMultiply,
}

type Frame struct {
	BlendMode BlendMode
	Image     image.Image
}

type Importer interface {
	MakeUI(win *ui.Window, recalcFunc func(index int), index int) (ui.Control, error)
	GetFrame(time int) (Frame, error)
	Cleanup() error
	Length() (int, error) // -1 for infinite
	Name() string
}
