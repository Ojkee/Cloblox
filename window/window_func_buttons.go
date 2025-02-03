package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"Cloblox/functools"
	"Cloblox/settings"
)

type FuncButton struct {
	content string
	rect    rl.Rectangle
}

func NewFuncButton(content string, rect rl.Rectangle) *FuncButton {
	return &FuncButton{
		content: content,
		rect:    rect,
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
	contentSize := functools.TextSizeEx(fb.content)
	contentPos := rl.NewVector2(
		fb.rect.X+(fb.rect.Width-contentSize.X)/2,
		fb.rect.Y+(fb.rect.Height-contentSize.Y)/2,
	)
	rl.DrawTextEx(
		settings.FONT,
		fb.content,
		contentPos,
		float32(settings.FONT_SIZE),
		settings.FONT_SPACING,
		settings.FONT_COLOR,
	)
}

func (fb *FuncButton) InRect(mousePos rl.Vector2) bool {
	return (mousePos.X >= fb.rect.X && mousePos.X <= fb.rect.X+fb.rect.Width) &&
		(mousePos.Y >= fb.rect.Y && mousePos.Y <= fb.rect.Y+fb.rect.Height)
}
