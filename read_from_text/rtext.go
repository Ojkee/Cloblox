package read_from_text

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/sqweek/dialog"

	"Cloblox/graph"
	"Cloblox/window"
)

func ReadFromTxt(blocks []graph.Graph, connections []*window.Connection) error {
	// Open window to choose the file
	filePath, err := dialog.File().Title("Select a file to read").Load()
	if err != nil {
		log.Fatalf("Failed to select file: %v", err)
	}
	fmt.Printf("Selected file: %s\n", filePath)

	// Read file
	fileBuffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filePath, err)
	}
	inputData := string(fileBuffer)

	// Divide data on rows
	rows := strings.Split(inputData, "\n")
	for _, row := range rows {
		fmt.Println("row:", row)
	}

	return nil
}
