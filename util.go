package main

func (s *Sequencer) calcClipPos(clip clipData) (x, y, w, h float64) {
	x = (float64(clip.startFrame) / float64(s.animationLength) * s.width) + 2
	w = float64(clip.trimLength) / float64(s.animationLength) * s.width
	y = (s.layerHeight * float64(clip.yindex) * s.height) + 2
	h = s.layerHeight * s.height
	return
}
