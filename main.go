package main

import "Cloblox/window"

func main() {
	mainWindow := window.NewWindow(
		"Cloblox",
		window.WINDOW_WIDTH,
		window.WINDOW_HEIGHT,
	)
	mainWindow.MainLoop()
}
