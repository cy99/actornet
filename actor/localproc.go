package actor

import (
	"github.com/davyxu/actornet/mailbox"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/actornet/util"
)

type Process interface {
	Tell(interface{})

	Stop()

	PID() *PID
}

type LocalProcess struct {
	mailbox mailbox.MailBox

	pid *PID

	a Actor
}

func (self *LocalProcess) Serialize(ser Serializer) {
	self.a.(Serializable).OnSerialize(ser)
}

func (self *LocalProcess) notifySystem(data interface{}) {
	self.Tell(&Message{
		Data:      data,
		SourcePID: self.pid,
		TargetPID: self.pid,
	})
}

func (self *LocalProcess) CreateRPC(waitCallID int64) *util.Future {

	f := util.NewFuture()

	self.mailbox.Hijack(func(rpcBack interface{}) bool {

		rpcMsg := rpcBack.(*Message)
		if rpcMsg.CallID == waitCallID {

			self.mailbox.Hijack(nil)
			f.Done(rpcMsg)
			return true
		}

		return false
	})

	return f
}

func (self *LocalProcess) PID() *PID {
	return self.pid
}

func (self *LocalProcess) Tell(data interface{}) {

	if EnableDebug {
		log.Debugf("#notify %s", data.(Context).String())
	}

	self.mailbox.Post(data)
}

func (self *LocalProcess) Stop() {

	self.notifySystem(&proto.Stop{})
}

func (self *LocalProcess) OnRecv(data interface{}) {

	ctx := data.(Context)

	if EnableDebug {
		log.Debugf("#recv %s", data.(Context).String())
	}

	self.a.OnRecv(ctx)
}

func (self *LocalProcess) Init(a Actor, pid *PID) Process {

	self.mailbox = mailbox.NewUnbouned()
	self.a = a
	self.pid = pid

	self.mailbox.Start(self)

	self.notifySystem(&proto.Start{})

	return self
}
