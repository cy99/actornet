package actor

import (
	"github.com/davyxu/actornet/mailbox"
	"github.com/davyxu/actornet/proto"
)

type Process interface {
	Notify(*Message)

	Call(*Message) *Message

	Stop()

	PID() *PID
}

type localProcess struct {
	mailbox mailbox.MailBox

	pid *PID

	a Actor
}

func (self *localProcess) notifySystem(data interface{}) {
	self.Notify(&Message{
		Data:      data,
		SourcePID: self.pid,
		TargetPID: self.pid,
	})
}

func (self *localProcess) Call(m *Message) *Message {

	m.CallID = AllocRPCSeq()

	reply := make(chan *Message)

	self.mailbox.Hijack(func(rpcBack interface{}) bool {

		rpcMsg := rpcBack.(*Message)
		if rpcMsg.CallID == m.CallID {
			reply <- rpcMsg
			return true
		}

		return false
	})

	m.TargetPID.Notify(m)

	msgReply := <-reply

	self.mailbox.Hijack(nil)

	return msgReply
}

func (self *localProcess) PID() *PID {
	return self.pid
}

func (self *localProcess) Notify(m *Message) {

	//log.Debugf("[%s] LocalProcess.Notify %v", self.pid.String(), *m)

	self.mailbox.Push(m)
}

func (self *localProcess) Stop() {

	self.notifySystem(&proto.Stop{})
}

func (self *localProcess) OnRecv(data interface{}) {

	msg := data.(*Message)

	//log.Debugf("[%s] LocalProcess.Notify %v", self.pid.String(), *msg)

	self.a.OnRecv(msg)
}

func newLocalProcess(a Actor, pid *PID) *localProcess {

	self := &localProcess{
		mailbox: mailbox.NewBounded(10),
		a:       a,
		pid:     pid,
	}

	self.mailbox.Start(self)

	self.notifySystem(&proto.Start{})

	return self
}
