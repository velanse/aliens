package world

import (
	"math/rand"

	"github.com/velanse/aliens/printer"
)

type Alien struct {
	Name      string
	Node      *Node
	Alive     bool
	MovesMade int
}

func (a *Alien) getRandomDestination() *Node {
	destinations := a.Node.getDestinations()
	if len(destinations) == 0 {
		return nil
	}
	r := rand.Intn(len(destinations))

	return destinations[r]
}

func (a *Alien) goTo(newNode *Node, p printer.Printer) (terminate bool) {
	a.Node.Alien = nil
	a.Node = newNode

	p.Debug("Alien %s arrived to %s \n", a.Name, newNode.Name)

	if newNode.Alien != nil {
		a.Node.destroyByAliens(newNode.Alien, a, p)
		p.Debug("Alien %s is dead \n", a.Name)

		return true
	} else {
		newNode.Alien = a

		return false
	}
}

func (a *Alien) makeMove(p printer.Printer) (terminate bool) {
	if !a.Alive {
		p.Debug("Alien %s is dead \n", a.Name)

		return true
	}

	d := a.getRandomDestination()
	if d == nil {
		p.Debug("Alien %s is stuck in %s \n", a.Name, a.Node.Name)
		return true
	}

	return a.goTo(d, p)
}
