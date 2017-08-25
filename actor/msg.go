package actor

import (
	"fmt"
	"github.com/davyxu/goobjfmt"
)

// 信息上下文
type Context interface {

	// 消息本体
	Msg() interface{}

	// 消息来源方，可能为空
	Source() *PID

	// Actor本体的PID
	Self() *PID

	Parent() *PID

	// 当对方用Call调用时， 需要用Reply回应
	Reply(data interface{})

	String() string
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

func (self *Message) Parent() *PID {
	return self.TargetPID.ref().ParentPID()
}

func (self *Message) String() string {

	var source string

	if self.SourcePID != nil {
		source = self.SourcePID.String()
	}

	return fmt.Sprintf("(%s)->(%s) callid:%d | (%T) %s", source, self.TargetPID.String(), self.CallID, self.Data, goobjfmt.CompactTextString(self.Data))
}

func (self *Message) Reply(data interface{}) {

	self.SourcePID.ref().Tell(&Message{
		Data:      data,
		TargetPID: self.SourcePID,
		SourcePID: self.TargetPID,
		CallID:    self.CallID,
	})

	// 标记已经处理
	self.CallID = 0
}
