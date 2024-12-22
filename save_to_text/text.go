package functools

import (
	"Cloblox/shapes"
	"Cloblox/window"
	"fmt"
	"os"
)

// SaveToTxt zapisuje schemat blokowy do pliku .txt
func SaveToTxt(filename string, blocks []shapes.Shape, connections []*window.Connection) error {

	// Otwórzenie pliku do zapisu
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("błąd tworzenia pliku: %v", err)
	}
	defer file.Close()

	// Mapowanie typów bloczków na znaczniki
	tags := map[shapes.SHAPE_TYPE]string{
		shapes.START:    "h1",
		shapes.STOP:     "h2",
		shapes.IF:       "f",
		shapes.VARIABLE: "v",
		shapes.ACTION:   "a",
	}

	// Zapisz bloki
	_, err = file.WriteString("Bloki:\n")
	if err != nil {
		return fmt.Errorf("błąd zapisu do pliku: %v", err)
	}
	for _, block := range blocks {
		// Pobierz zawartość i typ bloczka
		content := block.GetContent()
		blockID := block.GetBlockId()
		blockType := block.GetType()

		// Pobierz odpowiedni znacznik
		tag, exists := tags[blockType]
		if !exists {
			tag = "unknown" // Domyślny znacznik dla nieznanego typu
		}

		// Sformatuj zawartość bloczka jako tekst w odpowiednich znacznikach
		for _, line := range content {
			_, err = file.WriteString(fmt.Sprintf("<%s>\"%s\",%d</%s>\n", tag, line, blockID, tag))
			if err != nil {
				return fmt.Errorf("błąd zapisu bloku do pliku: %v", err)
			}
		}
	}

	// Zapisz połączenia
	_, err = file.WriteString("\nPołączenia:\n")
	if err != nil {
		return fmt.Errorf("błąd zapisu do pliku: %v", err)
	}

	for _, conn := range connections {
		startID := conn.GetInShapeId() // ID bloku początkowego
		stopID := conn.GetOutShapeId() // ID bloku końcowego
		_, err = file.WriteString(fmt.Sprintf("<c>%d,%d</c>\n", startID, stopID))
		if err != nil {
			return fmt.Errorf("błąd zapisu połączenia do pliku: %v", err)
		}
	}

	return nil
}
