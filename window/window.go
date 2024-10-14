package window

import rl "github.com/gen2brain/raylib-go/raylib"

type Window struct {
	name   string
	height int32
	width  int32

	backgroundColor rl.Color
	fontColor       rl.Color

	buildingShapes []ShapeDrawable
	shapes         []ShapeDrawable
}

func NewWindow(name string, height, width int32) *Window {
	return &Window{
		name:   name,
		height: height,
		width:  width,

		backgroundColor: BACKGROUND_COLOR,
		fontColor:       rl.NewColor(255, 248, 231, 255),

		buildingShapes: initBuildingShapes(width, height),
		shapes:         make([]ShapeDrawable, 0),
	}
}

func initBuildingShapes(width, height int32) []ShapeDrawable {
	offsetX := float32(width/2 + 10)
	gap := float32(SHAPE_HEIGHT + 16)
	offsetY := float32(height)/2.0 - gap*2.5
	buildingShapes := make([]ShapeDrawable, 0)
	buildingShapes = append(buildingShapes, NewStartShape(offsetX, offsetY))
	offsetY += gap
	buildingShapes = append(buildingShapes, NewVariableShape(offsetX, offsetY))
	offsetY += gap
	buildingShapes = append(buildingShapes, NewIfShape(offsetX, offsetY))
	offsetY += gap
	buildingShapes = append(buildingShapes, NewPrintShape(offsetX, offsetY))
	offsetY += gap
	buildingShapes = append(buildingShapes, NewStopShape(offsetX, offsetY))
	return buildingShapes
}

func (window *Window) MainLoop() {
	rl.InitWindow(window.width, window.height, "Cloblox")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !(rl.WindowShouldClose() || rl.IsKeyPressed(rl.KeyQ)) {
		window.draw()
	}
}

func (window *Window) draw() {
	rl.BeginDrawing()
	rl.MaximizeWindow()

	rl.ClearBackground(window.backgroundColor)

	rl.DrawLine(window.width/2, 0, window.width/2, window.height, rl.NewColor(255, 248, 231, 255))
	for _, shape := range window.buildingShapes {
		shape.Draw()
	}
	for _, shape := range window.shapes {
		shape.Draw()
	}

	rl.EndDrawing()
}
