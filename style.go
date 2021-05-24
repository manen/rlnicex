package rlnicex

import rl "github.com/gen2brain/raylib-go/raylib"

type Style struct {
	// BackgroundColor is the color of the background of all widgets.
	BackgroundColor rl.Color

	// Font is the font **everything** is rendered with
	Font rl.Font
	// FontSize is the... font size.
	FontSize float64
	// FontSpacing is the space between the letters. It might be called letter
	// spacing in other places
	FontSpacing float64
}

var DefaultStyle Style = Style{
	BackgroundColor: rl.GetColor(0x141516ff),
	FontSize:        16,
	FontSpacing:     2.4,
}

var style Style = DefaultStyle

func SetStyle(s Style) {
	style = s
}
