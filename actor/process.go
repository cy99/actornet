package actor

import (
	"github.com/davyxu/actornet/mailbox"
	"github.com/davyxu/actornet/proto"
)

type Process interface {
	Send(interface{})
	Stop()
}

type localProcess struct {
	mailbox mailbox.MailBox

	a Actor

	thisMsg interface{}
}

func (self *localProcess) Send(data interface{}) {

	self.mailbox.Push(data)
}

func (self *localProcess) Stop() {

	self.Send(&proto.Stop{})
}

func (self *localProcess) Recv(msg interface{}) {

	self.thisMsg = msg
	self.a.Receive(self)
	self.thisMsg = nil
}

func (self *localProcess) Msg() interface{} {

	return self.thisMsg
}

func NewLocalProcess(a Actor) Process {

	self := &localProcess{
		mailbox: mailbox.NewBounded(10),
		a:       a,
	}

	self.mailbox.Start(self)

	self.Send(&proto.Start{})

	return self
}
