package socket

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
)

type socketProcess struct {
	pid     actor.PID
	callseq int64
}

func (self *socketProcess) PID() *actor.PID {
	return &self.pid
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

func (self *socketProcess) Call(m *actor.Message) *actor.Message {

	return nil
}

func init() {

	actor.RemoteProcessCreator = func(pid *actor.PID) actor.Process {

		return &socketProcess{
			pid: *pid,
		}
	}
}
