package world

import (
	"fmt"
	"math/rand"
	"strconv"
)

const(
	North = "north"
	South = "south"
	West = "west"
	East = "east"
	maxMoves = 10000
)

var OppositeDirection = map[string]string {
	North: South,
	South: North,
	West: East,
	East: West,
}

func NewNode(name string) *Node {
	return &Node{
		Name: name,
		Active: true,
		Destinations: make(map[string]*Node),
	}
}

func (n *Node) destroyByAliens(x, y *Alien) {
	n.Active = false
	x.Alive = false
	y.Alive = false

	fmt.Printf("%s has been destroyed by alien %s and alien %s! \n", n.Name, x.Name, y.Name)
}

func (n *Node) accommodateAlien(a *Alien) {
	n.Alien = a
	a.Node = n

	fmt.Printf("Alien %s arrived to %s \n", a.Name, n.Name)
}

func (n *Node) getDestinations() []*Node {
	var destinations []*Node

	for _, d := range n.Destinations {
		if d.Active {
			destinations = append(destinations, d)
		}
	}

	return destinations
}

func (a *Alien) getRandomDestination() *Node {
	destinations := a.Node.getDestinations()
	if len(destinations) == 0 {
		return nil
	}
	r := rand.Intn(len(destinations))

	return destinations[r]
}

func (a *Alien) goTo(newNode *Node) (terminate bool) {
	a.Node.Alien = nil
	a.Node = newNode

	fmt.Printf("Alien %s arrived to %s \n", a.Name, newNode.Name)

	if newNode.Alien != nil {
		a.Node.destroyByAliens(newNode.Alien, a)
		return true
	} else {
		newNode.Alien = a

		return false
	}
}

func (a *Alien) makeMove() (terminate bool) {
	operatedNodes := append(a.Node.getDestinations(), a.Node)
	lockNodes(operatedNodes)
	defer unlockNodes(operatedNodes)

	if !a.Alive {
		fmt.Printf("Alien %s is dead \n", a.Name)

		return true
	}

	// Calculating destinations one more time inside the lock
	d := a.getRandomDestination()
	if d == nil  {
		fmt.Printf("Alien %s is stuck in %s \n", a.Name, a.Node.Name)
		return true
	}
	return a.goTo(d)

	return false
}

func (a *Alien) Dispatch() {
	for i:=0; i<maxMoves; i++ {
		if a.makeMove() {
			break
		}
	}
}

func lockNodes(nodes []*Node) {
	for _, d := range nodes {
		d.mu.Lock()
	}
}

func unlockNodes(nodes []*Node) {
	for _, d := range nodes {
		d.mu.Unlock()
	}
}

func InhabitWithAliens(nodes map[string]*Node, nodeNames []string, n int) []*Alien {
	aliens := make([]*Alien, 0, n)
	for i:=0; i<n; i++ {
		randomKey := nodeNames[rand.Intn(len(nodeNames))]
		randomNode := nodes[randomKey]
		if !randomNode.Active {
			fmt.Printf("Alien %d landed to dead city %s \n", i, randomNode.Name)

			// let's consider if an Alien arrives to a dead city - he dies immediately (we are not creating it)
			continue
		}

		a := &Alien{
			Name: strconv.Itoa(i),
			Node: randomNode,
			Alive: true,
		}
		fmt.Printf("Alien %s landed to %s \n", a.Name, randomNode.Name)

		if randomNode.Alien != nil {
			//node is destroyed immediately
			fmt.Printf("%s has been destroyed by alien %s and alien %s immediately! \n", randomNode.Name, randomNode.Alien.Name, a.Name)

			randomNode.destroyByAliens(randomNode.Alien, a)
		}

		a.Node = randomNode
		randomNode.Alien = a

		aliens = append(aliens, a)
	}

	return aliens
}

func IsValidDirection(d string) bool {
	return d == North || d == South || d == East || d == West
}

func GetNodeNames(nodes map[string]*Node) []string {
	nodeNames := make([]string, 0, len(nodes))
	for k := range nodes {
		nodeNames = append(nodeNames, k)
	}
	return nodeNames
}