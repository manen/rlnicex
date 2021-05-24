package rlnicex

import rl "github.com/gen2brain/raylib-go/raylib"

type Offset rl.Vector2

func NewOffset(x, y float64) Offset {
	return Offset{
		X: float32(x),
		Y: float32(y),
	}
}

func (r Offset) Sub(x, y float64) Offset {
	return Offset{
		X: float32(float64(r.X) + x),
		Y: float32(float64(r.Y) + y),
	}
}
