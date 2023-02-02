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
	if n := rbt.root.find(id); n == nil {
		return 0, errors.New("node with id " + strconv.Itoa(id) + " not exists")
	} else {
		return n.t.GetTime(), nil
	}
}

func (rbt *RBTree) StartTimer(id int) error {
	if n := rbt.root.find(id); n == nil {
		return errors.New("node with id " + strconv.Itoa(id) + " not exists")
	} else {
		n.t.Start()
		return nil
	}
}

func (rbt *RBTree) PauseTimer(id int) error {
	if n := rbt.root.find(id); n == nil {
		return errors.New("node with id " + strconv.Itoa(id) + " not exists")
	} else {
		n.t.Pause()
		return nil
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

func (rbt *RBTree) DeleteNode(id int) error {
	if !rbt.delete(rbt.root.find(id)) {
		return errors.New("node with id " + strconv.Itoa(id) + " not exists")
	} else {
		rbt.size--
		return nil
	}
}

func (rbt *RBTree) ResetTimer(id int) error {
	if n := rbt.root.find(id); n == nil {
		return errors.New("node with id " + strconv.Itoa(id) + " not exists")
	} else {
		n.t.Reset()
		return nil
	}
}

func (rbt *RBTree) Length() int {
	return rbt.size
}
