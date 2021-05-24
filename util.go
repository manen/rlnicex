package rlnicex

import rl "github.com/gen2brain/raylib-go/raylib"

func getFinal(base rl.Rectangle, r Renderer) rl.RectangleInt32 {
	final := base.ToInt32()
	final.X += int32(r.Offset.X)
	final.Y += int32(r.Offset.Y)

	return final
}
