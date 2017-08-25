package gate

import (
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
)

type inboundHandler struct {
}

// 收到客户端发过来的消息
func (self *inboundHandler) Call(ev *cellnet.Event) {

	if ev.Type == cellnet.Event_Recv {

		backendPID, outboundPID := GetSessionBinding(ev.Ses)

		if outboundPID != nil && backendPID != nil {

			log.Debugf("direct route: %s -> %s", outboundPID.String(), backendPID.String())

			backendPID.TellBySender(ev.Msg, outboundPID)

		} else {

			// TODO 检查消息是否在透传列表中

			switch ev.Msg.(type) {
			case *proto.BindClientREQ:

				backendAssitPID.TellBySender(&proto.BindClientREQ{ev.Ses.ID()}, receiptorPID)
			}

		}

	}
}

func newInboundHandler() cellnet.EventHandler {
	return &inboundHandler{}
}
