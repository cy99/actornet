package gate

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
)

type outboundClient struct {
	ses cellnet.Session
}

func (self *outboundClient) OnRecv(c actor.Context) {

	switch c.Msg().(type) {
	case *proto.Start:
	default:
		//  后台服务器发送给客户端
		self.ses.Send(c.Msg())
	}

}

func newOutboundClient(ses cellnet.Session) actor.Actor {
	return &outboundClient{
		ses: ses,
	}
}
