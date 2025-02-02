package iostate

import (
	"fmt"
	"os"
	"strings"

	"Cloblox/blocks"
	"Cloblox/graph"
)

// SaveToTxt zapisuje strukturę grafu do pliku tekstowego
func SaveToTxt(filename string, graph *graph.Graph) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, "Blocks:")

	// --- DEBUG: Sprawdź ilość bloczków ---
	blocksList := graph.GetAllBlocks()
	fmt.Printf("DEBUG: Found %d blocks in graph\n", len(blocksList))

	for _, block := range blocksList {
		var znacznik, content string

		// Określenie znacznika i zawartości
		switch b := block.(type) {
		case *blocks.StartBlock:
			znacznik = "h1"
			content = "Start"
		case *blocks.StopBlock:
			znacznik = "h2"
			content = "Stop"
		case *blocks.IfBlock:
			znacznik = "f"
			content = b.GetConditionExprString()
		case *blocks.VariablesBlock:
			znacznik = "v"
			// Pobieramy wartość zmiennej
			varsMap := b.GetVars()
			if len(varsMap) == 0 {
				content = "NO VARIABLES"
			} else {
				var varContent []string
				for key, val := range varsMap {
					varContent = append(varContent, fmt.Sprintf("%s = %v", key, val))
				}
				content = strings.Join(varContent, "; ") // Formatowanie jako lista zmiennych
			}
		case *blocks.ActionBlock:
			znacznik = "a"
			content = b.GetConditionExprString()
		default:
			znacznik = "?"
			content = "UNKNOWN"
		}

		// Zapis do pliku
		fmt.Fprintf(file, "<%s>\"%s\", %d</%s>\n", znacznik, content, block.GetId(), znacznik)

		// --- DEBUG: Wyświetlamy w terminalu ---
		fmt.Printf("DEBUG: Block ID: %d, Type: %s, Content: %s\n", block.GetId(), znacznik, content)
	}

	// Funkcja do obsługi połączeń
	fmt.Fprintln(file, "\nConnections:")

	// Zmienna do trzymania połączeń, żeby uniknąć dublowania
	addedConnections := make(map[string]bool)

	// Iterujemy po każdym bloku w blocksList
	var previousBlock *blocks.Block
	for _, currentBlock := range blocksList {
		inID := -1
		if previousBlock != nil {
			inID = (*previousBlock).GetId()
		}
		outID := currentBlock.GetId()

		// // Jeżeli to blok Stop, nie zapisuj połączeń wychodzących z niego
		// if _, ok := currentBlock.(*blocks.StopBlock); ok {
		// 	break // Zatrzymujemy iterację, gdy napotkamy StopBlock
		// }

		// Dodajemy połączenie między blokami
		if previousBlock != nil && inID != outID {
			connectionKey := fmt.Sprintf("%d,%d", inID, outID)
			if !addedConnections[connectionKey] {
				fmt.Fprintf(file, "<c>s%d,%d</c>\n", inID, outID) // Zwykłe połączenie
				addedConnections[connectionKey] = true
			}
		}

		// Sprawdzamy, czy blok jest typu 'IfBlock'
		if block, ok := currentBlock.(*blocks.IfBlock); ok {
			if block.GetNextTrue() != nil {
				trueBlock := block.GetNextTrue()
				if !addedConnections[fmt.Sprintf("%d,%d", outID, (*trueBlock).GetId())] {
					// Zapis połączenia dla prawdy (prawa strona)
					fmt.Fprintf(file, "<c>r%d,%d</c>\n", outID, (*trueBlock).GetId())
					addedConnections[fmt.Sprintf("%d,%d", outID, (*trueBlock).GetId())] = true
				}
			}

			if block.GetNextFalse() != nil {
				falseBlock := block.GetNextFalse()
				if !addedConnections[fmt.Sprintf("%d,%d", outID, (*falseBlock).GetId())] {
					// Zapis połączenia dla fałszu (lewa strona)
					fmt.Fprintf(file, "<c>l%d,%d</c>\n", outID, (*falseBlock).GetId())
					addedConnections[fmt.Sprintf("%d,%d", outID, (*falseBlock).GetId())] = true
				}
			}
		} else if _, ok := currentBlock.(*blocks.StopBlock); ok {
			// Blok Stop: Połączenia tylko przychodzące
			if previousBlock != nil && inID != outID {
				connectionKey := fmt.Sprintf("%d,%d", inID, outID)
				if !addedConnections[connectionKey] {
					fmt.Fprintf(file, "<c>s%d,%d</c>\n", outID, inID)
					addedConnections[connectionKey] = true
				}
			}
		} else {
			// Inne bloki: Zwykłe połączenie
			if previousBlock != nil && inID != outID {
				connectionKey := fmt.Sprintf("%d,%d", inID, outID)
				if !addedConnections[connectionKey] {
					fmt.Fprintf(file, "<c>s%d,%d</c>\n", inID, outID) // Zwykłe połączenie
					addedConnections[connectionKey] = true
				}
			}
		}

		// Przechodzimy do kolejnego bloku
		previousBlock = &currentBlock
	}

	return nil
}
