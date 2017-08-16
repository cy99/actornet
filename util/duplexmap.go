package util

import (
	"errors"
	"sync"
)

// 双向关联map
type DuplexMap struct {
	main2slave map[interface{}]interface{}
	slave2main map[interface{}]interface{}
	guard      sync.RWMutex
}

func (self *DuplexMap) Add(main, slave interface{}) {

	self.guard.Lock()
	self.main2slave[main] = slave
	self.slave2main[slave] = main
	self.guard.Unlock()
}

var ErrNotExists = errors.New("Not exists")

func (self *DuplexMap) RemoveByMain(main interface{}) (interface{}, error) {

	self.guard.Lock()

	defer self.guard.Unlock()

	if slave, ok := self.main2slave[main]; ok {
		delete(self.main2slave, main)
		delete(self.slave2main, slave)

		return slave, nil
	} else {

		return nil, ErrNotExists
	}

}

func (self *DuplexMap) RemoveBySlave(slave interface{}) (interface{}, error) {

	self.guard.Lock()

	defer self.guard.Unlock()

	if main, ok := self.slave2main[slave]; ok {
		delete(self.main2slave, main)
		delete(self.slave2main, main)
		return main, nil
	} else {

		return nil, ErrNotExists
	}

}

func (self *DuplexMap) SlaveByMain(main interface{}) (interface{}, bool) {

	self.guard.RLock()
	defer self.guard.RUnlock()

	v, ok := self.main2slave[main]
	return v, ok
}

func (self *DuplexMap) MainBySlave(slave interface{}) (interface{}, bool) {

	self.guard.RLock()
	defer self.guard.RUnlock()

	v, ok := self.slave2main[slave]
	return v, ok
}

func (self *DuplexMap) Visit(callback func(main, slave interface{}) bool) {

	for main, slave := range self.main2slave {

		if !callback(main, slave) {
			break
		}
	}

}

func NewDuplexMap() *DuplexMap {

	return &DuplexMap{
		main2slave: make(map[interface{}]interface{}),
		slave2main: make(map[interface{}]interface{}),
	}
}
