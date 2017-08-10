package actor

import (
	"github.com/davyxu/actornet/mailbox"
	"github.com/davyxu/actornet/proto"
)

type Process interface {
	Notify(data interface{}, sender *PID)
	Stop()

	PID() *PID
}

type localProcess struct {
	mailbox mailbox.MailBox

	pid PID

	a Actor
}

func (self *localProcess) PID() *PID {
	return &self.pid
}

func (self *localProcess) Notify(data interface{}, sender *PID) {

	self.mailbox.Push(&mailContext{
		msg:  data,
		src:  sender,
		self: &self.pid,
	})
}

func (self *localProcess) Stop() {

	self.Notify(&proto.Stop{}, &self.pid)
}

func (self *localProcess) OnRecv(msg interface{}) {

	self.a.OnRecv(msg.(Context))
}

func NewLocalProcess(a Actor, pid PID) *localProcess {

	self := &localProcess{
		mailbox: mailbox.NewBounded(10),
		a:       a,
		pid:     pid,
	}

	self.pid.proc = self

	self.mailbox.Start(self)

	self.Notify(&proto.Start{}, &self.pid)

	return self
}
