package save_to_text_test

import (
	"io/ioutil"
	"os"
	"testing"

	"Cloblox/save_to_text"
	"Cloblox/shapes"
	"Cloblox/window"
)

func TestSaveToTxt(t *testing.T) {
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

	blocks := []shapes.Shape{block1, block2}

	connections := []*window.Connection{
		window.NewConnection(0, 0, 1, 1, 1, 2, false, false),
	}

	tmpFile, err := ioutil.TempFile("", "test_save_to_txt")
	if err != nil {
		t.Fatalf("Temp file err: %v\n", err)
	}
	defer os.Remove(tmpFile.Name())
	err = save_to_text.SaveToTxt(tmpFile.Name(), blocks, connections)
	if err != nil {
		t.Fatalf("Error saving to file\n: %v", err)
	}

	data, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Error reading from file: %v", err)
	}

	expectedContent := `Blocks:
<h1>"Start",1</h1>
<h2>"Stop",2</h2>

Connections:
<c>1,2</c>
`

	// Checking if the expected content is the same as the content
	if string(data) != expectedContent {
		t.Errorf("Nieprawidłowa zawartość pliku.\nOczekiwano:\n%s\nOtrzymano:\n%s", expectedContent, string(data))
	}
}

func TestSaveToTxtBubbleSort(t *testing.T) {
	// Creating blocks for bubble sort
	block1 := &shapes.ShapeDefault{}
	block1.SetBlockId(1)
	block1.SetName("Start")
	block1.SetContent(&[]string{"Start"})
	block1.SetShapeType(shapes.START)

	block2 := &shapes.ShapeDefault{}
	block2.SetBlockId(2)
	block2.SetName("Initialize")
	block2.SetContent(&[]string{"int i, j;"})
	block2.SetShapeType(shapes.ACTION)

	block3 := &shapes.ShapeDefault{}
	block3.SetBlockId(3)
	block3.SetName("Outer Loop")
	block3.SetContent(&[]string{"for i = 0 to n-1"})
	block3.SetShapeType(shapes.VARIABLE)

	block4 := &shapes.ShapeDefault{}
	block4.SetBlockId(4)
	block4.SetName("Inner Loop")
	block4.SetContent(&[]string{"for j = 0 to n-i-1"})
	block4.SetShapeType(shapes.VARIABLE)

	block5 := &shapes.ShapeDefault{}
	block5.SetBlockId(5)
	block5.SetName("Comparison")
	block5.SetContent(&[]string{"if arr[j] > arr[j+1]"})
	block5.SetShapeType(shapes.IF)

	block6 := &shapes.ShapeDefault{}
	block6.SetBlockId(6)
	block6.SetName("Swap")
	block6.SetContent(&[]string{"swap(arr[j], arr[j+1])"})
	block6.SetShapeType(shapes.ACTION)

	block7 := &shapes.ShapeDefault{}
	block7.SetBlockId(7)
	block7.SetName("Stop")
	block7.SetContent(&[]string{"Stop"})
	block7.SetShapeType(shapes.STOP)

	blocks := []shapes.Shape{block1, block2, block3, block4, block5, block6, block7}

	// Connections for bubble sort
	connections := []*window.Connection{
		window.NewConnection(0, 0, 1, 1, 1, 2, false, false),
		window.NewConnection(0, 0, 2, 2, 2, 3, false, false),
		window.NewConnection(0, 0, 3, 3, 3, 4, false, false),
		window.NewConnection(0, 0, 4, 4, 4, 5, false, false),
		window.NewConnection(0, 0, 5, 5, 5, 6, true, false), // True branch
		window.NewConnection(0, 0, 5, 5, 5, 4, false, true), // False branch
		window.NewConnection(0, 0, 6, 6, 6, 4, false, false),
		window.NewConnection(0, 0, 4, 4, 4, 3, false, true),
		window.NewConnection(0, 0, 3, 3, 3, 7, false, true),
	}

	tmpFile, err := ioutil.TempFile("", "test_save_to_txt_bubble_sort")
	if err != nil {
		t.Fatalf("Temp file err: %v\n", err)
	}
	defer os.Remove(tmpFile.Name())
	err = save_to_text.SaveToTxt(tmpFile.Name(), blocks, connections)
	if err != nil {
		t.Fatalf("Error saving to file\n: %v", err)
	}

	data, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Error reading from file: %v", err)
	}

	// Expected content of the txt
	expectedContent := `Blocks:
<h1>"Start",1</h1>
<a>"int i, j;",2</a>
<v>"for i = 0 to n-1",3</v>
<v>"for j = 0 to n-i-1",4</v>
<f>"if arr[j] > arr[j+1]",5</f>
<a>"swap(arr[j], arr[j+1])",6</a>
<h2>"Stop",7</h2>

Connections:
<c>1,2</c>
<c>2,3</c>
<c>3,4</c>
<c>4,5</c>
<c>5,6</c>
<c>5,4</c>
<c>6,4</c>
<c>4,3</c>
<c>3,7</c>
`
	if string(data) != expectedContent {
		t.Errorf(
			"Expected:\n%s\nGiven:\n%s",
			expectedContent,
			string(data),
		)
	}
}
