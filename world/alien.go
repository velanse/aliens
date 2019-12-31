package world

import (
	"math/rand"
	"time"

	"github.com/velanse/aliens/printer"
)

type Alien struct {
	Name  string
	Node  *Node
	Alive bool
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
		return true
	} else {
		newNode.Alien = a

		return false
	}
}

func (a *Alien) makeMove(p printer.Printer, async bool) (terminate bool) {
	operatedNodes := append(a.Node.getDestinations(), a.Node)
	if async {
		lockNodes(operatedNodes)
		defer unlockNodes(operatedNodes)
	}

	if !a.Alive {
		p.Debug("Alien %s is dead \n", a.Name)

		return true
	}

	// Calculating destinations one more time inside the lock
	d := a.getRandomDestination()
	if d == nil {
		p.Debug("Alien %s is stuck in %s \n", a.Name, a.Node.Name)
		return true
	}
	t := a.goTo(d, p)
	return t
}

func (a *Alien) Dispatch(p printer.Printer, delay uint) {
	for i := 0; i < maxMoves; i++ {
		if delay > 0 {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(int(delay))))
		}

		if a.makeMove(p, true) {
			break
		}
	}
}
