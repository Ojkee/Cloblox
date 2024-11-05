package main

import (
	"Cloblox/settings"
	"Cloblox/window"
)

func main() {
	mainWindow := window.NewWindow(
		"Cloblox",
		settings.WINDOW_HEIGHT,
		settings.WINDOW_WIDTH,
	)
	mainWindow.MainLoop()
}
