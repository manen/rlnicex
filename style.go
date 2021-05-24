package rlnicex

import (
	"encoding/json"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/imdario/mergo"
)

type StyleConfig struct {
	Base    Style `json:"base"`
	Hovered Style `json:"hover"`
	Held    Style `json:"held"`
}

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

var DefaultBaseStyle Style = Style{
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
var DefaultHoverStyle Style = Style{
	BackgroundColor: rl.Color{
		R: 200,
		G: 200,
		B: 200,
		A: 255,
	},
}
var DefaultHeldStyle Style = Style{
	BackgroundColor: rl.Color{
		R: 200,
		G: 200,
		B: 200,
		A: 255,
	},
	FontColor: rl.Color{
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	},
}

var (
	baseStyle    Style
	hoveredStyle Style
	heldStyle    Style
)

func init() {
	SetBaseStyle(DefaultBaseStyle)
	SetHoveredStyle(DefaultHoverStyle)
	SetHeldStyle(DefaultHeldStyle)
}

func SetBaseStyle(s Style) error {
	newStyle, err := FixBaseStyle(s)
	if err != nil {
		return err
	}
	baseStyle = newStyle
	return nil
}
func SetHoveredStyle(s Style) error {
	newStyle, err := FixStyle(s)
	if err != nil {
		return err
	}
	hoveredStyle = newStyle
	return nil
}
func SetHeldStyle(s Style) error {
	newStyle, err := FixStyle(s)
	if err != nil {
		return err
	}
	heldStyle = newStyle
	return nil
}

func LoadStyle(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	newStyle := StyleConfig{}
	d := json.NewDecoder(file)
	d.Decode(&newStyle)

	err = SetBaseStyle(newStyle.Base)
	if err != nil {
		return err
	}
	err = SetHoveredStyle(newStyle.Hovered)
	if err != nil {
		return err
	}
	err = SetHeldStyle(newStyle.Held)
	if err != nil {
		return err
	}
	return nil
}

func FixBaseStyle(s Style) (Style, error) {
	err := mergo.Merge(&s, DefaultBaseStyle, mergo.WithTypeCheck)
	return s, err
}
func FixStyle(s Style) (Style, error) {
	err := mergo.Merge(&s, baseStyle, mergo.WithTypeCheck)
	log.Println(s)
	return s, err
}

func getUsedStyle(c Hoverable, r Offset) Style {
	if c.IsHeld(r) {
		return heldStyle
	}
	if c.IsHovered(r) {
		return hoveredStyle
	}
	return baseStyle
}
