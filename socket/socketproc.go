package socket

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
)

type socketProcess struct {
	pid *actor.PID

	hijack func(*actor.Message) bool
}

func (self *socketProcess) PID() *actor.PID {
	return self.pid
}

func (self *socketProcess) Notify(m *actor.Message) {

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

func (self *socketProcess) BeginHijack(waitCallID int64) chan *actor.Message {
	reply := make(chan *actor.Message)

	self.hijack = func(rpcMsg *actor.Message) bool {

		if rpcMsg.CallID == waitCallID {
			reply <- rpcMsg
			return true
		}

		return false
	}

	addHijack(self)

	return reply
}

func (self *socketProcess) EndHijack(reply chan *actor.Message) *actor.Message {

	msgReply := <-reply

	self.hijack = nil

	return msgReply
}


func init() {

	actor.RemoteProcessCreator = func(pid *actor.PID) actor.Process {

		return &socketProcess{
			pid: pid,
		}
	}
}
