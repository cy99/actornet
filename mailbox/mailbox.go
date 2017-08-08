package mailbox

type MailBox interface {
	Push(interface{})
	Recv() interface{}
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

func NewBounded(size int) MailBox {

	return &Bounded{
		q: make(chan interface{}, size),
	}
}
