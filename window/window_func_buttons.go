package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/functools"
	"Cloblox/settings"
)

type FuncButton struct {
	content    string
	rect       rl.Rectangle
	contentPos rl.Vector2
}

func NewFuncButton(content string, rect rl.Rectangle) *FuncButton {
	contentSize := functools.TextSizeEx(content)
	_contentPos := rl.NewVector2(
		rect.X+(rect.Width-contentSize.X)/2,
		rect.Y+(rect.Height-contentSize.Y)/2,
	)
	return &FuncButton{
		content:    content,
		rect:       rect,
		contentPos: _contentPos,
	}
}

func (fb *FuncButton) Draw() {
	var roundness float32 = 0.5
	var segments int32 = 10
	rl.DrawRectangleRounded(
		fb.rect,
		roundness,
		segments,
		settings.SIMULATE_BUTTON_COLOR,
	)
	rl.DrawTextEx(
		settings.FONT,
		fb.content,
		fb.contentPos,
		float32(settings.FONT_SIZE),
		settings.FONT_SPACING,
		settings.FONT_COLOR,
	)
}

func (fb *FuncButton) InRect(mousePos rl.Vector2) bool {
	return (mousePos.X >= fb.rect.X && mousePos.X <= fb.rect.X+fb.rect.Width) &&
		(mousePos.Y >= fb.rect.Y && mousePos.Y <= fb.rect.Y+fb.rect.Height)
}
