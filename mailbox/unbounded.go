package mailbox

import (
	"runtime"
	"sync/atomic"
)

const (
	idle int32 = iota
	running
)

type Unbounded struct {
	hijack func(interface{}) bool

	q *Queue

	schedulerStatus int32
	userMessages    int32

	receiver MailReceiver
}

func (self *Unbounded) Hijack(callback func(interface{}) bool) {

	self.hijack = callback
}

func (self *Unbounded) Post(data interface{}) {

	if self.hijack != nil && self.hijack(data) {
		return
	}

	self.q.Push(data)

	atomic.AddInt32(&self.userMessages, 1)

	self.schedule()
}

func (self *Unbounded) processMessage() {
	for {

		self.run()

		atomic.StoreInt32(&self.schedulerStatus, idle)
		user := atomic.LoadInt32(&self.userMessages)
		// check if there are still messages to process (sent after the message loop ended)
		if user > 0 {
			// try setting the mailbox back to running
			if atomic.CompareAndSwapInt32(&self.schedulerStatus, idle, running) {
				//	fmt.Printf("looping %v %v %v\n", sys, user, m.suspended)
				continue
			}
		}

		break
	}
}

func (self *Unbounded) schedule() {

	if atomic.CompareAndSwapInt32(&self.schedulerStatus, idle, running) {
		go self.processMessage()
	}
}

func (self *Unbounded) run() {

	for {

		runtime.Gosched()

		data, ok := self.q.Pop()

		if ok {

			self.receiver.OnRecv(data)

		} else {

			// 处理完一批次后退出
			break
		}
	}
}

func (self *Unbounded) Start(mr MailReceiver) {

	self.receiver = mr
}

func NewUnbouned() MailBox {
	return &Unbounded{
		q: NewQueue(10),
	}
}
