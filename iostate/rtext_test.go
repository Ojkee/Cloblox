package iostate_test

import (
	"os"
	"testing"

	iostate "Cloblox/iostate"
	"Cloblox/shapes"
)

func TestReadFromTxt(t *testing.T) {
	// Expected blocks and connections
	block1 := &shapes.ShapeDefault{}
	block1.SetBlockId(1)
	block1.SetName("Start")
	block1.SetContent(&[]string{"Start"})
	block1.SetShapeType(shapes.START)

	block2 := &shapes.ShapeDefault{}
	block2.SetBlockId(2)
	block2.SetName("Stop")
	block2.SetContent(&[]string{"Stop"})
	block2.SetShapeType(shapes.STOP)

	expectedBlocks := []shapes.Shape{block1, block2}

	expectedConnections := []shapes.Connection{
		*shapes.NewConnection(0, 0, 1, 1, 1, 2, false, false),
	}

	// Temporary file content for testing
	fileContent := `Blocks:
<h1>"Start",1</h1>
<h2>"Stop",2</h2>

Connections:
<c>s1,2</c>
`

	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "test_read_from_txt")
	if err != nil {
		t.Fatalf("Temp file creation failed: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write the content to the temporary file
	if _, err := tmpFile.WriteString(fileContent); err != nil {
		t.Fatalf("Error writing to temp file: %v", err)
	}
	tmpFile.Close()

	// Call the ReadFromTxt function
	parsedBlocks, parsedConnections, err := iostate.ReadFromTxt(tmpFile.Name())
	if err != nil {
		t.Fatalf("Error reading from file: %v", err)
	}

	// Validate blocks
	if len(parsedBlocks) != len(expectedBlocks) {
		t.Fatalf("Block count mismatch: expected %d, got %d", len(expectedBlocks), len(parsedBlocks))
	}

	for i, block := range parsedBlocks {
		if block.GetId() != expectedBlocks[i].GetBlockId() || block.GetName() != expectedBlocks[i].GetName() {
			t.Errorf("Block mismatch at index %d: expected ID=%d, Name=%s; got ID=%d, Name=%s",
				i, expectedBlocks[i].GetBlockId(), expectedBlocks[i].GetName(), block.GetId(), block.GetName())
		}
	}

	// Validate connections
	if len(parsedConnections) != len(expectedConnections) {
		t.Fatalf("Connection count mismatch: expected %d, got %d", len(expectedConnections), len(parsedConnections))
	}

	for i, conn := range parsedConnections {
		expectedConn := expectedConnections[i]
		if conn.GetInShapeId() != expectedConn.GetInShapeId() || conn.GetOutShapeId() != expectedConn.GetOutShapeId() {
			t.Errorf("Connection mismatch at index %d: expected (%d -> %d), got (%d -> %d)",
				i, expectedConn.GetInShapeId(), expectedConn.GetOutShapeId(), conn.GetInShapeId(), conn.GetOutShapeId())
		}
	}
}

func TestReadFromTxtBubbleSort(t *testing.T) {
	// Expected blocks and connections
	block1 := &shapes.ShapeDefault{}
	block1.SetBlockId(1)
	block1.SetName("Start")
	block1.SetContent(&[]string{"Start"})
	block1.SetShapeType(shapes.START)

	block2 := &shapes.ShapeDefault{}
	block2.SetBlockId(2)
	block2.SetName("Action Block")
	block2.SetContent(&[]string{"int i, j;"})
	block2.SetShapeType(shapes.ACTION)

	block3 := &shapes.ShapeDefault{}
	block3.SetBlockId(3)
	block3.SetName("Variable Block")
	block3.SetContent(&[]string{"for i = 0 to n-1"})
	block3.SetShapeType(shapes.VARIABLE)

	block4 := &shapes.ShapeDefault{}
	block4.SetBlockId(4)
	block4.SetName("Variable Block")
	block4.SetContent(&[]string{"for j = 0 to n-i-1"})
	block4.SetShapeType(shapes.VARIABLE)

	block5 := &shapes.ShapeDefault{}
	block5.SetBlockId(5)
	block5.SetName("If Block")
	block5.SetContent(&[]string{"if arr[j] > arr[j+1]"})
	block5.SetShapeType(shapes.IF)

	block6 := &shapes.ShapeDefault{}
	block6.SetBlockId(6)
	block6.SetName("Action Block")
	block6.SetContent(&[]string{"swap(arr[j], arr[j+1])"})
	block6.SetShapeType(shapes.ACTION)

	block7 := &shapes.ShapeDefault{}
	block7.SetBlockId(7)
	block7.SetName("Stop")
	block7.SetContent(&[]string{"Stop"})
	block7.SetShapeType(shapes.STOP)

	expectedBlocks := []shapes.Shape{block1, block2, block3, block4, block5, block6, block7}

	// Connections for bubble sort
	expectedConnections := []shapes.Connection{
		*shapes.NewConnection(0, 0, 1, 1, 1, 2, false, false),
		*shapes.NewConnection(0, 0, 2, 2, 2, 3, false, false),
		*shapes.NewConnection(0, 0, 3, 3, 3, 4, false, false),
		*shapes.NewConnection(0, 0, 4, 4, 4, 5, false, false),
		*shapes.NewConnection(0, 0, 5, 5, 5, 6, true, false), // True branch
		*shapes.NewConnection(0, 0, 5, 5, 5, 4, false, true), // False branch
		*shapes.NewConnection(0, 0, 6, 6, 6, 4, false, false),
		*shapes.NewConnection(0, 0, 4, 4, 4, 3, false, true),
		*shapes.NewConnection(0, 0, 3, 3, 3, 7, false, true),
	}

	// Temporary file content for testing
	fileContent := `Blocks:
<h1>"Start",1</h1>
<a>"int i, j;",2</a>
<v>"for i = 0 to n-1",3</v>
<v>"for j = 0 to n-i-1",4</v>
<f>"if arr[j] > arr[j+1]",5</f>
<a>"swap(arr[j], arr[j+1])",6</a>
<h2>"Stop",7</h2>

Connections:
<c>s1,2</c>
<c>s2,3</c>
<c>s3,4</c>
<c>s4,5</c>
<c>r5,6</c>
<c>l5,4</c>
<c>s6,4</c>
<c>s4,3</c>
<c>s3,7</c>
`

	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "test_read_from_txt")
	if err != nil {
		t.Fatalf("Temp file creation failed: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write the content to the temporary file
	if _, err := tmpFile.WriteString(fileContent); err != nil {
		t.Fatalf("Error writing to temp file: %v", err)
	}
	tmpFile.Close()

	// Call the ReadFromTxt function
	parsedBlocks, parsedConnections, err := iostate.ReadFromTxt(tmpFile.Name())
	if err != nil {
		t.Fatalf("Error reading from file: %v", err)
	}

	// Validate blocks
	if len(parsedBlocks) != len(expectedBlocks) {
		t.Fatalf("Block count mismatch: expected %d, got %d", len(expectedBlocks), len(parsedBlocks))
	}

	for i, block := range parsedBlocks {
		if block.GetId() != expectedBlocks[i].GetBlockId() || block.GetName() != expectedBlocks[i].GetName() {
			t.Errorf("Block mismatch at index %d: expected ID=%d, Name=%s; got ID=%d, Name=%s",
				i, expectedBlocks[i].GetBlockId(), expectedBlocks[i].GetName(), block.GetId(), block.GetName())
		}
	}

	// Validate connections
	if len(parsedConnections) != len(expectedConnections) {
		t.Fatalf("Connection count mismatch: expected %d, got %d", len(expectedConnections), len(parsedConnections))
	}

	for i, conn := range parsedConnections {
		expectedConn := expectedConnections[i]
		if conn.GetInShapeId() != expectedConn.GetInShapeId() || conn.GetOutShapeId() != expectedConn.GetOutShapeId() {
			t.Errorf("Connection mismatch at index %d: expected (%d -> %d), got (%d -> %d)",
				i, expectedConn.GetInShapeId(), expectedConn.GetOutShapeId(), conn.GetInShapeId(), conn.GetOutShapeId())
		}
	}
}
