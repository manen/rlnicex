package rlnicex

import (
	"encoding/json"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Style struct {
	// BackgroundColor is the color of the background of all widgets.
	BackgroundColor rl.Color `json:"backgroundColor"`

	// FontColor
	FontColor rl.Color `json:"fontColor"`
	// FontSize is the... font size.
	FontSize float64 `json:"fontSize"`
	// FontSpacing is the space between the letters. It might be called letter
	// spacing in other places
	FontSpacing float64 `json:"fontSpacing"`

	// BorderWidth is the... Width of the borders. It's not that hard
	BorderWidth float64 `json:"borderWidth"`
	// BorderColor is the color of the borders. I wish they were all this easy.
	BorderColor rl.Color `json:"borderColor"`
}

var DefaultStyle Style = Style{
	BackgroundColor: rl.White,

	FontColor:   rl.Black,
	FontSize:    16,
	FontSpacing: 2.4,

	BorderWidth: 2,
	BorderColor: rl.Color{
		R: 200,
		G: 200,
		B: 200,
		A: 255,
	},
}

var style Style = DefaultStyle

func SetStyle(s Style) {
	style = s
}

func LoadStyle(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	d := json.NewDecoder(file)
	d.Decode(&style)

	return nil
}
