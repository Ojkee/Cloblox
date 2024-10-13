package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Window struct {
	name   string
	height int32
	width  int32

	backgroundColor rl.Color
}

func NewWindow(name string, height, width int32) *Window {
	return &Window{
		name:   name,
		height: height,
		width:  width,

		backgroundColor: rl.NewColor(51, 51, 51, 255),
	}
}

func (window *Window) MainLoop() {
	rl.InitWindow(window.width, window.height, "Cloblox")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.MaximizeWindow()
		rl.ClearBackground(window.backgroundColor)

		rl.EndDrawing()
	}
}
