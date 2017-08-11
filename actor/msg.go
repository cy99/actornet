package actor


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

	// 标记已经处理
	self.CallID = 0
}
