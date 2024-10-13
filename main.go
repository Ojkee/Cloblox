package main

import "Cloblox/window"

func main() {
	mainWindow := window.NewWindow("Cloblox ", 900, 1800)
	mainWindow.MainLoop()
}
