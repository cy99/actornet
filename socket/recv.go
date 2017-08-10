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

	address, _ := AddressBySession(ev.Ses)

	tgtProc := actor.LocalPIDManager.Get(msg.TargetID)

	if tgtProc != nil {

		if msg.SourceID != "" {
			tgtProc.Notify(userMsg, actor.NewPID(address, msg.SourceID))
		} else {
			tgtProc.Notify(userMsg, nil)
		}

	} else {
		log.Errorln("node not found:", msg.TargetID)
	}

}
