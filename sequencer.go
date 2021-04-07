package main

import (
	"github.com/Nv7-Github/NVid7/importers"
	"github.com/andlabs/ui"
)

func NewSequencer() *Sequencer {
	return &Sequencer{
		layerHeight: 0.2,
	}
}

type clipData struct {
	yindex     int
	startFrame int
	length     int
	trimLength int
}

type Sequencer struct {
	layerHeight float64

	clips     []importers.Importer
	clipDatas []clipData
	selected  int
	width     float64
	height    float64

	animationLength int
	a               *ui.Area
}

func (s *Sequencer) Draw(a *ui.Area, dp *ui.AreaDrawParams) {
	// TODO: Make scrollable
	s.width = dp.AreaWidth
	s.height = dp.AreaHeight

	// Draw
	for i, clip := range s.clipDatas {
		p := ui.DrawNewPath(ui.DrawFillModeWinding)
		x := (float64(clip.startFrame) / float64(s.animationLength) * s.width) + 2
		w := float64(clip.trimLength) / float64(s.animationLength) * s.width
		y := (s.layerHeight * float64(clip.yindex) * s.height) + 2
		p.AddRectangle(x, y, w, s.layerHeight*s.height)
		p.End()
		dp.Context.Fill(p, &rectBrush)
		if s.selected == i {
			dp.Context.Stroke(p, &selectedBrush, &strokeParams)
		} else {
			dp.Context.Stroke(p, &outlineBrush, &strokeParams)
		}
		p.Free()
	}

	for i, clip := range s.clipDatas {
		x := float64(clip.startFrame) / float64(s.animationLength) * s.width
		y := s.layerHeight * float64(clip.yindex) * s.height
		textParams.Width = float64(clip.trimLength) / float64(s.animationLength) * s.width
		textParams.String = ui.NewAttributedString(s.clips[i].Name())
		txt := ui.DrawNewTextLayout(&textParams)
		dp.Context.Text(txt, x+4, y+2)
		txt.Free()
	}
}

func (s *Sequencer) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {

}

func (s *Sequencer) AddClip(clip importers.Importer) {
	s.clips = append(s.clips, clip)
	length, err := clip.Length()
	handle(err)
	if length == -1 {
		length = s.animationLength / 5
	}
	s.clipDatas = append(s.clipDatas, clipData{
		yindex:     len(s.clipDatas),
		startFrame: 0,
		length:     length,
		trimLength: length,
	})
}

func (s *Sequencer) SetAnimationLength(length int)                           { s.animationLength = length }
func (s *Sequencer) Update()                                                 { s.a.QueueRedrawAll() }
func (s *Sequencer) SetArea(a *ui.Area)                                      { s.a = a }
func (s *Sequencer) MouseCrossed(a *ui.Area, left bool)                      {}
func (s *Sequencer) DragBroken(a *ui.Area)                                   {}
func (s *Sequencer) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) { return false }
