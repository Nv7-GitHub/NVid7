package main

import (
	"image"

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

	x float64
}

type Sequencer struct {
	layerHeight float64

	clips     []importers.Importer
	clipDatas []clipData
	selected  int
	width     float64
	height    float64

	outWidth   int
	outHeight  int
	dragStartX float64
	dragStartY float64

	dragging bool

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
		x, y, w, h := s.calcClipPos(clip)
		p.AddRectangle(x, y, w, h)
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
		x, y, w, _ := s.calcClipPos(clip)
		textParams.Width = w
		textParams.String = ui.NewAttributedString(s.clips[i].Name())
		txt := ui.DrawNewTextLayout(&textParams)
		dp.Context.Text(txt, x+4, y+2)
		txt.Free()
	}
}

func (s *Sequencer) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
	// TODO: Make scrollable
	s.width = me.AreaWidth
	s.height = me.AreaHeight

	if me.Down == 1 {
		mouse := image.Pt(int(me.X), int(me.Y))
		for i, clip := range s.clipDatas {
			x, y, w, h := s.calcClipPos(clip)
			rct := image.Rect(int(x), int(y), int(w)+int(x), int(h)+int(y))
			if mouse.In(rct) {
				s.SetSelected(i)
				s.dragging = true
				s.dragStartX = me.X
				s.dragStartY = me.Y
				s.Update()
				return
			}
		}
	}

	if s.dragging {
		if len(me.Held) == 0 && me.Down == 0 {
			s.dragging = false
		}

		s.clipDatas[s.selected].x += (me.X - s.dragStartX) / s.width
		s.clipDatas[s.selected].startFrame = int(s.clipDatas[s.selected].x * float64(s.animationLength))
		s.dragStartX = me.X
		s.dragStartY = me.Y
		s.Update()
	}
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
func (s *Sequencer) ClipCount() int                                          { return len(s.clips) }
func (s *Sequencer) SetSelected(index int) {
	s.selected = index
	clip := s.clips[index]
	ui, err := clip.MakeUI(win, s.RecalcLength, index)
	handle(err)
	inspector.SetChild(ui)
	s.Update()
}
func (s *Sequencer) RecalcLength(index int) {
	clip := s.clips[index]
	length, err := clip.Length()
	handle(err)
	if length == -1 {
		length = s.animationLength / 5
	}
	clipD := s.clipDatas[index]
	if clipD.trimLength > length {
		clipD.trimLength = length
	}
	clipD.length = length
	s.clipDatas[index] = clipD
	s.Update()
}

func (s *Sequencer) SetOutputDimensions(sz image.Rectangle) {
	s.outWidth = sz.Dx()
	s.outHeight = sz.Dy()
}

func (s *Sequencer) Bounds() image.Rectangle {
	return image.Rect(0, 0, s.outWidth, s.outHeight)
}
