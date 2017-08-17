package nexus

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/actornet/util"
	"github.com/davyxu/cellnet"
)

type nexusProcess struct {
	pid *actor.PID

	hijack func(*actor.Message) bool
}

func (self *nexusProcess) PID() *actor.PID {
	return self.pid
}

func (self *nexusProcess) Notify(data interface{}) {

	m := data.(*actor.Message)

	if actor.EnableDebug {
		log.Debugf("#notify %s", data.(actor.Context).String())
	}

	msgdata, msgid, err := cellnet.EncodeMessage(m.Data)
	if err != nil {
		log.Errorln(err)
		return
	}

	msg := &proto.RouteACK{
		TargetID: self.pid.Id,
		MsgID:    msgid,
		MsgData:  msgdata,
		CallID:   m.CallID,
	}

	if m.SourcePID != nil {
		msg.SourceID = m.SourcePID.Id
	}

	sendToDomain(self.pid.Domain, msg)
}

func (self *nexusProcess) Stop() {

}

func (self *nexusProcess) CreateRPC(waitCallID int64) *util.Future {
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

		return &nexusProcess{
			pid: pid,
		}
	}
}
