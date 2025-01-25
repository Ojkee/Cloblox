package iostate

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	ID    int
	Type  string
	Label string
}

type Graph struct {
	Nodes     []Node
	Adjacency [][]int
}

// parseGraph reads a graph definition from a file and returns a Graph struct.
// The file is expected to have a specific format:
// - The first line contains node definitions in the format: (type)ID {label}, ...
// - The second line contains the header "adjacency:"
// - The remaining lines contain the adjacency matrix in the format: [row1],[row2],...
func parseGraph(filePath string) Graph {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the first line for node definitions
	if !scanner.Scan() {
		log.Fatalf("Failed to read the first line for nodes")
	}
	nodesLine := scanner.Text()
	nodes := []Node{}
	var currentNode strings.Builder
	openBraces := 0
	for _, char := range nodesLine {
		switch char {
		case '{':
			openBraces++
			currentNode.WriteRune(char)
		case '}':
			openBraces--
			currentNode.WriteRune(char)
		case ',':
			if openBraces == 0 {
				entry := strings.TrimSpace(currentNode.String())
				if entry != "" {
					parseNodeEntry(entry, &nodes)
				}
				currentNode.Reset()
			} else {
				currentNode.WriteRune(char)
			}
		default:
			currentNode.WriteRune(char)
		}
	}
	if currentNode.Len() > 0 {
		entry := strings.TrimSpace(currentNode.String())
		if entry != "" {
			parseNodeEntry(entry, &nodes)
		}
	}

	// Read the second line for the adjacency matrix header
	if !scanner.Scan() {
		log.Fatalf("Failed to read the second line for adjacency header")
	}
	adjacencyHeader := scanner.Text()
	if !strings.HasPrefix(adjacencyHeader, "adjacency:") {
		log.Fatalf("Invalid input: expected 'adjacency:' on the second line")
	}

	// Read the remaining lines for the adjacency matrix
	matrixBuilder := strings.Builder{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			matrixBuilder.WriteString(line)
		}
	}
	matrixData := matrixBuilder.String()

	// Validate and parse the adjacency matrix
	if strings.HasPrefix(matrixData, "[") && strings.HasSuffix(matrixData, "]") {
		matrixData = strings.Trim(matrixData, "[] ")
	} else {
		log.Fatalf("Invalid adjacency matrix format: missing outer brackets")
	}

	rows := strings.Split(matrixData, "],[")
	adjacency := [][]int{}
	for i, row := range rows {
		row = strings.Trim(row, "[] ")
		numbers := strings.Split(row, ",")
		rowInt := []int{}
		for j, num := range numbers {
			value, err := strconv.Atoi(strings.TrimSpace(num))
			if err != nil {
				log.Fatalf("Invalid adjacency matrix value at row %d, column %d: %s", i+1, j+1, num)
			}
			rowInt = append(rowInt, value)
		}
		adjacency = append(adjacency, rowInt)
	}

	return Graph{
		Nodes:     nodes,
		Adjacency: adjacency,
	}
}

// parseNodeEntry parses a single node entry from a string and appends it to the nodes slice.
// The entry string is expected to be in the format: (type)ID {label}
// If the entry does not match the expected format, the function logs a fatal error.
//
// Parameters:
// - entry: A string representing a node entry in the format (type)ID {label}.
// - nodes: A pointer to a slice of Node structs where the parsed node will be appended.
func parseNodeEntry(entry string, nodes *[]Node) {
	// Compile a regular expression to match the node entry format
	nodeRegex := regexp.MustCompile(`\((.*?)\)(\d+) \{(.*)\}`)

	// Find submatches in the entry string
	matches := nodeRegex.FindStringSubmatch(entry)

	// Check if the entry matches the expected format
	if len(matches) != 4 {
		log.Fatalf("Invalid node entry: %s", entry)
	}

	// Convert the node ID from string to integer
	id, err := strconv.Atoi(matches[2])
	if err != nil {
		log.Fatalf("Invalid node ID: %s", matches[2])
	}

	// Append the parsed node to the nodes slice
	*nodes = append(*nodes, Node{
		ID:    id,
		Type:  matches[1],
		Label: strings.TrimSpace(matches[3]),
	})
}

// generateTikZ generates a TikZ representation of the given graph.
// It constructs a LaTeX document with TikZ code to visualize the graph.
//
// Parameters:
// - graph: A Graph struct containing nodes and adjacency matrix.
//
// Returns:
// - A string containing the LaTeX document with TikZ code.
func generateTikZ(graph Graph) string {
	var tikz string

	// Add LaTeX document preamble and TikZ styles
	tikz += `\documentclass{article}
\usepackage{tikz}
\usetikzlibrary{shapes.geometric, arrows}
\begin{document}
\begin{tikzpicture}[node distance=1.5cm and 1.5cm]
\tikzstyle{startstop} = [ellipse, minimum width=3cm, minimum height=1cm, text centered, draw=black, fill=purple!30]
\tikzstyle{process} = [rectangle, minimum width=3cm, minimum height=1cm, text centered, draw=black, fill=purple!20]
\tikzstyle{decision} = [diamond, minimum width=3cm, minimum height=1cm, text centered, draw=black, fill=purple!50]
\tikzstyle{data} = [trapezium, trapezium left angle=75, trapezium right angle=105, minimum width=3cm, minimum height=1cm, text centered, draw=black, fill=purple!40]
\tikzstyle{arrow} = [thick,->,>=stealth]
`

	verticalPos := 0.0
	horizontalOffset := 0.0

	// Add nodes to the TikZ picture
	for _, node := range graph.Nodes {
		style := "process"
		switch node.Type {
		case "start", "stop":
			style = "startstop"
		case "if":
			style = "decision"
			verticalPos -= 2
		}
		tikz += fmt.Sprintf("\\node[%s] (%d) at (%.2f, %.2f) {%s};\n", style, node.ID, horizontalOffset, verticalPos, node.Label)
		if node.Type == "if" {
			verticalPos -= 4
		} else {
			verticalPos -= 2
		}
	}

	lastNode := len(graph.Nodes)

	// Add edges to the TikZ picture based on the adjacency matrix
	for i := 0; i < lastNode; i++ {
		for j, count := 0, 0; j < lastNode; j++ {
			if graph.Adjacency[i][j] == 1 {
				bendAmount := 20
				if j != i+1 {
					bendAmount = 50
				}

				if graph.Nodes[i].Type == "if" {
					if count == 0 {
						tikz += fmt.Sprintf("\\draw[arrow, bend left=%d] (%d.west) to (%d.north);\n", bendAmount, i+1, j+1)
					} else {
						tikz += fmt.Sprintf("\\draw[arrow, bend left=%d] (%d.east) to (%d.north);\n", bendAmount, i+1, j+1)
					}
					count++
				} else if graph.Nodes[j].Type == "stop" {
					tikz += fmt.Sprintf("\\draw[arrow, bend left=%d] (%d.east) to (%d.north);\n", bendAmount, i+1, j+1)
				} else {
					tikz += fmt.Sprintf("\\draw[arrow, bend left=%d] (%d.south) to (%d.north);\n", bendAmount, i+1, j+1)
				}
			}
		}
	}

	// End TikZ picture and LaTeX document
	tikz += "\\end{tikzpicture}\n"
	tikz += "\\end{document}\n"

	return tikz
}

// saveTikZToFile saves the given TikZ code to a file with the specified filename.
// If the file cannot be created or written to, the function logs a fatal error.
//
// Parameters:
// - tikz: A string containing the TikZ code to be saved.
// - filename: A string specifying the name of the file to save the TikZ code to.
func saveTikZToFile(tikz string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(tikz)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	log.Printf("TikZ code saved to %s", filename)
}

// compileTexToPDF compiles a TeX file to a PDF using the pdflatex command.
// It checks if pdflatex is available in the system's PATH, runs the command,
// and cleans up auxiliary files generated during the compilation.
//
// Parameters:
// - texFile: A string specifying the path to the TeX file to be compiled.
func compileTexToPDF(texFile string) {
	if _, err := exec.LookPath("pdflatex"); err != nil {
		log.Fatalf("pdflatex not found: %v", err)
	}

	cmd := exec.Command("pdflatex", texFile)
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to compile TeX to PDF: %v", err)
	}

	// Clean up auxiliary files
	auxFile := strings.Replace(texFile, ".tex", ".aux", 1)
	logFile := strings.Replace(texFile, ".tex", ".log", 1)
	os.Remove(auxFile)
	os.Remove(logFile)

	log.Printf("PDF successfully generated from %s", texFile)
}

// processInput processes the input file to generate a graph, convert it to TikZ code,
// save the TikZ code to a .tex file, and compile the .tex file to a .pdf file.
//
// Main function of this section, it processes the input file to generate a graph, convert it to TikZ code, save the TikZ code to a .tex file, and compile the .tex file to a .pdf file.
//
// Use: processInput("input.json")
//
// Parameters:
// - filePath: A string specifying the path to the input file.
func processInput(filePath string) {
	// Parse the graph from the input JSON file
	graph := parseGraph(filePath)

	// Generate TikZ code
	tikz := generateTikZ(graph)

	// Save the TikZ code to a .tex file
	texFilename := "output.tex"
	saveTikZToFile(tikz, texFilename)

	// Compile the .tex file to a .pdf file
	compileTexToPDF(texFilename)
}
