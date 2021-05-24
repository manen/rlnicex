package rlnicex

import rl "github.com/gen2brain/raylib-go/raylib"

type Label struct {
	Label    string
	Centered bool
	Pos      rl.Vector2
}

func NewLabel(label string, centered bool, x, y float64) Label {
	return Label{
		Label:    label,
		Centered: centered,
		Pos: rl.Vector2{
			X: float32(x),
			Y: float32(y),
		},
	}
}

func NewLabelSimple(label string) Label {
	return NewLabel(label, true, 0, 0)
}

func (l Label) RenderWithStyle(r Offset, style Style) {
	final := getFinal(rl.Rectangle{
		X: float32(l.Pos.X),
		Y: float32(l.Pos.Y),
	}, r)

	if l.Centered {
		o := rl.MeasureTextEx(rl.GetFontDefault(), l.Label, float32(style.FontSize), float32(style.FontSpacing))
		final.X -= int32(o.X / 2)
		final.Y -= int32(o.Y / 2)
	}

	rl.DrawTextEx(rl.GetFontDefault(), l.Label, rl.Vector2{
		X: float32(final.X),
		Y: float32(final.Y),
	}, float32(style.FontSize), float32(style.FontSpacing), style.FontColor)
}

func (l Label) Render(r Offset) {
	l.RenderWithStyle(r, baseStyle)
}
