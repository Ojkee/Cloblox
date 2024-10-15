package main

import "Cloblox/window"

func main() {
	mainWindow := window.NewWindow(
		"Cloblox",
		window.WINDOW_HEIGHT,
		window.WINDOW_WIDTH,
	)
	mainWindow.MainLoop()
}
