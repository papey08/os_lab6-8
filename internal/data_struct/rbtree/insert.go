package rbtree

import (
	"errors"
	"strconv"
)

// fixInsert rotates & recolors some nodes for implementing rbtree's properties if need
func (rbt *RBTree) fixInsert(x *node) {
	for x != rbt.root && x.parent.color == red {
		if x.parent == x.parent.parent.left {
			y := x.parent.parent.right
			if y.color == red { // found red uncle => recoloring, no rotations
				x.parent.color = black
				y.color = black
				x.parent.parent.color = red
				x = x.parent.parent
			} else { // recoloring & rotating
				if x == x.parent.right {
					x = x.parent
					rbt.rotateLeft(x)
				}
				x.parent.color = black
				x.parent.parent.color = red
				rbt.rotateRight(x.parent.parent)
			}
		} else {
			y := x.parent.parent.left
			if y.color == red { // found red uncle
				x.parent.color = black
				y.color = black
				x.parent.parent.color = red
				x = x.parent.parent
			} else { // recoloring & rotating
				if x == x.parent.left {
					x = x.parent
					rbt.rotateRight(x)
				}
				x.parent.color = black
				x.parent.parent.color = red
				rbt.rotateLeft(x.parent.parent)
			}
		}
	}
	rbt.root.color = black
}

func (rbt *RBTree) insert(id int) error {
	var current, parent *node
	current = rbt.root
	parent = nil
	// searching a place where to insert a new node like in BFS
	for current != &leafNode {
		if id < current.id {
			parent = current
			current = current.left
		} else if id > current.id {
			parent = current
			current = current.right
		} else {
			return errors.New("node with id " + strconv.Itoa(id) +
				" already exists")
		}
	}
	// inserting new node
	x := newNode(id, red)
	x.parent = parent
	if parent != nil {
		if parent.id > id {
			parent.left = x
		} else {
			parent.right = x
		}
	} else {
		rbt.root = new(node)
		rbt.root = x
	}
	// calling a function to fix rbtree
	rbt.fixInsert(x)
	return nil
}
