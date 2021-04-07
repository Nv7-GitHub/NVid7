package main

import (
	"github.com/Nv7-Github/NVid7/importers"
	"github.com/andlabs/ui"
)

func NewSequencer() *Sequencer {
	return &Sequencer{}
}

type Sequencer struct {
	clips []importers.Importer

	a *ui.Area
}

func (s *Sequencer) Draw(a *ui.Area, dp *ui.AreaDrawParams) {

}

func (s *Sequencer) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {

}

func (s *Sequencer) AddClip(clip importers.Importer) {
	s.clips = append(s.clips, clip)
}

func (s *Sequencer) Update()                                                 { s.a.QueueRedrawAll() }
func (s *Sequencer) SetArea(a *ui.Area)                                      { s.a = a }
func (s *Sequencer) MouseCrossed(a *ui.Area, left bool)                      {}
func (s *Sequencer) DragBroken(a *ui.Area)                                   {}
func (s *Sequencer) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) { return false }
