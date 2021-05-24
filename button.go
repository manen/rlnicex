package rlnicex

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	Label Label
	Pos   rl.Rectangle
}

func NewButton(label Label, x, y, w, h float64) Button {
	return Button{
		Label: label,
		Pos:   rl.NewRectangle(float32(x), float32(y), float32(w), float32(h)),
	}
}

func (b Button) Render(r Offset) {
	final := getFinal(b.Pos, r)

	// Draw the background
	rl.DrawRectangle(final.X, final.Y, final.Width, final.Height, style.BackgroundColor)
	// Render the label
	b.Label.Render(r.Sub(float64(b.Pos.X)+float64(b.Pos.Width)/2, float64(b.Pos.Y)+float64(b.Pos.Height)/2))

	DrawBorder(final)
}
