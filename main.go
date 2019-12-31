package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/velanse/aliens/printer"

	"github.com/velanse/aliens/io"
	"github.com/velanse/aliens/world"
)

func main() {
	var (
		filename string
		n        int
		async    bool
		debug    bool
		delay    uint
	)

	fmt.Println("--- Aliens invasion started ---")

	flag.StringVar(&filename, "map", "", "A path to file that contains a map of cities")
	flag.IntVar(&n, "N", 0, "Number of aliens")
	flag.BoolVar(&async, "async", true, "Set to true to make aliens move asynchronously")
	flag.BoolVar(&debug, "debug", false, "Set to true in order to output debug messages")
	flag.UintVar(&delay, "delay", 0, "For real life invasion simulation you can specify maximum time in Milliseconds that alien stays in the city before next move. Only for async mode")

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

	if async {
		world.RunAsync(aliens, delay, printer)
	} else {
		world.RunSync(aliens, printer)
	}

	fmt.Println("--- World map after invasion ---")

	io.PrintNodes(nodes, printer)
}
