package rlnicex_test

import (
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
	rlx "github.com/manen/rlnicex"
)

func TestFull(t *testing.T) {
	rl.InitWindow(690, 420, "Test") // I can't keep memorizing these small 16:9 ratios
	rl.SetTargetFPS(60)

	rlx.LoadStyle("./test_assets/style.json")

	r := rlx.NewOffset(0, 0)
	btn := rlx.NewButton(rlx.NewLabelSimple("Nice"), 10, 10, 140, 40)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		r.Y += 0.33
		btn.Render(r)

		rl.EndDrawing()
	}
}
