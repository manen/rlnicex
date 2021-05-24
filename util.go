package rlnicex

import rl "github.com/gen2brain/raylib-go/raylib"

func getFinal(base rl.Rectangle, r Offset) rl.RectangleInt32 {
	final := base.ToInt32()
	final.X += int32(r.X)
	final.Y += int32(r.Y)

	return final
}

func DrawBorder(pos rl.RectangleInt32) {
	if style.BorderWidth <= 0 {
		return
	}

	// Top
	rl.DrawRectangle(pos.X-int32(style.BorderWidth/2), pos.Y-int32(style.BorderWidth/2), pos.Width, int32(style.BorderWidth), style.BorderColor)
	// Left
	rl.DrawRectangle(pos.X-int32(style.BorderWidth/2), pos.Y-int32(style.BorderWidth/2), int32(style.BorderWidth), pos.Height, style.BorderColor)
	// Bottom
	rl.DrawRectangle(pos.X-int32(style.BorderWidth/2), pos.Y+pos.Height-int32(style.BorderWidth/2), pos.Width, int32(style.BorderWidth), style.BorderColor)
	// Right
	rl.DrawRectangle(pos.X+pos.Width-int32(style.BorderWidth/2), pos.Y-int32(style.BorderWidth/2), int32(style.BorderWidth), pos.Height+int32(style.BorderWidth), style.BorderColor)
}
