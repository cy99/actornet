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

func (self *socketProcess) Send(target *actor.PID, data interface{}) {

	msgdata, msgid, err := cellnet.EncodeMessage(data)
	if err != nil {
		log.Errorln(err)
		return
	}

	SendToSession(target.Address, &proto.Route{
		SourceID: self.pid.Id,
		TargetID: target.Id,
		MsgID:    msgid,
		MsgData:  msgdata,
	})
}

func (self *socketProcess) Stop() {

}

func init() {

	actor.RemoteProcessCreator = func() actor.Process {

		return new(socketProcess)
	}
}
