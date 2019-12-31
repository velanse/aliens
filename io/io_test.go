package io

import (
	"errors"
	"github.com/velanse/aliens/printer"
	"reflect"
	"testing"

	"github.com/velanse/aliens/world"
)

func TestReadFile(t *testing.T) {
	filename := "../testdata/test01.txt"

	expectedNodesNumber := 5

	nodes, err := ReadFile(filename)

	if err != nil {
		t.Errorf("There should be no error reading existing file %s", filename)
	}

	if len(nodes) != expectedNodesNumber {
		t.Errorf("Wrong nodes amount obtained: Expected: %v, got: %v", expectedNodesNumber, len(nodes))
	}
}

func TestReadFileWithData(t *testing.T) {
	filename := "../testdata/test02.txt"
	expected := setupWorld()

	nodes, err := ReadFile(filename)

	if err != nil {
		t.Errorf("There should be no error reading existing file %s", filename)
	}

	if !reflect.DeepEqual(nodes, expected) {
		t.Errorf("Wrong data obtained from file. Expected: %+v, got: %+v", expected, nodes)
	}
}

func TestReadFileError(t *testing.T) {
	filename := "noexistingfile.txt"
	nodes, err := ReadFile(filename)

	if err == nil {
		t.Errorf("An error expected")
	}

	if len(nodes) != 0 {
		t.Errorf("Nodes list should be empty")
	}
}

func TestReadFileWrongData(t *testing.T) {
	filename := "../testdata/test03.txt"
	nodes, err := ReadFile(filename)

	expectedError := errors.New("wrong direction")

	if err == nil {
		t.Errorf("An error expected")
	}

	if !reflect.DeepEqual(err, expectedError) {
		t.Errorf("An error expected: %v, got %v", expectedError, err)
	}

	if len(nodes) != 0 {
		t.Errorf("Nodes list should be empty")
	}
}

func TestPrintNodes(t *testing.T) {
	nodes := setupWorld()

	expectedCitiesPrinted := 3

	printer := &printer.MockPrinter{}

	PrintNodes(nodes, printer)

	if len(printer.Messages) != 3 {
		t.Errorf("Wrong cities amount printed. Expected: %d, got: %+v", expectedCitiesPrinted, printer.Messages)
	}
}

func setupWorld() map[string]*world.Node {
	A := &world.Node{
		Name:         "A",
		Active:       true,
		Destinations: make(map[string]*world.Node),
	}

	B := &world.Node{
		Name:         "B",
		Active:       true,
		Destinations: make(map[string]*world.Node),
	}

	C := &world.Node{
		Name:         "C",
		Active:       true,
		Destinations: make(map[string]*world.Node),
	}

	A.Destinations["east"] = B
	B.Destinations["west"] = A
	B.Destinations["east"] = C
	C.Destinations["west"] = B

	return map[string]*world.Node{
		"A": A,
		"B": B,
		"C": C,
	}
}
