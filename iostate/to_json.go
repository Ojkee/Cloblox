package iostate

import (
	"fmt"
	"os"
	"strings"

	"Cloblox/shapes"
)

func SaveToJson(filename string, blocks []shapes.Shape, connections []shapes.Connection) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Error creating file: %v", err)
	}
	defer file.Close()

	tags := map[shapes.SHAPE_TYPE]string{
		shapes.START:    "start",
		shapes.STOP:     "stop",
		shapes.IF:       "if",
		shapes.VARIABLE: "variable",
		shapes.ACTION:   "action",
	}

	_, err = file.WriteString("nodes: ")
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
			_, err = file.WriteString(fmt.Sprintf("(%s)%d {\"%s\"},", tag, blockID, line))
			if err != nil {
				return fmt.Errorf("Error saving to file: %v", err)
			}
		}
	}

	edgeMatrix := buildAdjacencyMatrix(blocks, connections)
	adjacencySerialized := serializeAdjacencyMatrix(edgeMatrix)

	_, err = file.WriteString("\nadjacency: \n")
	if err != nil {
		return fmt.Errorf("Error saving to file: %v", err)
	}

	_, err = file.WriteString(adjacencySerialized)
	if err != nil {
		return fmt.Errorf("Error saving to file: %v", err)
	}

	fmt.Println("Graph saved to JSON.")
	return nil
}

// Serialization of matrix to one line
func serializeAdjacencyMatrix(matrix [][]int) string {
	var rows []string
	for _, row := range matrix {
		rows = append(
			rows,
			fmt.Sprintf(
				"[%s]",
				strings.Trim(strings.Join(strings.Fields(fmt.Sprint(row)), ","), "[]"),
			),
		)
	}
	return fmt.Sprintf("[%s]", strings.Join(rows, ",\n"))
}

// Tworzenie macierzy sąsiedztwa na podstawie węzłów i połączeń
func buildAdjacencyMatrix(blocks []shapes.Shape, connections []shapes.Connection) [][]int {
	// Mapa węzłów na indeksy w macierzy
	nodeIndexMap := make(map[int]int)
	for i, block := range blocks {
		nodeIndexMap[block.GetBlockId()] = i
	}

	// Tworzenie pustej macierzy NxN
	N := len(blocks)
	edgeMatrix := make([][]int, N)
	for i := 0; i < N; i++ {
		edgeMatrix[i] = make([]int, N)
	}

	// Dodawanie połączeń do macierzy
	for _, conn := range connections {
		inID := conn.GetInShapeId()
		outID := conn.GetOutShapeId()

		inIndex, okIn := nodeIndexMap[inID]
		outIndex, okOut := nodeIndexMap[outID]

		if okIn && okOut {
			edgeMatrix[inIndex][outIndex] = 1
		} else {
			fmt.Printf("Błąd: Nie znaleziono indeksów dla połączenia %d -> %d\n", inID, outID)
		}
	}

	return edgeMatrix
}
