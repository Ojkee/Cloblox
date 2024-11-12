package settings

import rl "github.com/gen2brain/raylib-go/raylib"

// WINDOW
const (
	WINDOW_HEIGHT = 800
	WINDOW_WIDTH  = 1400
)

var (
	FONT      rl.Font
	FONT_SIZE int32 = 16
)

// COLORS
var (
	BACKGROUND_COLOR       = rl.NewColor(51, 51, 51, 255)
	FONT_COLOR             = rl.NewColor(255, 248, 231, 255)
	FONT_COLOR_TRANSPARENT = rl.NewColor(255, 248, 231, 100)
	IF_COLOR               = rl.NewColor(171, 110, 164, 255)
	IF_COLOR_TRANSPARENT   = rl.NewColor(171, 110, 164, 100)
	START_STOP_COLOR       = rl.NewColor(118, 105, 126, 255)
	VARIABLE_COLOR         = rl.NewColor(153, 141, 153, 255)
	PRINT_COLOR            = rl.NewColor(102, 96, 102, 255)
	CONNECTION_COLOR       = rl.NewColor(77, 74, 77, 255)
	BLOCK_POSITIVE         = rl.NewColor(108, 123, 41, 255)
	BLOCK_NEGATIVE         = rl.NewColor(179, 34, 53, 255)
	HIGHLIGHT_COLOR        = rl.NewColor(204, 199, 186, 255)
)

// SHAPES
var (
	SHAPE_MIN_WIDTH  float32 = 96
	SHAPE_MIN_HEIGHT float32 = 32
	SHAPE_TEXT_GAP   float32 = 10

	MAX_CONTENT_LINES int     = 4
	MAX_CONTENT_CHARS int     = 30
	MARGIN_HORIZONTAL float32 = 10
	MARGIN_VERTICAL   float32 = 10

	HIGHLIGHT_PAD float32 = 4
)

// DEBUG
var (
	DEBUG_BLOCKS_POINTERS = false
	DEBUG_SHAPE_CONTENT   = true
)
