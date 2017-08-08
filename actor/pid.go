package actor

type PID struct {
	Index int64
	ID    int64

	proc Process
}

func (self *PID) Send(data interface{}) {

	self.proc.Send(data)
}
