package iostate_test

// import (
// 	"io/ioutil"
// 	"os"
// 	"testing"

// 	"Cloblox/blocks"
// 	"Cloblox/graph"
// 	iostate "Cloblox/iostate"
// )

// func TestSaveToJson(t *testing.T) {
// 	// Tworzymy blok Start
// 	block1 := &blocks.StartBlock{}
// 	block1.SetId(1)

// 	// Tworzymy blok Stop
// 	block2 := &blocks.StopBlock{}
// 	block2.SetId(2)

// 	// Tworzymy nowy graf
// 	g := graph.NewGraph()

// 	// Dodajemy bloki do grafu
// 	g.AddBlock(block1)
// 	g.AddBlock(block2)

// 	// Tworzymy połączenie między blokiem Start a Stop
// 	g.AddConnection(block1, block2)

// 	// Tworzymy plik tymczasowy
// 	tmpFile, err := ioutil.TempFile("", "test_save_to_json")
// 	if err != nil {
// 		t.Fatalf("Temp file err: %v\n", err)
// 	}
// 	defer os.Remove(tmpFile.Name())

// 	// Zapisujemy graf do pliku
// 	err = iostate.SaveToJson(tmpFile.Name(), g)
// 	if err != nil {
// 		t.Fatalf("Error saving to file\n: %v", err)
// 	}

// 	// Odczytujemy dane z pliku
// 	data, err := ioutil.ReadFile(tmpFile.Name())
// 	if err != nil {
// 		t.Fatalf("Error reading from file: %v", err)
// 	}

// 	// Oczekiwana zawartość
// 	expectedContent := `nodes: (start)1 {"Start"},(stop)2 {"Stop"},
// adjacency:
// [[0,1],
// [0,0]]`

// 	// Sprawdzamy, czy zawartość jest zgodna
// 	if string(data) != expectedContent {
// 		t.Errorf(
// 			"Nieprawidłowa zawartość pliku.\nOczekiwano:\n%s\nOtrzymano:\n%s",
// 			expectedContent,
// 			string(data),
// 		)
// 	}
// }

// func TestSaveToJsonBubble1(t *testing.T) {
// 	// Tworzymy bloki dla algorytmu Bubble Sort
// 	block1 := &blocks.StartBlock{}
// 	block1.SetId(1)

// 	block2 := &blocks.ActionBlock{}
// 	block2.SetId(2)
// 	block2.SetContent("int i, j;")

// 	block3 := &blocks.VariablesBlock{}
// 	block3.SetId(3)
// 	block3.SetContent("for i = 0 to n-1")

// 	block4 := &blocks.VariablesBlock{}
// 	block4.SetId(4)
// 	block4.SetContent("for j = 0 to n-i-1")

// 	block5 := &blocks.IfBlock{}
// 	block5.SetId(5)
// 	block5.SetConditionExpr("arr[j] > arr[j+1]")

// 	block6 := &blocks.ActionBlock{}
// 	block6.SetId(6)
// 	block6.SetContent("swap(arr[j], arr[j+1])")

// 	block7 := &blocks.StopBlock{}
// 	block7.SetId(7)

// 	// Tworzymy graf
// 	g := graph.NewGraph()

// 	// Dodajemy bloki do grafu
// 	g.AddBlock(block1)
// 	g.AddBlock(block2)
// 	g.AddBlock(block3)
// 	g.AddBlock(block4)
// 	g.AddBlock(block5)
// 	g.AddBlock(block6)
// 	g.AddBlock(block7)

// 	// Tworzymy połączenia między blokami
// 	g.AddConnection(block1, block2)
// 	g.AddConnection(block2, block3)
// 	g.AddConnection(block3, block4)
// 	g.AddConnection(block4, block5)
// 	g.AddConnection(block5, block6) // true branch
// 	g.AddConnection(block5, block4) // false branch
// 	g.AddConnection(block6, block4)
// 	g.AddConnection(block4, block3)
// 	g.AddConnection(block3, block7)

// 	// Tworzymy plik tymczasowy
// 	tmpFile, err := ioutil.TempFile("", "test_save_to_json_bubble")
// 	if err != nil {
// 		t.Fatalf("Temp file err: %v\n", err)
// 	}
// 	defer os.Remove(tmpFile.Name())

// 	// Zapisujemy graf do pliku
// 	err = iostate.SaveToJson(tmpFile.Name(), g)
// 	if err != nil {
// 		t.Fatalf("Error saving to file\n: %v", err)
// 	}

// 	// Odczytujemy dane z pliku
// 	data, err := ioutil.ReadFile(tmpFile.Name())
// 	if err != nil {
// 		t.Fatalf("Error reading from file: %v", err)
// 	}

// 	// Oczekiwana zawartość
// 	expectedContent := `nodes: (start)1 {"Start"},(action)2 {"int i, j;"},(variable)3 {"for i = 0 to n-1"},(variable)4 {"for j = 0 to n-i-1"},(if)5 {"if arr[j] > arr[j+1]"},(action)6 {"swap(arr[j], arr[j+1])"},(stop)7 {"Stop"},
// adjacency:
// [[0,1,0,0,0,0,0],
// [0,0,1,0,0,0,0],
// [0,0,0,1,0,0,1],
// [0,0,1,0,1,0,0],
// [0,0,0,3,0,2,0],
// [0,0,0,1,0,0,0],
// [0,0,0,0,0,0,0]]`

// 	// Sprawdzamy, czy zawartość jest zgodna
// 	if string(data) != expectedContent {
// 		t.Errorf(
// 			"Nieprawidłowa zawartość pliku.\nOczekiwano:\n%s\nOtrzymano:\n%s",
// 			expectedContent,
// 			string(data),
// 		)
// 	}
// }
