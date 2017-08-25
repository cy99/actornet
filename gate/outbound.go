package gate

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
)

type outboundClient struct {
	actor.LocalProcess
	ses cellnet.Session
}

func (self *outboundClient) OnRecv(c actor.Context) {

	switch c.Msg().(type) {
	case *proto.Start: // 防止控制消息传到客户端
	default:
		//  后台服务器发送给客户端
		self.ses.Send(c.Msg())
	}

}

func newOutboundClient(ses cellnet.Session) func() actor.Actor {
	return func() actor.Actor {

		return &outboundClient{
			ses: ses,
		}
	}
}
