package actor

type Context interface {
	Msg() interface{}

	Source() *PID

	Self() *PID

	Reply(data interface{})
}

type Message struct {
	Data      interface{}
	SourcePID *PID
	TargetPID *PID
	CallID    int64
}

func (self *Message) Msg() interface{} {
	return self.Data
}

func (self *Message) Source() *PID {
	return self.SourcePID
}

func (self *Message) Self() *PID {
	return self.TargetPID
}

func (self *Message) Reply(data interface{}) {

	self.SourcePID.Notify(&Message{
		Data:      data,
		TargetPID: self.SourcePID,
		SourcePID: self.TargetPID,
		CallID:    self.CallID,
	})
}
