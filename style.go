package rlnicex

import rl "github.com/gen2brain/raylib-go/raylib"

type Style struct {
	// BackgroundColor is the color of the background of all widgets.
	BackgroundColor rl.Color

	// FontSize is the... font size.
	FontSize float64
	// FontSpacing is the space between the letters. It might be called letter
	// spacing in other places
	FontSpacing float64

	// BorderWidth is the... Width of the borders. It's not that hard
	BorderWidth float64
	// BorderColor is the color of the borders. I wish they were all this easy.
	BorderColor rl.Color
}

var DefaultStyle Style = Style{
	BackgroundColor: rl.GetColor(0x141516ff),

	FontSize:    16,
	FontSpacing: 2.4,

	BorderWidth: 1,
	BorderColor: rl.GetColor(0x343536ff),
}

var style Style = DefaultStyle

func SetStyle(s Style) {
	style = s
}
