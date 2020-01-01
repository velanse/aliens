package world

import (
	"math/rand"
	"strconv"

	"github.com/velanse/aliens/printer"
)

const (
	North    = "north"
	South    = "south"
	West     = "west"
	East     = "east"
	maxMoves = 10000
)

var OppositeDirection = map[string]string{
	North: South,
	South: North,
	West:  East,
	East:  West,
}

func IsValidDirection(d string) bool {
	return d == North || d == South || d == East || d == West
}

func InhabitWithAliens(nodes map[string]*Node, nodeNames []string, n int, p printer.Printer) []*Alien {
	aliens := make([]*Alien, 0, n)
	for i := 0; i < n; i++ {
		randomKey := nodeNames[rand.Intn(len(nodeNames))]
		randomNode := nodes[randomKey]
		if !randomNode.Active {
			p.Debug("Alien %d landed to dead city %s \n", i, randomNode.Name)

			// let's consider if an Alien arrives to a dead city - he dies immediately (we are not creating it)
			continue
		}

		a := &Alien{
			Name:  strconv.Itoa(i),
			Node:  randomNode,
			Alive: true,
		}
		p.Debug("Alien %s landed to %s \n", a.Name, randomNode.Name)

		if randomNode.Alien != nil {
			//node is destroyed immediately
			p.Debug("%s has been destroyed by alien %s and alien %s immediately! \n", randomNode.Name, randomNode.Alien.Name, a.Name)

			randomNode.destroyByAliens(randomNode.Alien, a, p)
		}

		a.Node = randomNode
		randomNode.Alien = a

		aliens = append(aliens, a)
	}

	return aliens
}

func RunSimulation(aliens []*Alien, p printer.Printer) {
	for len(aliens) > 0 {
		randomKey := rand.Intn(len(aliens))

		randomAlien := aliens[randomKey]
		terminated := randomAlien.makeMove(p)
		randomAlien.MovesMade++

		if terminated || randomAlien.MovesMade >= maxMoves {
			aliens[randomKey] = aliens[len(aliens)-1]
			aliens = aliens[:len(aliens)-1]
		}
	}
}
