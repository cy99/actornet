package util

import "container/list"

// 类似于C#的Action, 一种delegate
type Delegate struct {
	list *list.List
}

func (self *Delegate) Invoke(args ...interface{}) error {

	if self.list == nil {
		return nil
	}

	for e := self.list.Front(); e != nil; e = e.Next() {

		v := e.Value.(func(...interface{}) error)

		if err := v(args...); err != nil {
			return err
		}

	}

	return nil
}

func (self *Delegate) Remove(element *list.Element) {

	if self.list == nil {
		return
	}

	if element == nil {
		return
	}

	self.list.Remove(element)
}

func (self *Delegate) Add(callback func(...interface{}) error) *list.Element {

	if self.list == nil {
		self.list = list.New()
	}

	return self.list.PushBack(callback)
}

func (self *Delegate) Clear() {
	if self.list == nil {
		return
	}
	self.list.Init()
}
