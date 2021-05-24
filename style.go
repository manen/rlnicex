package rlnicex

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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

type StyleConfig struct {
	Base    Style `json:"base"`
	Hovered Style `json:"hover"`
	Held    Style `json:"held"`
}

func (sc StyleConfig) Apply() error {
	err := sc.Base.setBaseStyle()
	if err != nil {
		return err
	}
	err = sc.Hovered.setHoveredStyle()
	if err != nil {
		return err
	}
	err = sc.Held.setHeldStyle()
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
	base, err := fixStyle(rsc.Base, defaultBaseStyle)
	if err != nil {
		return StyleConfig{}, err
	}
	hovered, err := fixStyle(rsc.Hovered, base)
	if err != nil {
		return StyleConfig{}, err
	}
	held, err := fixStyle(rsc.Held, base)
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

	// So you wanna add a new style element, huh?
	// Sure...
	//
	// You'll need to add it to this struct, in Style.ToMap, implement defaults
	// for it in defaultBaseStyle, defaultHoveredStyle, and defaultHeldStyle,
	// and make it work in fixStyle. Also, if you're planning on creating another
	// object inside of a style, you'll need to write some shit code like I did
	// with anyToColor. Have fun.
	//
	// Don't even get me started on adding a new style type, like if key x is
	// pressed, do something else. This is shit code.
}

func (s Style) setBaseStyle() error {
	log.Println("Setting base style", s)
	newStyle, err := fixStyle(s.ToMap(), defaultBaseStyle)
	baseStyle = newStyle
	return err
}
func (s Style) setHoveredStyle() error {
	log.Println("Setting hovered style", s)
	newStyle, err := fixStyle(s.ToMap(), baseStyle)
	hoveredStyle = newStyle
	return err
}
func (s Style) setHeldStyle() error {
	log.Println("Setting held style", s)
	newStyle, err := fixStyle(s.ToMap(), baseStyle)
	heldStyle = newStyle
	return err
}

func (s Style) ToMap() map[string]interface{} {
	m := map[string]interface{}{}

	m["backgroundColor"] = s.BackgroundColor
	m["fontColor"] = s.FontColor
	m["fontSize"] = s.FontSize
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
	baseStyle    Style
	hoveredStyle Style
	heldStyle    Style
)

func init() {
	defaultStyleConfig.Apply()
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

func fixStyle(raw map[string]interface{}, def Style) (Style, error) {
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

func getUsedStyle(c Hoverable, r Offset) Style {
	if c.IsHeld(r) {
		return heldStyle
	}
	if c.IsHovered(r) {
		return hoveredStyle
	}
	return baseStyle
}
