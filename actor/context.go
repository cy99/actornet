package actor

type Context interface {
	Msg() interface{}

	Source() *PID

	Self() *PID
}

type mailContext struct {
	msg  interface{}
	src  *PID
	self *PID
}

func (self *mailContext) Msg() interface{} {
	return self.msg
}

func (self *mailContext) Source() *PID {
	return self.src
}

func (self *mailContext) Self() *PID {
	return self.self
}
