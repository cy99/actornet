package nexus

import (
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
)

func shareInit(peer cellnet.Peer) {

	cellnet.RegisterMessage(peer, "coredef.SessionClosed", func(ev *cellnet.Event) {

		// 其他服务器断开
		if domain := removeServiceSession(ev.Ses); domain != "" {

			broardCast(&proto.NexusClose{
				Domain: domain,
			})

		}

	})

	cellnet.RegisterMessage(peer, "proto.RouteACK", onRouter)
}
