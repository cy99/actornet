package socket

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/actornet/util"
	"github.com/davyxu/cellnet"
)

type socketProcess struct {
	pid *actor.PID

	hijack func(*actor.Message) bool
}

func (self *socketProcess) PID() *actor.PID {
	return self.pid
}

func (self *socketProcess) Notify(data interface{}) {

	m := data.(*actor.Message)

	msgdata, msgid, err := cellnet.EncodeMessage(m.Data)
	if err != nil {
		log.Errorln(err)
		return
	}

	msg := &proto.Route{
		TargetID: self.pid.Id,
		MsgID:    msgid,
		MsgData:  msgdata,
		CallID:   m.CallID,
	}

	if m.SourcePID != nil {
		msg.SourceID = m.SourcePID.Id
	}

	SendToSession(self.pid.Address, msg)
}

func (self *socketProcess) Stop() {

}

func (self *socketProcess) CreateRPC(waitCallID int64) *util.Future {
	f := util.NewFuture()

	self.hijack = func(rpcMsg *actor.Message) bool {

		if rpcMsg.CallID == waitCallID {
			self.hijack = nil
			f.Done(rpcMsg)
			return true
		}

		return false
	}

	addHijack(self)

	return f
}

func init() {

	actor.RemoteProcessCreator = func(pid *actor.PID) actor.Process {

		return &socketProcess{
			pid: pid,
		}
	}
}
