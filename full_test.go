package rlnicex_test

import (
	"log"
	"testing"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	rlx "github.com/manen/rlnicex"
)

func TestFull(t *testing.T) {
	rl.InitWindow(690, 420, "Test") // I can't keep memorizing these small 16:9 ratios
	rl.SetTargetFPS(60)

	err := rlx.LoadStyle("./test_assets/style.json")
	if err != nil {
		log.Panicln(err)
	}

	r := rlx.NewOffset(0, 0)
	btn := rlx.NewButton(rlx.NewLabelSimple("Nice"), 10, 10, 140, 40)
	originalLabel := btn.Label.Label

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		btn.Render(r)
		if btn.IsClicked(r) && btn.Label.Label == originalLabel {
			go func() {
				btn.Label.Label = "Clicked!"
				time.Sleep(500 * time.Millisecond)
				btn.Label.Label = originalLabel
			}()
		}

		rl.DrawFPS(4, 4)

		rl.EndDrawing()
	}
}
