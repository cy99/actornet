package actor

type PID struct {
	Index int64
	ID    int64

	p Process
}

func (self *PID) Send(data interface{}) {

	self.p.Send(data)
}
