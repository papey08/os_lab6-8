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
func (n *node) find(id int) *node {
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
	return current
}

func (rbt *RBTree) rotateLeft(x *node) {
	y := x.right
	x.right = y.left
	if y.left != &leafNode {
		y.left.parent = x
	}
	if y != &leafNode {
		y.parent = x.parent
	}
	if x.parent != nil {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	} else {
		rbt.root = y
	}
	y.left = x
	if x != &leafNode {
		x.parent = y
	}
}

func (rbt *RBTree) rotateRight(x *node) {
	y := x.left
	x.left = y.right
	if y.right != &leafNode {
		y.right.parent = x
	}
	if y != &leafNode {
		y.parent = x.parent
	}
	if x.parent != nil {
		if x == x.parent.right {
			x.parent.right = y
		} else {
			x.parent.left = y
		}
	} else {
		rbt.root = y
	}
	y.right = x
	if x != &leafNode {
		x.parent = y
	}
}
