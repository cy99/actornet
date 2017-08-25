package actor

type ActorTemplate struct {
	id   string
	ac   ActorCreator
	ins  Actor
	pid  *PID
	ppid *PID
}

func (self *ActorTemplate) WithParent(pid *PID) *ActorTemplate {
	self.ppid = pid
	return self
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
	self.ac = newFuncActor(f)

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

func NewTemplate() *ActorTemplate {
	return &ActorTemplate{}
}
