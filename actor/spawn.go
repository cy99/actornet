package actor

import (
	"github.com/davyxu/actornet/util"
	"strconv"
)

type ActorTemplate struct {
	name string
	a    Actor
	pid  *PID
}

func (self *ActorTemplate) WithName(name string) *ActorTemplate {
	self.name = name
	return self
}

func (self *ActorTemplate) WithFunc(f ActorFunc) *ActorTemplate {
	self.a = f
	return self
}

func (self *ActorTemplate) WithInstance(a Actor) *ActorTemplate {
	self.a = a
	return self
}

func (self *ActorTemplate) Spawn() *PID {

	// 生成流水名字
	if self.name == "" {
		self.name = strconv.FormatInt(util.GenPersistantID(0), 10)
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

	pid := NewLocalPID(t.name)

	proc := newLocalProcess(t.a, pid)

	if err := LocalPIDManager.Add(proc); err != nil {
		panic(err)
	}

	pid.proc = proc

	return proc.pid
}
