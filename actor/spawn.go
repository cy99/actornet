package actor

import (
	"github.com/davyxu/actornet/util"
	"reflect"
	"strconv"
)

type actorFuncHelper struct {
	LocalProcess
	f func(c Context)
}

func (self *actorFuncHelper) OnRecv(c Context) {
	self.f(c)
}

type ActorTemplate struct {
	id  string
	ac  ActorCreator
	ins Actor
	pid *PID
}

func (self *ActorTemplate) WithID(id string) *ActorTemplate {
	self.id = id
	return self
}

func (self *ActorTemplate) WithCreator(ac ActorCreator) *ActorTemplate {
	self.ac = ac
	return self
}

func (self *ActorTemplate) WithInstance(a Actor) *ActorTemplate {
	self.ins = a
	return self
}

func (self *ActorTemplate) WithFunc(f func(c Context)) *ActorTemplate {
	self.ac = func() Actor {
		return &actorFuncHelper{f: f}
	}
	return self
}

func (self *ActorTemplate) newActor() Actor {

	if self.ac != nil {
		return self.ac()
	}

	if self.ins != nil {
		return self.ins
	}

	return nil
}

func (self *ActorTemplate) Spawn() *PID {

	// 生成流水名字
	if self.id == "" {
		self.id = strconv.FormatInt(util.GenPersistantID(0), 10)
	}

	return spawn(self)
}

func NewTemplate() *ActorTemplate {
	return &ActorTemplate{}
}

func spawn(t *ActorTemplate) *PID {

	if !inited {
		panic("Call actor.StartSystem first")
	}

	pid := NewLocalPID(t.id)

	a := t.newActor()

	initor, ok := a.(interface {
		Init(Actor, *PID) Process
	})

	if !ok {
		panic("Require actor.LocalProcess in your actor visitor: " + reflect.TypeOf(a).Elem().Name())
	}

	proc := initor.Init(a, pid)

	if err := LocalPIDManager.Add(proc); err != nil {
		panic(err)
	}

	pid.proc = proc

	log.Debugf("#spawn actor: %s", pid.String())

	return pid
}
