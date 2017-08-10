package socket

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
)

type socketProcess struct {
	pid actor.PID
}

func (self *socketProcess) PID() *actor.PID {
	return &self.pid
}

func (self *socketProcess) Notify(data interface{}, sender *actor.PID) {

	msgdata, msgid, err := cellnet.EncodeMessage(data)
	if err != nil {
		log.Errorln(err)
		return
	}

	msg := &proto.Route{
		TargetID: self.pid.Id,
		MsgID:    msgid,
		MsgData:  msgdata,
	}

	if sender != nil {
		msg.SourceID = sender.Id
	}

	SendToSession(self.pid.Address, msg)
}

func (self *socketProcess) Stop() {

}

func init() {

	actor.RemoteProcessCreator = func(pid *actor.PID) actor.Process {

		return &socketProcess{
			pid: *pid,
		}
	}
}
