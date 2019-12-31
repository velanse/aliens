package io

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/velanse/aliens/world"
	"log"
	"os"
	"strings"
)

const(
	separator = "="
)

func ReadFile(filename string) (map[string]*world.Node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	nodes := make(map[string]*world.Node)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineSlice := strings.Fields(scanner.Text())
		if len(lineSlice) == 0 {
			return nil, errors.New("map file should not contain empty lines")
		}
		nodeName := lineSlice[0]

		var node *world.Node
		if n, ok := nodes[nodeName]; ok {
			node = n
		} else {
			node = world.NewNode(nodeName)
		}

		for i:=1; i<len(lineSlice); i++ {
			connection := strings.Split(lineSlice[i], separator)
			if len(connection) < 2 {
				return nil, errors.New("invalid connection mapping")
			}
			dir := connection[0]
			if !world.IsValidDirection(dir) {
				return nil, errors.New("wrong direction")
			}

			if n, ok := nodes[connection[1]]; ok {
				node.Destinations[dir] = n
			} else {
				nodes[connection[1]] = world.NewNode(connection[1])
				node.Destinations[dir] = nodes[connection[1]]
			}

			// since the return way may not be specified in the input file
			node.Destinations[dir].Destinations[world.OppositeDirection[dir]] = node
		}
		nodes[nodeName] = node
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nodes, nil
}

func PrintNodes(nodes map[string]*world.Node) {
	for _, node := range nodes {
		if !node.Active {
			continue
		}
		output := node.Name
		for k, v := range node.Destinations {
			if !v.Active {
				continue
			}
			output += fmt.Sprintf(" %s=%s", k, v.Name)
		}

		fmt.Println(output)
	}
}
