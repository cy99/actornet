package socket

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
)

func onRouter(ev *cellnet.Event) {

	msg := ev.Msg.(*proto.Route)

	userMsg, err := cellnet.DecodeMessage(msg.MsgID, msg.MsgData)
	if err != nil {
		log.Errorln(err)
		return
	}

	sourceProc := actor.LocalPIDManager.Get(msg.SourceID)

	sourceProc.Send(actor.NewLocalPID(msg.TargetID), userMsg)

}
