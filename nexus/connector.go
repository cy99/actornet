package nexus

import (
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
)

func ConnectSingleton(address string, domain string) {

	con(address, domain, true)
}

func ConnectMulti(address string, domain string) {

	con(address, domain, false)
}

// 启动本机的listen
func con(address string, domain string, singleton bool) {

	peer := socket.NewConnector(nil)
	peer.Start(address)

	cellnet.RegisterMessage(peer, "coredef.SessionConnected", func(ev *cellnet.Event) {

		ev.Send(&proto.DomainIdentifyACK{
			Domain:    domain,
			Singleton: singleton,
		})
	})

	cellnet.RegisterMessage(peer, "proto.DomainIdentifyACK", func(ev *cellnet.Event) {
		msg := ev.Msg.(*proto.DomainIdentifyACK)

		broardCast(&proto.NexusOpen{
			Domain: msg.Domain,
		})

		addServiceSession(msg.Domain, ev.Ses)
	})

	shareInit(peer)
}
