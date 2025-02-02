package iostate

import (
	"fmt"
	"os"
	"strings"

	"Cloblox/blocks"
	"Cloblox/graph"
)

// SaveToJson zapisuje strukturę grafu do pliku w formacie tekstowym
func SaveToJson(filename string, graph *graph.Graph) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Bloki
	blocksList := graph.GetAllBlocks()

	// --- Zapis bloków ---
	//_, err = file.WriteString("nodes: ")
	if err != nil {
		return fmt.Errorf("Error saving to file: %v", err)
	}

	var blocksText []string
	for _, block := range blocksList {
		var tag, content string

		// Określenie znacznika i zawartości
		switch b := block.(type) {
		case *blocks.StartBlock:
			tag = "start"
			content = "Start"
		case *blocks.StopBlock:
			tag = "stop"
			content = "Stop"
		case *blocks.IfBlock:
			tag = "if"
			content = b.GetConditionExprString()
		case *blocks.VariablesBlock:
			tag = "variable"
			varsMap := b.GetVars()
			if len(varsMap) == 0 {
				content = "NO VARIABLES"
			} else {
				var varContent []string
				for key, val := range varsMap {
					varContent = append(varContent, fmt.Sprintf("%s = %v", key, val))
				}
				content = strings.Join(varContent, "; ")
			}
		case *blocks.ActionBlock:
			tag = "action"
			content = b.GetConditionExprString()
		default:
			tag = "unknown"
			content = "UNKNOWN"
		}

		// Dodajemy blok do listy
		blocksText = append(blocksText, fmt.Sprintf("(%s)%d {\"%s\"}", tag, block.GetId(), content))
	}

	// Zapisujemy wszystkie bloki do pliku
	_, err = file.WriteString(strings.Join(blocksText, ","))
	if err != nil {
		return fmt.Errorf("Error saving to file: %v", err)
	}

	// --- Zapis połączeń ---
	_, err = file.WriteString("\nadjacency: \n")
	if err != nil {
		return fmt.Errorf("Error saving to file: %v", err)
	}

	// Przygotowujemy macierz sąsiedztwa
	N := len(blocksList)
	edgeMatrix := make([][]int, N)
	for i := 0; i < N; i++ {
		edgeMatrix[i] = make([]int, N)
	}

	// Połączenia
	addedConnections := make(map[string]bool)
	var previousBlock *blocks.Block

	// Iterujemy po każdym bloku w blocksList
	for _, currentBlock := range blocksList {
		inID := -1
		if previousBlock != nil {
			inID = (*previousBlock).GetId()
		}
		outID := currentBlock.GetId()

		// // Ignorujemy połączenia wychodzące z bloków Stop
		// if _, ok := currentBlock.(*blocks.StopBlock); ok {
		// 	break
		// }

		// Dodajemy połączenie między blokami
		if previousBlock != nil && inID != outID {
			connectionKey := fmt.Sprintf("%d,%d", inID, outID)
			if !addedConnections[connectionKey] {
				edgeMatrix[inID][outID] = 1 // Zwykłe połączenie
				addedConnections[connectionKey] = true
			}
		}

		// Sprawdzamy, czy blok jest typu 'IfBlock'
		if block, ok := currentBlock.(*blocks.IfBlock); ok {
			if block.GetNextTrue() != nil {
				trueBlock := block.GetNextTrue()
				connectionKey := fmt.Sprintf("%d,%d", outID, (*trueBlock).GetId())
				if !addedConnections[connectionKey] {
					edgeMatrix[outID][(*trueBlock).GetId()] = 3 // Połączenie prawdziwe (w prawo)
					addedConnections[connectionKey] = true
				}
			}

			if block.GetNextFalse() != nil {
				falseBlock := block.GetNextFalse()
				connectionKey := fmt.Sprintf("%d,%d", outID, (*falseBlock).GetId())
				if !addedConnections[connectionKey] {
					edgeMatrix[outID][(*falseBlock).GetId()] = 2 // Połączenie fałszywe (w lewo)
					addedConnections[connectionKey] = true
				}
			}
		} else {
			// Dodajemy standardowe połączenia (dla innych bloków)
			if previousBlock != nil && inID != outID {
				connectionKey := fmt.Sprintf("%d,%d", inID, outID)
				if !addedConnections[connectionKey] {
					edgeMatrix[inID][outID] = 1 // Zwykłe połączenie
					addedConnections[connectionKey] = true
				}
			}
		}

		// Przechodzimy do kolejnego bloku
		previousBlock = &currentBlock
	}

	// Zapisujemy macierz sąsiedztwa do pliku w wymaganym formacie
	_, err = file.WriteString("[\n")
	if err != nil {
		return fmt.Errorf("Error saving to file: %v", err)
	}

	for i, row := range edgeMatrix {
		_, err = file.WriteString(fmt.Sprintf("[%s]", strings.Join(intArrayToString(row), ",")))
		if err != nil {
			return fmt.Errorf("Error saving to file: %v", err)
		}
		if i < len(edgeMatrix)-1 {
			_, err = file.WriteString(",\n")
		} else {
			_, err = file.WriteString("\n")
		}
	}

	_, err = file.WriteString("]\n")
	if err != nil {
		return fmt.Errorf("Error saving to file: %v", err)
	}

	fmt.Println("Graph saved to JSON.")
	return nil
}

// Funkcja pomocnicza do zamiany tablicy intów na formatowany string
func intArrayToString(arr []int) []string {
	var result []string
	for _, v := range arr {
		result = append(result, fmt.Sprintf("%d", v))
	}
	return result
}
