package settings

import rl "github.com/gen2brain/raylib-go/raylib"

// WINDOW
const (
	WINDOW_HEIGHT = 800
	WINDOW_WIDTH  = 1400
)

// DEBUG
var (
	DEBUG_BLOCKS_POINTERS             = false
	DEBUG_SHAPE_CONTENT               = false
	DEBUG_DIAGRAM_DETAILS             = false
	DEBUG_DIAGRAM_WINDOW_SIDE_DETAILS = true
	DEBUG_ERRORS                      = false
)

// FONT
var (
	FONT         rl.Font
	FONT_SIZE    int32   = 16
	FONT_SPACING float32 = 1
	FONT_PATH    string  = "fonts/Metropolis-Medium.otf"
)

// COLORS
var (
	DEBUG_COLOR                           = rl.NewColor(253, 87, 87, 255)
	BACKGROUND_COLOR                      = rl.NewColor(51, 51, 51, 255)
	FONT_COLOR                            = rl.NewColor(255, 248, 231, 255)
	FONT_ERROR_COLOR                      = rl.NewColor(243, 170, 154, 255)
	FONT_ERROR_STRONG_COLOR               = rl.NewColor(230, 92, 76, 255)
	IF_COLOR                              = rl.NewColor(171, 110, 164, 255)
	START_STOP_COLOR                      = rl.NewColor(118, 105, 126, 255)
	VARIABLE_COLOR                        = rl.NewColor(153, 141, 153, 255)
	PRINT_COLOR                           = rl.NewColor(102, 96, 102, 255)
	CONNECTION_COLOR                      = rl.NewColor(77, 74, 77, 255)
	HIGHLIGHT_COLOR                       = rl.NewColor(204, 199, 186, 255)
	POSITIVE_VAL_COLOR                    = rl.NewColor(54, 227, 130, 255)
	NEGATIVE_VAL_COLOR                    = rl.NewColor(165, 61, 40, 255)
	HELP_OUTER_BORDER_COLOR               = rl.NewColor(81, 81, 81, 255)
	HELP_INNER_BORDER_COLOR               = rl.NewColor(31, 31, 31, 255)
	BUTTON_COLOR                          = rl.NewColor(31, 31, 31, 255)
	BUTTON_COLOR_SELECTED                 = rl.NewColor(71, 71, 71, 255)
	FUNCTION_BUTTON_COLOR                 = BUTTON_COLOR
	SIMULATE_BUTTON_COLOR                 = rl.NewColor(31, 31, 31, 255)
	SIMULATE_BUTTON_ATTRIB_COLOR          = rl.NewColor(100, 100, 100, 255)
	SIMULATE_BUTTON_ATTRIB_COLOR_SELECTED = FONT_COLOR
)

// SHAPES
var (
	SHAPE_MIN_WIDTH  float32 = 96
	SHAPE_MIN_HEIGHT float32 = 32
	SHAPE_TEXT_GAP   float32 = 10

	MAX_CONTENT_LINES int     = 4
	MAX_CONTENT_CHARS int     = 50
	MARGIN_HORIZONTAL float32 = 10
	MARGIN_VERTICAL   float32 = 10

	HIGHLIGHT_PAD float32 = 4
)

// CONSOLE
const (
	CONSOLE_WIDTH               = WINDOW_WIDTH / 2
	CONSOLE_HEIGHT              = WINDOW_HEIGHT / 4
	CONSOLE_PREFIX              = ">"
	CONSOLE_PREFIX_STRONG_ERROR = "!!"
	CONSOLE_MAX_LINES           = 10
	CONSOLE_MAX_LINES_HISTORY   = 64
	CONSOLE_MAX_LINE_WIDTH      = CONSOLE_WIDTH - 50
	CONSOLE_MARGIN              = 10
	CONSOLE_BORDER_WIDTH        = 2
)

// BUTTONS
const (
	BUTTON_MARGIN float32 = 10
	BUTTON_WIDTH  float32 = 60
	BUTTON_HEIGHT float32 = 30
	BUTTON_X      float32 = WINDOW_WIDTH/2 - BUTTON_WIDTH - BUTTON_MARGIN
	BUTTON_Y      float32 = BUTTON_MARGIN
	BUTTON_GAP    float32 = 4
	BUTTON_LIMIT  int     = 10
)

// SIMULATION OPTIONS
const (
	SIMULATION_TIME_STEP_MS     = 500
	SIMULATION_TIME_STEP_MS_MIN = 40
	SIMULATION_TIME_STEP_MS_MAX = 4000
)

// SIMULATION SLICE
const (
	SLICE_UP_MARGIN     = 10
	SLICE_BOTTON_MARGIN = 10
	SLICE_HIGH_POS      = WINDOW_HEIGHT - SLICE_BOTTON_MARGIN - CONSOLE_HEIGHT
	SLICE_LOW_POS       = SLICE_UP_MARGIN
)

// SIMULATE BUTTONS
const (
	SIMULATE_BUTTON_GAP    = 10
	SIMULATE_BUTTON_WIDTH  = 70
	SIMULATE_BUTTON_HEIGHT = 40
	SIMULATE_BUTTON_OFFSET = 40
	SIMULATE_BUTTON_POS_X  = WINDOW_WIDTH/2 - 3*SIMULATE_BUTTON_WIDTH - 2*SIMULATE_BUTTON_GAP - SIMULATE_BUTTON_OFFSET
	SIMULATE_BUTTON_POS_Y  = WINDOW_HEIGHT - CONSOLE_HEIGHT - SIMULATE_BUTTON_HEIGHT

	SIMULATE_BUTTON_PAUSE_SIZE    = 16
	SIMULATE_BUTTON_TRIANGLE_SIZE = 16
)

const (
	PATH_TXT           = "records/saves/save1.txt"
	PATH_JSON          = "records/saves/save1.json"
	PATH_PYTHON        = "records/code/save1.py"
	PATH_PDF           = "records/pdfs/save1.pdf"
	PATH_PDF_TEMP_JSON = "records/pdfs/save1.json"
)
