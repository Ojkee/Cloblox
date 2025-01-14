package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/settings"
)

// Prints help when mouse hoovers over certain shape
func (window *Window) drawHelp(mousePos *rl.Vector2) {
	helpInfo := window.defaultHelpMessage()
	window.drawHelpBackgourd()
	for _, shape := range window.buildingShapes {
		if rl.CheckCollisionPointRec(*mousePos, shape.GetRect()) {
			helpInfo = shape.Info()
		}
	}

	rl.DrawTextEx(
		settings.FONT,
		helpInfo,
		rl.NewVector2(16, 16),
		float32(settings.FONT_SIZE),
		settings.FONT_SPACING,
		settings.FONT_COLOR,
	)
}

func (window *Window) drawHelpBackgourd() {
	var margin float32 = 10
	var borderWidth float32 = 2
	var roundFactor float32 = 0.05
	rl.DrawRectangleRounded(
		rl.NewRectangle(
			margin,
			margin,
			settings.WINDOW_WIDTH/2-2*margin,
			settings.WINDOW_HEIGHT-2*margin,
		),
		roundFactor,
		16,
		settings.HELP_OUTER_BORDER_COLOR,
	)
	rl.DrawRectangleRounded(
		rl.NewRectangle(
			margin+borderWidth,
			12,
			settings.WINDOW_WIDTH/2-2*(margin+borderWidth),
			settings.WINDOW_HEIGHT-2*(margin+borderWidth),
		),
		roundFactor,
		8,
		settings.HELP_INNER_BORDER_COLOR,
	)
}

func (window *Window) defaultHelpMessage() string {
	retVal := `
	Controls 
  Press key suggested within [ * ] to enter mode.
  In upper right corner You can find information about
  current mode.

	Building
  [ B ] , [ 1 ], [ Esc ]  
	Allows user to place new blocks (left click) and make
  connections between blocks (right click).

	Insering
  [ I ] , [ 2 ]
	Allows user to insert content such as: 
	- If statement
	- New Variables
	- Math operations
  Insertion supports pasting from clipboard via ctrl + V.

	Removing                       
  [ R ] , [ 3 ]
	Allows user to remove blocks with its connections (left click).

	Simulating   
  [ S ] , [ 4 ]
	Starts simulation mode, where user can select which array variable
  program should visualize and whether algorithm makes step by step
  simulation or moves automatically.


  Console
  Appears in simulation mode and shows every print action,
  strong and weak errors.
  - Strong error - '!! [red]' - Doesn't let simulation to start.
  - Weak error - '> [orange]' - Lets simulation to start.
    May couse Strong error later on.
  `
	return retVal
}
