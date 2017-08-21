package mailbox

import "github.com/davyxu/actornet/proto"

type Bounded struct {
	q chan interface{}

	hijack func(interface{}) bool
}

func (self *Bounded) Hijack(callback func(interface{}) bool) {

	self.hijack = callback
}

func (self *Bounded) Post(data interface{}) {

	if self.hijack != nil && self.hijack(data) {
		return
	}

	self.q <- data
}

func (self *Bounded) Start(mr MailReceiver) {

	go self.recvThread(mr)
}

func (self *Bounded) recvThread(mr MailReceiver) {
	for {

		msg := <-self.q

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
