package data_struct

import (
	"errors"
	"os_lab6-8/internal/timer"
)

// DefaultMap is a cover over map from standard golang library implements DataStructure interface
type DefaultMap struct {
	tmap map[int]*timer.TTimer
}

func NewDefaultMap() *DefaultMap {
	newDefaultMap := new(DefaultMap)
	newDefaultMap.tmap = make(map[int]*timer.TTimer)
	return newDefaultMap
}

func (d *DefaultMap) GetTime(id int) (int, error) {
	if _, ok := d.tmap[id]; ok {
		return d.tmap[id].GetTime(), nil
	} else {
		return 0, errors.New("node not exist")
	}
}

func (d *DefaultMap) StartTimer(id int) error {
	if _, ok := d.tmap[id]; ok {
		d.tmap[id].Start()
		return nil
	} else {
		return errors.New("node not exist")
	}
}

func (d *DefaultMap) InsertNode(id int) error {
	if _, ok := d.tmap[id]; ok {
		return errors.New("node already exists")
	} else {
		d.tmap[id] = new(timer.TTimer)
		return nil
	}
}

func (d *DefaultMap) DeleteNode(id int) error {
	if _, ok := d.tmap[id]; !ok {
		return errors.New("node not exists")
	} else {
		delete(d.tmap, id)
		return nil
	}
}

func (d *DefaultMap) Length() int {
	return len(d.tmap)
}
