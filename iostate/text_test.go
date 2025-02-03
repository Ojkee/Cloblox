package iostate_test

// import (
// 	"io/ioutil"
// 	"os"
// 	"testing"

// 	"Cloblox/blocks"
// 	"Cloblox/graph"
// 	"Cloblox/iostate"
// 	"Cloblox/shapes"
// )

// func TestSaveToTxt(t *testing.T) {
// 	// Tworzymy blok Start
// 	block1 := &shapes.ShapeDefault{}
// 	block1.SetBlockId(1)
// 	block1.SetName("Start")
// 	content1 := &[]string{"Start"}
// 	block1.SetContent(content1) // Bez wskaźnika
// 	block1.SetShapeType(shapes.START)

// 	// Tworzymy blok Stop
// 	block2 := &shapes.ShapeDefault{}
// 	block2.SetBlockId(2)
// 	block2.SetName("Stop")
// 	content2 := &[]string{"Stop"}
// 	block2.SetContent(content2)
// 	block2.SetShapeType(shapes.STOP)

// 	// Tworzenie bloków Shape
// 	shape1 := &shapes.ShapeDefault{}
// 	shape1.SetBlockId(1)
// 	shape1.SetName("Start")
// 	shape1.SetContent(content1)
// 	shape1.SetShapeType(shapes.START)

// 	shape2 := &shapes.ShapeDefault{}
// 	shape2.SetBlockId(2)
// 	shape2.SetName("Stop")
// 	shape2.SetContent(content2)
// 	shape2.SetShapeType(shapes.STOP)

// 	// Tworzymy obiekt Graph i dodajemy bloki
// 	// Adaptery zamieniające Shape na Block
// 	block12 := iostate.NewShapeBlockAdapter(shape1)
// 	block22 := iostate.NewShapeBlockAdapter(shape2)

// 	// Łączenie bloków
// 	block12.SetNext(block22)

// 	// Tworzenie grafu
// 	blocksList := []blocks.Block{block12, block22}
// 	g := graph.NewGraph(&blocksList)

// 	// Tworzymy połączenie między blokami
// 	connections := []shapes.Connection{
// 		*shapes.NewConnection(0, 0, 1, 1, 1, 2, false, false),
// 	}

// 	// Tworzymy plik tymczasowy
// 	tmpFile, err := ioutil.TempFile("", "test_save_to_txt")
// 	if err != nil {
// 		t.Fatalf("Temp file err: %v\n", err)
// 	}
// 	defer os.Remove(tmpFile.Name()) // Usunięcie pliku po teście

// 	// Wywołujemy funkcję zapisu
// 	err = iostate.SaveToTxt(tmpFile.Name(), g, []shapes.Shape{block1, block2}, connections)
// 	if err != nil {
// 		t.Fatalf("Error saving to file: %v", err)
// 	}

// 	// Odczytujemy zawartość pliku
// 	data, err := ioutil.ReadFile(tmpFile.Name())
// 	if err != nil {
// 		t.Fatalf("Error reading from file: %v", err)
// 	}

// 	// Oczekiwany wynik
// 	expectedContent := `Blocks:
// <h1>"Start",1</h1>
// <h2>"Stop",2</h2>

// Connections:
// <c>s1,2</c>
// `

// 	// Sprawdzamy poprawność zapisu
// 	if string(data) != expectedContent {
// 		t.Errorf(
// 			"Nieprawidłowa zawartość pliku.\nOczekiwano:\n%s\nOtrzymano:\n%s",
// 			expectedContent,
// 			string(data),
// 		)
// 	}
// }

// // func TestSaveToTxtBubbleSort(t *testing.T) {
// // 	// Creating blocks for bubble sort
// // 	block1 := &shapes.ShapeDefault{}
// // 	block1.SetBlockId(1)
// // 	block1.SetName("Start")
// // 	block1.SetContent(&[]string{"Start"})
// // 	block1.SetShapeType(shapes.START)

// // 	block2 := &shapes.ShapeDefault{}
// // 	block2.SetBlockId(2)
// // 	block2.SetName("Action Block")
// // 	block2.SetContent(&[]string{"int i, j;"})
// // 	block2.SetShapeType(shapes.ACTION)

// // 	block3 := &shapes.ShapeDefault{}
// // 	block3.SetBlockId(3)
// // 	block3.SetName("Variable Block")
// // 	block3.SetContent(&[]string{"for i = 0 to n-1"})
// // 	block3.SetShapeType(shapes.VARIABLE)

// // 	block4 := &shapes.ShapeDefault{}
// // 	block4.SetBlockId(4)
// // 	block4.SetName("Variable Block")
// // 	block4.SetContent(&[]string{"for j = 0 to n-i-1"})
// // 	block4.SetShapeType(shapes.VARIABLE)

// // 	block5 := &shapes.ShapeDefault{}
// // 	block5.SetBlockId(5)
// // 	block5.SetName("If Block")
// // 	block5.SetContent(&[]string{"if arr[j] > arr[j+1]"})
// // 	block5.SetShapeType(shapes.IF)

// // 	block6 := &shapes.ShapeDefault{}
// // 	block6.SetBlockId(6)
// // 	block6.SetName("Action Block")
// // 	block6.SetContent(&[]string{"swap(arr[j], arr[j+1])"})
// // 	block6.SetShapeType(shapes.ACTION)

// // 	block7 := &shapes.ShapeDefault{}
// // 	block7.SetBlockId(7)
// // 	block7.SetName("Stop")
// // 	block7.SetContent(&[]string{"Stop"})
// // 	block7.SetShapeType(shapes.STOP)

// // 	blocks := []shapes.Shape{block1, block2, block3, block4, block5, block6, block7}

// // 	// Connections for bubble sort
// // 	connections := []shapes.Connection{
// // 		*shapes.NewConnection(0, 0, 1, 1, 1, 2, false, false),
// // 		*shapes.NewConnection(0, 0, 2, 2, 2, 3, false, false),
// // 		*shapes.NewConnection(0, 0, 3, 3, 3, 4, false, false),
// // 		*shapes.NewConnection(0, 0, 4, 4, 4, 5, false, false),
// // 		*shapes.NewConnection(0, 0, 5, 5, 5, 6, true, false), // Prawa strona - tak
// // 		*shapes.NewConnection(0, 0, 5, 5, 5, 4, false, true), // Lewa srona - nie
// // 		*shapes.NewConnection(0, 0, 6, 6, 6, 4, false, false),
// // 		*shapes.NewConnection(0, 0, 4, 4, 4, 3, false, true),
// // 		*shapes.NewConnection(0, 0, 3, 3, 3, 7, false, true),
// // 	}

// // 	tmpFile, err := ioutil.TempFile("", "test_save_to_txt_bubble_sort")
// // 	if err != nil {
// // 		t.Fatalf("Temp file err: %v\n", err)
// // 	}
// // 	defer os.Remove(tmpFile.Name())
// // 	err = iostate.SaveToTxt(tmpFile.Name(), blocks, connections)
// // 	if err != nil {
// // 		t.Fatalf("Error saving to file\n: %v", err)
// // 	}

// // 	data, err := ioutil.ReadFile(tmpFile.Name())
// // 	if err != nil {
// // 		t.Fatalf("Error reading from file: %v", err)
// // 	}

// // 	// Expected content of the txt
// // 	expectedContent := `Blocks:
// // <h1>"Start",1</h1>
// // <a>"int i, j;",2</a>
// // <v>"for i = 0 to n-1",3</v>
// // <v>"for j = 0 to n-i-1",4</v>
// // <f>"if arr[j] > arr[j+1]",5</f>
// // <a>"swap(arr[j], arr[j+1])",6</a>
// // <h2>"Stop",7</h2>

// // Connections:
// // <c>s1,2</c>
// // <c>s2,3</c>
// // <c>s3,4</c>
// // <c>s4,5</c>
// // <c>r5,6</c>
// // <c>l5,4</c>
// // <c>s6,4</c>
// // <c>s4,3</c>
// // <c>s3,7</c>
// // `
// // 	if string(data) != expectedContent {
// // 		t.Errorf(
// // 			"Expected:\n%s\nGiven:\n%s",
// // 			expectedContent,
// // 			string(data),
// // 		)
// // 	}
// // }
