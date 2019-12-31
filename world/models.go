package world

import "sync"

type Alien struct {
	Name string
	Node *Node
	Alive bool
}

type Node struct {
	Name string
	Alien *Alien
	Destinations map[string]*Node
	Active bool
	mu  sync.RWMutex
}