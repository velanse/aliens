package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/velanse/aliens/io"
	"github.com/velanse/aliens/world"
)

func main() {
	var (
		filename string
		n int
	)

	fmt.Println("--- Aliens invasion started ---")

	flag.StringVar(&filename, "map", "", "A path to file that contains a map of cities")
	flag.IntVar(&n, "N", 0, "Number of aliens")

	flag.Parse()

	if len(filename) == 0 {
		fmt.Println("Please provide a path to a filename containing a map of cities")
	}

	nodes, err := io.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	nodeNames := world.GetNodeNames(nodes)

	aliens := world.InhabitWithAliens(nodes, nodeNames, n)

	var wg sync.WaitGroup
	wg.Add(len(aliens))

	for _, a := range aliens {
		go func() {
			defer wg.Done()
			a.Dispatch()
		}()
	}

	wg.Wait()

	io.PrintNodes(nodes)
}