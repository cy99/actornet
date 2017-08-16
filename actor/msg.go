package actor

import (
	"fmt"
	"github.com/davyxu/goobjfmt"
)

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

func (self *Message) String() string {

	var source string

	if self.SourcePID != nil {
		source = self.SourcePID.String()
	}

	return fmt.Sprintf("(%s)->(%s) callid:%d | (%T) %s", source, self.TargetPID.String(), self.CallID, self.Data, goobjfmt.CompactTextString(self.Data))
}

func (self *Message) Reply(data interface{}) {

	self.SourcePID.ref().Notify(&Message{
		Data:      data,
		TargetPID: self.SourcePID,
		SourcePID: self.TargetPID,
		CallID:    self.CallID,
	})

	// 标记已经处理
	self.CallID = 0
}
