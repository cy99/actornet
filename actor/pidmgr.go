package actor

import (
	"errors"
	"strconv"
	"sync"

	"github.com/davyxu/actornet/util"
)

type PIDManager struct {
	Domain string

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

func (self *PIDManager) GetByID(id string) Process {

	if proc, ok := self.processByID[id]; ok {
		return proc
	}

	return nil
}

func (self *PIDManager) Get(pid *PID) Process {

	if pid.Domain != self.Domain {
		return nil
	}

	if proc, ok := self.processByID[pid.Id]; ok {
		return proc
	}

	return nil
}

func (self *PIDManager) Remove(pid *PID) {

	if pid.Domain != self.Domain {
		return
	}

	delete(self.processByID, pid.Id)
}

func NewPIDManager(domain string) *PIDManager {
	return &PIDManager{
		Domain:      domain,
		processByID: make(map[string]Process),
	}

}

var (
	LocalPIDManager *PIDManager

	pidmgrByDomain      = map[string]*PIDManager{}
	pidmgrByDomainGuard sync.Mutex
)

// 找到对应地址的远程pid管理器
func remotePIDManager(domain string) *PIDManager {

	pidmgrByDomainGuard.Lock()

	defer pidmgrByDomainGuard.Unlock()

	if mgr, ok := pidmgrByDomain[domain]; ok {
		return mgr
	}

	mgr := NewPIDManager(domain)

	pidmgrByDomain[domain] = mgr

	return mgr
}

func init() {

	OnReset.Add(func(...interface{}) error {

		LocalPIDManager = NewPIDManager("localhost")
		pidmgrByDomain = map[string]*PIDManager{}

		return nil
	})

}
