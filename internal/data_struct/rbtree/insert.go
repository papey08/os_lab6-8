package rbtree

import (
	"errors"
	"strconv"
)

func (rbt *RBTree) rotateLeft(x *node) {
	y := x.rightChild
	x.rightChild = y.leftChild
	if y.leftChild != &leafNode {
		y.leftChild.parentNode = x
	}
	if y != &leafNode {
		y.parentNode = x.parentNode
	}
	if x.parentNode != nil {
		if x == x.parentNode.leftChild {
			x.parentNode.leftChild = y
		} else {
			x.parentNode.rightChild = y
		}
	} else {
		rbt.root = y
	}
	y.leftChild = x
	if x != &leafNode {
		x.parentNode = y
	}
}

func (rbt *RBTree) rotateRight(x *node) {
	y := x.leftChild
	x.leftChild = y.rightChild
	if y.rightChild != &leafNode {
		y.rightChild.parentNode = x
	}
	if y != &leafNode {
		y.parentNode = x.parentNode
	}
	if x.parentNode != nil {
		if x == x.parentNode.rightChild {
			x.parentNode.rightChild = y
		} else {
			x.parentNode.leftChild = y
		}
	} else {
		rbt.root = y
	}
	y.rightChild = x
	if x != &leafNode {
		x.parentNode = y
	}
}

// fixInsert rotates & recolors some nodes for implementing rbtree's properties if need
func (rbt *RBTree) fixInsert(x *node) {
	for x != rbt.root && x.parentNode.nodeColor == red {
		if x.parentNode == x.parentNode.parentNode.leftChild {
			y := x.parentNode.parentNode.rightChild
			if y.nodeColor == red { // found red uncle => recoloring, no rotations
				x.parentNode.nodeColor = black
				y.nodeColor = black
				x.parentNode.parentNode.nodeColor = red
				x = x.parentNode.parentNode
			} else { // recoloring & rotating
				if x == x.parentNode.rightChild {
					x = x.parentNode
					rbt.rotateLeft(x)
				}
				x.parentNode.nodeColor = black
				x.parentNode.parentNode.nodeColor = red
				rbt.rotateRight(x.parentNode.parentNode)
			}
		} else {
			y := x.parentNode.parentNode.leftChild
			if y.nodeColor == red { // found red uncle
				x.parentNode.nodeColor = black
				y.nodeColor = black
				x.parentNode.parentNode.nodeColor = red
				x = x.parentNode.parentNode
			} else { // recoloring & rotating
				if x == x.parentNode.leftChild {
					x = x.parentNode
					rbt.rotateRight(x)
				}
				x.parentNode.nodeColor = black
				x.parentNode.parentNode.nodeColor = red
				rbt.rotateLeft(x.parentNode.parentNode)
			}
		}
	}
	rbt.root.nodeColor = black
}

func (rbt *RBTree) insert(id int) error {
	var current, parent *node
	current = rbt.root
	parent = nil
	// searching a place where to insert a new node like in BFS
	for current != &leafNode {
		if id < current.id {
			parent = current
			current = current.leftChild
		} else if id > current.id {
			parent = current
			current = current.rightChild
		} else {
			return errors.New("node with id " + strconv.Itoa(id) +
				" already exists")
		}
	}
	// inserting new node
	x := newNode(id, red)
	x.parentNode = parent
	if parent != nil {
		if parent.id > id {
			parent.leftChild = x
		} else {
			parent.rightChild = x
		}
	} else {
		rbt.root = new(node)
		rbt.root = x
	}
	// calling a function to fix rbtree
	rbt.fixInsert(x)
	return nil
}
