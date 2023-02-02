package rbtree

import (
	"errors"
	"strconv"
)

type RBTree struct {
	root *node
	size int
}

func NewRBTree() *RBTree {
	rbt := new(RBTree)
	rbt.root = &leafNode
	rbt.size = 0
	return rbt
}

func (rbt *RBTree) GetTime(id int) (int, error) {
	if t := rbt.root.find(id); t == nil {
		return 0, errors.New("node with id " + strconv.Itoa(id) + " not exist")
	} else {
		return t.GetTime(), nil
	}
}

func (rbt *RBTree) InsertNode(id int) error {
	if err := rbt.insert(id); err != nil {
		return err
	} else {
		rbt.size++
		return nil
	}
}
