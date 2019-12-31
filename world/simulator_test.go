package world

import (
	"reflect"
	"sort"
	"testing"

	"github.com/velanse/aliens/printer"
)

type scenario struct {
	name                    string
	nodes                   map[string]*Node
	n                       int
	expectedCitiesDestroyed int
	expectedCitiesLeft      int
	async                   bool
}

var scenarios = []scenario{
	{
		name:                    "2 Aliens sync mode",
		nodes:                   setupWorld(3),
		n:                       2,
		expectedCitiesDestroyed: 1,
		expectedCitiesLeft:      2,
		async:                   false,
	},
	{
		name:                    "1 alien struggle sync mode",
		nodes:                   setupWorld(3),
		n:                       1,
		expectedCitiesDestroyed: 0,
		expectedCitiesLeft:      3,
		async:                   false,
	},
	{
		name:                    "3 aliens invasion sync mode",
		nodes:                   setupWorld(3),
		n:                       3,
		expectedCitiesDestroyed: 1,
		expectedCitiesLeft:      2,
		async:                   false,
	},
	{
		name:                    "2 aliens async in 2 cities",
		nodes:                   setupWorld(2),
		n:                       2,
		expectedCitiesDestroyed: 1,
		expectedCitiesLeft:      1,
		async:                   true,
	},
}

func TestSync(t *testing.T) {
	for _, s := range scenarios {
		t.Logf("Testing invasion: %v", s.name)

		printer := &printer.MockPrinter{}
		nodeNames := GetNodeNames(s.nodes)

		aliens := InhabitWithAliens(s.nodes, nodeNames, s.n, printer)

		if s.async {
			RunAsync(aliens, 0, printer)
		} else {
			RunSync(aliens, printer)
		}

		if len(printer.Messages) != s.expectedCitiesDestroyed {
			t.Errorf("Wrong cities amount destroyed. Expected: %d, got: %d", s.expectedCitiesDestroyed, len(printer.Messages))
		}

		activeNodesNumber := 0
		for _, n := range s.nodes {
			if n.Active {
				activeNodesNumber++
			}
		}

		if activeNodesNumber != s.expectedCitiesLeft {
			t.Errorf("Wrong cities amount left after invasion. Expected: %d, got: %d", s.expectedCitiesLeft, len(s.nodes))
		}
	}
}

func TestGetNodeNames(t *testing.T) {
	nodes := setupWorld(3)

	expected := []string{"A", "B", "C"}

	nodeNames := GetNodeNames(nodes)
	sort.Strings(nodeNames)

	if !reflect.DeepEqual(nodeNames, expected) {
		t.Errorf("Got wrong node names. Expected: %v, got: %v", expected, nodeNames)
	}
}

func setupWorld(citiesNumber int) map[string]*Node {
	A := &Node{
		Name:         "A",
		Active:       true,
		Destinations: make(map[string]*Node),
	}

	B := &Node{
		Name:         "B",
		Active:       true,
		Destinations: make(map[string]*Node),
	}

	C := &Node{
		Name:         "C",
		Active:       true,
		Destinations: make(map[string]*Node),
	}

	switch citiesNumber {
	case 3:
		A.Destinations["east"] = B
		B.Destinations["west"] = A
		B.Destinations["east"] = C
		C.Destinations["west"] = B

		return map[string]*Node{
			"A": A,
			"B": B,
			"C": C,
		}
	case 2:
		A.Destinations["east"] = B
		B.Destinations["west"] = A

		return map[string]*Node{
			"A": A,
			"B": B,
		}
	}

	return map[string]*Node{}
}
