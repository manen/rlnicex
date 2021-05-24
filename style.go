package rlnicex

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StyleConfig struct {
	Base    Style `json:"base"`
	Hovered Style `json:"hover"`
	Held    Style `json:"held"`
}

func (sc StyleConfig) Apply() error {
	err := setBaseStyle(sc.Base)
	if err != nil {
		return err
	}
	err = setHoveredStyle(sc.Hovered)
	if err != nil {
		return err
	}
	err = setHeldStyle(sc.Held)
	if err != nil {
		return err
	}

	return nil
}

type rawStyleConfig struct {
	Base    map[string]interface{} `json:"base"`
	Hovered map[string]interface{} `json:"hover"`
	Held    map[string]interface{} `json:"held"`
}

func (rsc rawStyleConfig) ToStyleConfig() (StyleConfig, error) {
	base, err := fixBaseStyle(rsc.Base)
	if err != nil {
		return StyleConfig{}, err
	}
	hovered, err := _fixStyle(rsc.Hovered, base)
	if err != nil {
		return StyleConfig{}, err
	}
	held, err := _fixStyle(rsc.Held, base)
	if err != nil {
		return StyleConfig{}, err
	}

	return StyleConfig{
		Base:    base,
		Hovered: hovered,
		Held:    held,
	}, nil
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

func (s Style) ToMap() map[string]interface{} {
	m := map[string]interface{}{}

	m["backgroundColor"] = s.BackgroundColor
	m["fontColor"] = s.FontColor
	m["fontSpacing"] = s.FontSpacing
	m["borderWidth"] = s.BorderWidth
	m["borderColor"] = s.BorderColor

	return m
}

func (s Style) String() string {
	w := strings.Builder{}
	en := json.NewEncoder(&w)
	en.Encode(s)

	return w.String()
}

var (
	defaultBaseStyle Style = Style{
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
	defaultHoveredStyle Style = Style{
		BackgroundColor: rl.Color{
			R: 200,
			G: 200,
			B: 200,
			A: 255,
		},
	}
	defaultHeldStyle Style = Style{
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
	defaultStyleConfig StyleConfig = StyleConfig{
		Base:    defaultBaseStyle,
		Hovered: defaultHoveredStyle,
		Held:    defaultHeldStyle,
	}
)

var (
	baseStyle    Style
	hoveredStyle Style
	heldStyle    Style
)

func init() {
	defaultStyleConfig.Apply()
}

func setBaseStyle(s Style) error {
	log.Println("Setting base style", s)
	newStyle, err := fixBaseStyle(s.ToMap())
	baseStyle = newStyle
	return err
}
func setHoveredStyle(s Style) error {
	log.Println("Setting hovered style", s)
	newStyle, err := fixStyle(s.ToMap())
	hoveredStyle = newStyle
	return err
}
func setHeldStyle(s Style) error {
	log.Println("Setting held style", s)
	newStyle, err := fixStyle(s.ToMap())
	heldStyle = newStyle
	return err
}

func LoadStyle(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	newStyle := rawStyleConfig{}
	d := json.NewDecoder(file)
	d.Decode(&newStyle)

	sc, err := newStyle.ToStyleConfig()
	if err != nil {
		return err
	}
	return sc.Apply()
}

func fixBaseStyle(raw map[string]interface{}) (Style, error) {
	return _fixStyle(raw, defaultBaseStyle)
}
func fixStyle(raw map[string]interface{}) (Style, error) {
	return _fixStyle(raw, baseStyle)
}
func _fixStyle(raw map[string]interface{}, def Style) (Style, error) {
	s := Style{}

	log.Println("def:", def)

	// dogshit code ahead

	if raw["backgroundColor"] == nil {
		s.BackgroundColor = def.BackgroundColor
	} else {
		newColor, err := anyToColor(raw["backgroundColor"])
		if err != nil {
			return Style{}, err
		}
		s.BackgroundColor = newColor
	}
	if raw["fontColor"] == nil {
		s.FontColor = def.FontColor
	} else {
		newColor, err := anyToColor(raw["fontColor"])
		if err != nil {
			return Style{}, err
		}
		s.FontColor = newColor
	}
	if raw["fontSize"] == nil {
		s.FontSize = def.FontSize
	} else {
		s.FontSize = raw["fontSize"].(float64)
	}
	if raw["fontSpacing"] == nil {
		s.FontSpacing = def.FontSpacing
	} else {
		s.FontSpacing = raw["fontSpacing"].(float64)
	}
	if raw["borderWidth"] == nil {
		s.BorderWidth = def.BorderWidth
	} else {
		s.BorderWidth = raw["borderWidth"].(float64)
	}
	if raw["borderColor"] == nil {
		s.BorderColor = def.BorderColor
	} else {
		newColor, err := anyToColor(raw["borderColor"])
		if err != nil {
			return Style{}, err
		}
		s.BorderColor = newColor
	}

	return s, nil
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

func getUsedStyle(c Hoverable, r Offset) Style {
	if c.IsHeld(r) {
		return heldStyle
	}
	if c.IsHovered(r) {
		return hoveredStyle
	}
	return baseStyle
}
