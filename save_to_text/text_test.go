package functools

import (
	"Cloblox/shapes"
	"Cloblox/window"
	"io/ioutil"
	"os"
	"testing"
)

func TestSaveToTxt(t *testing.T) {
	// Przykładowe dane testowe
	block1 := &shapes.ShapeDefault{}
	block1.SetBlockId(1)
	block1.SetName("Start")
	block1.SetContent(&[]string{"Start"})
	block1.GetShapeType(shapes.START)

	block2 := &shapes.ShapeDefault{}
	block2.SetBlockId(2)
	block2.SetName("Stop")
	block2.SetContent(&[]string{"Stop"})
	block2.GetShapeType(shapes.STOP)

	blocks := []shapes.Shape{block1, block2}

	connections := []*window.Connection{
		window.NewConnection(0, 0, 1, 1, 1, 2, false, false),
	}

	// Stwórz plik tymczasowy
	tmpFile, err := ioutil.TempFile("", "test_save_to_txt")
	if err != nil {
		t.Fatalf("Nie udało się utworzyć pliku tymczasowego: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Usuń plik po zakończeniu testu

	// Wykonaj zapis
	err = SaveToTxt(tmpFile.Name(), blocks, connections)
	if err != nil {
		t.Fatalf("Błąd podczas zapisu do pliku: %v", err)
	}

	// Odczytaj zawartość pliku
	data, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Błąd podczas odczytu pliku: %v", err)
	}

	// Oczekiwana zawartość pliku
	expectedContent := `Bloki:
<h1>"Start",1</h1>
<h2>"Stop",2</h2>

Połączenia:
<c>1,2</c>
`

	// Porównaj zawartość pliku z oczekiwaną zawartością
	if string(data) != expectedContent {
		t.Errorf("Nieprawidłowa zawartość pliku.\nOczekiwano:\n%s\nOtrzymano:\n%s", expectedContent, string(data))
	}
}
