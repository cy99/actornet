package actor

import (
	"github.com/davyxu/actornet/mailbox"
	"github.com/davyxu/actornet/proto"
)

type Process interface {
	Notify(*Message)

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

func (self *localProcess) BeginHijack(waitCallID int64) chan *Message {

	reply := make(chan *Message)

	self.mailbox.Hijack(func(rpcBack interface{}) bool {

		rpcMsg := rpcBack.(*Message)
		if rpcMsg.CallID == waitCallID {
			reply <- rpcMsg
			return true
		}

		return false
	})

	return reply
}

func (self *localProcess) EndHijack(reply chan *Message) *Message {

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

	needReply := msg.CallID != 0

	self.a.OnRecv(msg)

	if needReply {
		log.Errorln("message not reply", *msg)
	}
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
