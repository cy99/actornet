package util

import "sync"

type Value interface{}

type Future struct {
	wg sync.WaitGroup
	r  Value
}

func NewFuture() *Future {

	self := &Future{}

	self.wg.Add(1)

	return self

}

func (self *Future) Done(r Value) {

	self.r = r
	self.wg.Done()
}

func (self *Future) Get() Value {

	self.wg.Wait()

	return self.r
}
