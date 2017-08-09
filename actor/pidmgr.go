package actor

import (
	"errors"
	"github.com/davyxu/actornet/util"
	"strconv"
)

type PIDManager struct {
	Address string

	processByID map[PID]Process
}

func (self *PIDManager) allocID() string {

	return strconv.FormatInt(util.GenPersistantID(0), 10)
}

func (self *PIDManager) Add(p Process) error {

	rawID := p.PID().raw()

	if _, ok := self.processByID[rawID]; ok {
		return errors.New("Duplicate id")
	}

	self.processByID[rawID] = p

	return nil
}

func (self *PIDManager) Get(pid *PID) Process {

	if proc, ok := self.processByID[pid.raw()]; ok {
		return proc
	}

	return nil
}

func NewPIDManager(address string) *PIDManager {
	return &PIDManager{
		Address:     address,
		processByID: make(map[PID]Process),
	}

}

var localPIDManager = NewPIDManager("local")
