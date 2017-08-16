package gate

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
)

type outboundClient struct {
}

func (self *outboundClient) OnRecv(c actor.Context) {

	switch c.Msg().(type) {
	case *proto.Start:

	default:
		BroardcastToClients(c.Msg())
	}
}

func newOutboundClient() actor.Actor {
	return &outboundClient{}
}
