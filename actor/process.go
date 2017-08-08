package actor

import "github.com/davyxu/actornet/mailbox"

type Process interface {
	Send(interface{})
	Stop()
}

type localProcess struct {
	mailbox mailbox.MailBox

	a Actor
}

func (self *localProcess) Send(interface{}) {

}

func (self *localProcess) Stop() {

}

func NewLocalProcess(a Actor) Process {

	return &localProcess{
		mailbox: mailbox.NewBounded(10),
		a:       a,
	}
}
