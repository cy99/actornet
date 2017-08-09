package actor

import (
	"errors"
	"github.com/davyxu/actornet/util"
	"strconv"
)

type PIDManager struct {
	Address string

	processByID map[string]Process
}

func (self *PIDManager) allocID() string {

	return strconv.FormatInt(util.GenPersistantID(0), 10)
}

func (self *PIDManager) Add(p Process) error {

	if _, ok := self.processByID[p.PID().Id]; ok {
		return errors.New("Duplicate id")
	}

	self.processByID[p.PID().Id] = p

	return nil
}

func (self *PIDManager) Get(id string) Process {

	if proc, ok := self.processByID[id]; ok {
		return proc
	}

	return nil
}

func NewPIDManager(address string) *PIDManager {
	return &PIDManager{
		Address:     address,
		processByID: make(map[string]Process),
	}

}

var LocalPIDManager = NewPIDManager("localhost")
