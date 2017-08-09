package mailbox

import "github.com/davyxu/actornet/proto"

type MailReceiver interface {
	OnRecv(interface{})
}

type MailBox interface {
	Push(interface{})
	Recv() interface{}
	Start(MailReceiver)
}

type Bounded struct {
	q chan interface{}
}

func (self *Bounded) Push(data interface{}) {
	self.q <- data
}

func (self *Bounded) Recv() interface{} {

	return <-self.q
}

func (self *Bounded) Start(mr MailReceiver) {

	go self.recvThread(mr)
}

func (self *Bounded) recvThread(mr MailReceiver) {
	for {

		msg := self.Recv()

		switch msg.(type) {
		case proto.Stop:
			goto EndRecv
		default:
			mr.OnRecv(msg)
		}

	}

EndRecv:
}

func NewBounded(size int) MailBox {

	return &Bounded{
		q: make(chan interface{}, size),
	}
}
