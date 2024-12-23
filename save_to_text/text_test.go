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
	if string(data) != expectedContent {
		t.Errorf(
			"Expected:\n%s\nGiven:\n%s",
			expectedContent,
			string(data),
		)
	}
}
