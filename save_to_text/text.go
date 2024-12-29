package save_to_text

import (
	"fmt"
	"os"

	"Cloblox/shapes"
	"Cloblox/window"
)

func SaveToTxt(filename string, blocks []shapes.Shape, connections []*window.Connection) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Error creating new file: %v\n", err)
	}
	defer file.Close()

	tags := map[shapes.SHAPE_TYPE]string{
		shapes.START:    "h1",
		shapes.STOP:     "h2",
		shapes.IF:       "f",
		shapes.VARIABLE: "v",
		shapes.ACTION:   "a",
	}

	_, err = file.WriteString("Blocks:\n")
	if err != nil {
		return fmt.Errorf("błąd zapisu do pliku: %v", err)
	}
	for _, block := range blocks {
		content := block.GetContent()
		blockID := block.GetBlockId()
		blockType := block.GetType()
		tag, exists := tags[blockType]
		if !exists {
			tag = "unknown"
		}
		for _, line := range content {
			_, err = file.WriteString(fmt.Sprintf("<%s>\"%s\",%d</%s>\n", tag, line, blockID, tag))
			if err != nil {
				return fmt.Errorf("Error saving to file: %v", err)
			}
		}
	}

	_, err = file.WriteString("\nConnections:\n")
	if err != nil {
		return fmt.Errorf("Error saving to file: %v\n", err)
	}

	for _, conn := range connections {
		startID := conn.GetInShapeId()
		stopID := conn.GetOutShapeId()
		_, err = file.WriteString(fmt.Sprintf("<c>%d,%d</c>\n", startID, stopID))
		if err != nil {
			return fmt.Errorf("Error saving to file: %v\n", err)
		}
	}
	return nil
}
