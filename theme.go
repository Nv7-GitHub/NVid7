package main

import "github.com/andlabs/ui"

var rectBrush = ui.DrawBrush{
	Type: ui.DrawBrushTypeSolid,
	R:    1,
	B:    1,
	G:    1,
	A:    1,
}
var outlineBrush = ui.DrawBrush{
	Type: ui.DrawBrushTypeSolid,
	R:    0.9,
	B:    0.9,
	G:    0.9,
	A:    1,
}
var selectedBrush = ui.DrawBrush{
	Type: ui.DrawBrushTypeSolid,
	R:    0,
	B:    1,
	G:    1,
	A:    1,
}
var textParams = ui.DrawTextLayoutParams{
	DefaultFont: ui.NewFontButton().Font(),
	Align:       ui.DrawTextAlignCenter,
}

var strokeParams = ui.DrawStrokeParams{
	Cap:       ui.DrawLineCapRound,
	Join:      ui.DrawLineJoinRound,
	Thickness: 2,
}
