package nexus

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
)

// 启动本机的listen
func Listen(address string, domain string) {

	peer := socket.NewAcceptor(nil)
	peer.Start(address)

	actor.LocalPIDManager.Domain = domain

	cellnet.RegisterMessage(peer, "proto.DomainIdentifyACK", func(ev *cellnet.Event) {
		msg := ev.Msg.(*proto.DomainIdentifyACK)

		ev.Send(&proto.DomainIdentifyACK{
			Domain: domain,
		})

		addServiceSession(msg.Domain, ev.Ses)
	})

	cellnet.RegisterMessage(peer, "coredef.SessionClosed", func(ev *cellnet.Event) {

		// 其他服务器断开
		removeServiceSession(ev.Ses)

	})

	cellnet.RegisterMessage(peer, "proto.RouteACK", onRouter)

}
