package rlnicex

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func getFinal(base rl.Rectangle, r Offset) rl.RectangleInt32 {
	final := base.ToInt32()
	final.X += int32(r.X)
	final.Y += int32(r.Y)

	return final
}

func DrawBorder(pos rl.RectangleInt32, style Style) {
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

type Hoverable interface {
	IsHovered(r Offset) bool
	IsHeld(r Offset) bool
}

func anyToColor(any interface{}) (rl.Color, error) {
	c, ok := any.(rl.Color)
	if ok {
		return c, nil
	}

	m, ok := any.(map[string]interface{})
	if !ok {
		return rl.Black, errors.New("color provided to anyToColor isn't a map[string]uint8")
	}

	r, ok := m["r"].(float64)
	if !ok {
		return rl.Black, errors.New("the r provided to anyToColor isn't a float64")
	}
	g, ok := m["g"].(float64)
	if !ok {
		return rl.Black, errors.New("the g provided to anyToColor isn't a float64")
	}
	b, ok := m["b"].(float64)
	if !ok {
		return rl.Black, errors.New("the b provided to anyToColor isn't a float64")
	}
	a, ok := m["a"].(float64)
	if !ok {
		return rl.Black, errors.New("the a provided to anyToColor isn't a float64")
	}

	return rl.Color{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}, nil
}
