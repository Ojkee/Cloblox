package iostate

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type NodePDF struct {
	ID    int
	Type  string
	Label string
}

type GraphPDF struct {
	Nodes     []NodePDF
	Adjacency [][]int
}

// SavePDF processes the input file to generate a graph, convert it to TikZ code,
// save the TikZ code to a .tex file, and compile the .tex file to a .pdf file.
//
// Parameters:
// - srcPathJSON: A string specifying the path to the JSON file.
// - dstPathPDF: A string specifying the path to save the resulting PDF file.
//
// Returns:
// - An error if any step in the process fails.
func SavePDF(srcPathJSON, dstPathPDF string) error {
	// Parse the graph from the input JSON file
	graph, err := parseGraph(srcPathJSON)
	if err != nil {
		return fmt.Errorf("Error parsing graph: %v", err)
	}

	// Generate TikZ code
	tikz := generateTikZ(graph)

	// Save the TikZ code to a .tex file
	texFilename := "output.tex"
	if err := saveTikZToFile(tikz, texFilename); err != nil {
		return fmt.Errorf("Error saving TikZ to file: %v", err)
	}

	// Compile the .tex file to a .pdf file
	if err := compileTexToPDF(texFilename, dstPathPDF); err != nil {
		return fmt.Errorf("Error compiling TeX to PDF: %v", err)
	}

	return nil
}

// parseGraph reads a graph definition from a file and returns a GraphPDF struct and an error.
func parseGraph(filePath string) (GraphPDF, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return GraphPDF{}, fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return GraphPDF{}, fmt.Errorf("Error reading the first line for nodes: file is empty")
	}
	nodesLine := scanner.Text()
	nodes := []NodePDF{}
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
					node, err := parseNodeEntry(entry)
					if err != nil {
						return GraphPDF{}, fmt.Errorf("Error parsing node entry: %v", err)
					}
					nodes = append(nodes, node)
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
			node, err := parseNodeEntry(entry)
			if err != nil {
				return GraphPDF{}, fmt.Errorf("Error parsing node entry: %v", err)
			}
			nodes = append(nodes, node)
		}
	}

	if !scanner.Scan() {
		return GraphPDF{}, fmt.Errorf("Error reading the second line for adjacency header: file is incomplete")
	}
	adjacencyHeader := scanner.Text()
	if !strings.HasPrefix(adjacencyHeader, "adjacency:") {
		return GraphPDF{}, fmt.Errorf("Error parsing adjacency header: expected 'adjacency:' on the second line")
	}

	matrixBuilder := strings.Builder{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			matrixBuilder.WriteString(line)
		}
	}
	matrixData := matrixBuilder.String()

	if strings.HasPrefix(matrixData, "[") && strings.HasSuffix(matrixData, "]") {
		matrixData = strings.Trim(matrixData, "[] ")
	} else {
		return GraphPDF{}, fmt.Errorf("Error parsing adjacency matrix: missing outer brackets")
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
				return GraphPDF{}, fmt.Errorf("Error parsing adjacency matrix value at row %d, column %d: %v", i+1, j+1, err)
			}
			rowInt = append(rowInt, value)
		}
		adjacency = append(adjacency, rowInt)
	}

	return GraphPDF{
		Nodes:     nodes,
		Adjacency: adjacency,
	}, nil
}

// parseNodeEntry parses a single node entry from a string and returns a NodePDF and an error.
func parseNodeEntry(entry string) (NodePDF, error) {
	nodeRegex := regexp.MustCompile(`\((.*?)\)(\d+) \{(.*)\}`)
	matches := nodeRegex.FindStringSubmatch(entry)
	if len(matches) != 4 {
		return NodePDF{}, fmt.Errorf("Error parsing node entry: invalid format in '%s'", entry)
	}

	id, err := strconv.Atoi(matches[2])
	if err != nil {
		return NodePDF{}, fmt.Errorf("Error parsing node ID: %v", err)
	}

	return NodePDF{
		ID:    id,
		Type:  matches[1],
		Label: strings.TrimSpace(matches[3]),
	}, nil
}

// generateTikZ generates a TikZ representation of the given graph.
func generateTikZ(graph GraphPDF) string {
	var tikz string

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

	tikz += "\\end{tikzpicture}\n"
	tikz += "\\end{document}\n"

	return tikz
}

// saveTikZToFile saves the given TikZ code to a file with the specified filename and returns an error.
func saveTikZToFile(tikz string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Error saving to file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(tikz)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}
	return nil
}

// compileTexToPDF compiles a TeX file to a PDF using the pdflatex command and returns an error.
func compileTexToPDF(texPath, pdfPath string) error {
	if _, err := exec.LookPath("pdflatex"); err != nil {
		return fmt.Errorf("Error finding pdflatex: %v", err)
	}

	outArg := fmt.Sprintf("--output-directory=%s", pdfPath)
	cmd := exec.Command("pdflatex", outArg, texPath)
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error compiling TeX to PDF: %v", err)
	}

	// Clean up auxiliary files
	auxFile := strings.Replace(texPath, ".tex", ".aux", 1)
	logFile := strings.Replace(texPath, ".tex", ".log", 1)
	_ = os.Remove(auxFile)
	_ = os.Remove(logFile)

	return nil
}
