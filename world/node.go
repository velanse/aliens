package world

import (
	"github.com/velanse/aliens/printer"
	"sync"
)

type Node struct {
	Name         string
	Alien        *Alien
	Destinations map[string]*Node
	Active       bool
	mu           sync.RWMutex
}

func NewNode(name string) *Node {
	return &Node{
		Name:         name,
		Active:       true,
		Destinations: make(map[string]*Node),
	}
}

func (n *Node) destroyByAliens(x, y *Alien, p printer.Printer) {
	n.Active = false
	x.Alive = false
	y.Alive = false

	p.Printf("%s has been destroyed by alien %s and alien %s! \n", n.Name, x.Name, y.Name)
}

func (n *Node) accommodateAlien(a *Alien, p printer.Printer) {
	n.Alien = a
	a.Node = n

	p.Debug("Alien %s arrived to %s \n", a.Name, n.Name)
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

func GetNodeNames(nodes map[string]*Node) []string {
	nodeNames := make([]string, 0, len(nodes))
	for k := range nodes {
		nodeNames = append(nodeNames, k)
	}
	return nodeNames
}
