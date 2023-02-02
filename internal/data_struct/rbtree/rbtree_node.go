package rbtree

import (
	"os_lab6-8/internal/timer"
)

type color int

const (
	red color = iota
	black
)

type node struct {
	id int
	t  *timer.TTimer

	nodeColor color

	leftChild  *node
	rightChild *node
	parentNode *node
}

// leafNode is an artificial node guarantees that any of leaves is black
var leafNode node = node{
	0, nil, black, nil, nil, nil,
}

// newNode is a constructor of the type node
func newNode(id int, nodeColor color) *node {
	n := new(node)
	n.id = id
	n.nodeColor = nodeColor
	n.t = new(timer.TTimer)

	n.leftChild = &leafNode
	n.rightChild = &leafNode
	n.parentNode = nil

	return n
}

// find returns timer from the node with specified id
func (n *node) find(id int) *timer.TTimer {
	var current = n
	// default searching of the node in BST
	for current.id != id {
		if id > current.id && current.rightChild != &leafNode {
			current = current.rightChild
		} else if id < current.id && current.leftChild != &leafNode {
			current = current.leftChild
		} else {
			return nil
		}
	}
	return current.t
}
