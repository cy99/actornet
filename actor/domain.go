package actor

import (
	"errors"
	"fmt"
	"github.com/davyxu/actornet/util"
	"github.com/davyxu/cellnet"
	"reflect"
	"strconv"
)

type Domain struct {
	Name string

	RemoteContext interface{}

	processByID map[string]Process
}

func (self *Domain) String() string {

	if self.RemoteContext != nil {

		return fmt.Sprintf("%s RemoteSes: %d", self.Name, self.RemoteContext.(cellnet.Session).ID())
	}

	return fmt.Sprintf("%s", self.Name)
}

func (self *Domain) allocID() string {

	return strconv.FormatInt(util.GenPersistantID(0), 10)
}

func (self *Domain) Add(p Process) error {

	if _, ok := self.processByID[p.PID().Id]; ok {
		return errors.New("Duplicate id")
	}

	self.processByID[p.PID().Id] = p

	return nil
}

func (self *Domain) GetByID(id string) Process {

	if proc, ok := self.processByID[id]; ok {
		return proc
	}

	return nil
}

func (self *Domain) Get(pid *PID) Process {

	if pid == nil {
		return nil
	}

	if pid.Domain != self.Name {
		return nil
	}

	if proc, ok := self.processByID[pid.Id]; ok {
		return proc
	}

	return nil
}

func (self *Domain) Kill(pid *PID) {

	if pid.Domain != self.Name {
		return
	}

	delete(self.processByID, pid.Id)
}

func (self *Domain) newPID(id string) *PID {
	return &PID{
		Domain: self.Name,
		Id:     id,
	}
}

func (self *Domain) Spawn(t *ActorTemplate) *PID {

	if !inited {
		panic("Call actor.StartSystem first")
	}

	// 生成流水名字
	if t.id == "" {
		t.id = strconv.FormatInt(util.GenPersistantID(0), 10)
	}

	pid := self.newPID(t.id)

	a := t.newActor()

	initor, ok := a.(interface {
		Init(Actor, *PID, *Domain) *LocalProcess
	})

	if !ok {
		panic("Require actor.LocalProcess in your actor visitor: " + reflect.TypeOf(a).Elem().Name())
	}

	proc := initor.Init(a, pid, self)

	if err := self.Add(proc); err != nil {
		panic(err)
	}

	pid.proc = proc

	pproc := t.ppid.ref()
	if pproc != nil {
		pproc.AddChild(pid)
	}

	log.Debugf("#spawn actor: %s", pid.String())

	return pid
}

func newDomain(domain string) *Domain {
	return &Domain{
		Name:        domain,
		processByID: make(map[string]Process),
	}

}
