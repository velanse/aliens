package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/velanse/aliens/io"
	"github.com/velanse/aliens/printer"
	"github.com/velanse/aliens/world"
)

func main() {
	var (
		filename string
		n        int
		debug    bool
	)

	fmt.Println("--- Aliens invasion started ---")

	// based on Go's "math/rand" implementation in order not to have deterministic values
	rand.Seed(time.Now().UnixNano())

	flag.StringVar(&filename, "map", "", "A path to file that contains a map of cities")
	flag.IntVar(&n, "N", 0, "Number of aliens")
	flag.BoolVar(&debug, "debug", false, "Set to true in order to output debug messages")

	flag.Parse()

	if len(filename) == 0 {
		fmt.Println("Please provide a path to a filename containing a valid map of cities")
	}

	printer := printer.NewStdoutPrinter(debug)

	nodes, err := io.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	nodeNames := world.GetNodeNames(nodes)
	aliens := world.InhabitWithAliens(nodes, nodeNames, n, printer)
	world.RunSimulation(aliens, printer)

	fmt.Println("--- World map after invasion ---")

	io.PrintNodes(nodes, printer)
}
