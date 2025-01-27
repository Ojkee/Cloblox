package iostate

import (
	"fmt"
	"os"

	"Cloblox/shapes"
)

func SaveToTxt(path string, blocks []shapes.Shape, connections []shapes.Connection) error {
	file, err := os.Create(path)
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
		return fmt.Errorf("Error saving to file: %v", err)
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

	// Mapa wezlow na indeksy w macierzy
	nodeIndexMap := make(map[int]int)
	for i, block := range blocks {
		nodeIndexMap[block.GetBlockId()] = i
	}

	for _, conn := range connections {
		inID := conn.GetInShapeId()
		outID := conn.GetOutShapeId()

		inIndex, okIn := nodeIndexMap[inID]

		if okIn {
			inShape := blocks[inIndex]
			var prefix string = "s"

			// Sprawdzenie, czy blok wejściowy to IF
			if inShape.GetType() == shapes.IF {
				if conn.IsCloserToRigth() {
					prefix = "l" // Lewa strona - nie
				} else {
					prefix = "r" // Prawa strona - tak
				}
			}

			// Zapis do pliku z odpowiednim prefiksem
			_, err := file.WriteString(fmt.Sprintf("<c>%s%d,%d</c>\n", prefix, inID, outID))
			if err != nil {
				fmt.Printf("Error saving connection %d -> %d: %v\n", inID, outID, err)
				return fmt.Errorf("error saving connection: %v", err)
			}
		} else {
			// Obsługa błędów, gdy indeksy nie zostaną znalezione
			fmt.Printf("Error: indices not found for connection %d -> %d\n", inID, outID)
		}
	}
	return nil
}
