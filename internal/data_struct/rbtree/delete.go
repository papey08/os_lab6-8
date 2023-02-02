package rbtree

func (rbt *RBTree) fixDelete(x *node) {
	for x != rbt.root && x.color == black {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == red {
				w.color = black
				x.parent.color = red
				rbt.rotateLeft(x.parent)
				w = x.parent.right
			}
			if w.left.color == black && w.right.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.right.color == black {
					w.left.color = black
					w.color = red
					rbt.rotateRight(w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = black
				w.right.color = black
				rbt.rotateLeft(x.parent)
				x = rbt.root
			}
		} else {
			w := x.parent.left
			if w.color == red {
				w.color = black
				x.parent.color = red
				rbt.rotateRight(x.parent)
				w = x.parent.left
			}
			if w.right.color == black && w.left.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.left.color == black {
					w.right.color = black
					w.color = red
					rbt.rotateLeft(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = black
				w.left.color = black
				rbt.rotateRight(x.parent)
				x = rbt.root
			}
		}
	}
	x.color = black
}

func (rbt *RBTree) delete(z *node) bool {
	var x, y *node
	if z == nil || z == &leafNode {
		return false
	}
	if z.left == &leafNode || z.right == &leafNode {
		y = z
	} else {
		y = z.right
		for y.left != &leafNode {
			y = y.left
		}
	}
	if y.left != &leafNode {
		x = y.left
	} else {
		x = y.right
	}
	x.parent = y.parent
	if y.parent != nil {
		if y == y.parent.left {
			y.parent.left = x
		} else {
			y.parent.right = x
		}
	} else {
		rbt.root = x
	}
	if y != z {
		z.id = y.id
		z.t = y.t
	}
	if y.color == black {
		rbt.fixDelete(x)
	}
	return true
}
