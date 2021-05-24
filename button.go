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

func (b Button) IsHovered(r Offset) bool {
	final := getFinal(b.Pos, r)
	return rl.CheckCollisionPointRec(rl.GetMousePosition(), final.ToFloat32())
}

func (b Button) IsClicked(r Offset) bool {
	return b.IsHovered(r) && rl.IsMouseButtonReleased(rl.MouseLeftButton)
}

func (b Button) IsHeld(r Offset) bool {
	return b.IsHovered(r) && rl.IsMouseButtonDown(rl.MouseLeftButton)
}

func (b Button) Render(r Offset) {
	final := getFinal(b.Pos, r)
	style := getUsedStyle(b, r)

	// Draw the background
	rl.DrawRectangle(final.X, final.Y, final.Width, final.Height, style.BackgroundColor)
	// Render the label
	b.Label.RenderWithStyle(r.Sub(float64(b.Pos.X)+float64(b.Pos.Width)/2, float64(b.Pos.Y)+float64(b.Pos.Height)/2), style)

	DrawBorder(final, style)
}
