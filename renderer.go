package rlnicex

import rl "github.com/gen2brain/raylib-go/raylib"

type Renderer struct {
	Offset rl.Vector2
}

func NewRenderer(x, y float64) Renderer {
	return Renderer{
		Offset: rl.Vector2{
			X: float32(x),
			Y: float32(y),
		},
	}
}

func (r Renderer) Sub(x, y float64) Renderer {
	return Renderer{
		Offset: rl.Vector2{
			X: float32(float64(r.Offset.X) + x),
			Y: float32(float64(r.Offset.Y) + y),
		},
	}
}
