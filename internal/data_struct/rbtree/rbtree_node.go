package rbtree

import (
	"os_lab6-8/internal/timer"
)

type nodeColor int

const (
	red nodeColor = iota
	black
)

type node struct {
	id int
	t  *timer.TTimer

	color nodeColor

	left   *node
	right  *node
	parent *node
}

// leafNode is an artificial node guarantees that any of leaves is black
var leafNode node = node{
	0, nil, black, nil, nil, nil,
}

// newNode is a constructor of the type node
func newNode(id int, nodeColor nodeColor) *node {
	n := new(node)
	n.id = id
	n.color = nodeColor
	n.t = new(timer.TTimer)

	n.left = &leafNode
	n.right = &leafNode
	n.parent = nil

	return n
}

// find returns timer from the node with specified id
func (n *node) find(id int) *timer.TTimer {
	if n == &leafNode {
		return nil
	}
	var current = n
	// default searching of the node in BST
	for current.id != id {
		if id > current.id && current.right != &leafNode {
			current = current.right
		} else if id < current.id && current.left != &leafNode {
			current = current.left
		} else {
			return nil
		}
	}
	return current.t
}
