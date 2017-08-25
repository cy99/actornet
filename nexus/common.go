package nexus

import (
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
)

func shareInit(peer cellnet.Peer) {

	cellnet.RegisterMessage(peer, "coredef.SessionClosed", func(ev *cellnet.Event) {

		removeDomains(ev.Ses)
	})

	cellnet.RegisterMessage(peer, "proto.DomainSyncACK", func(ev *cellnet.Event) {
		msg := ev.Msg.(*proto.DomainSyncACK)

		addDomains(msg.DomainNames, ev.Ses)
	})

	cellnet.RegisterMessage(peer, "proto.RouteACK", onRouter)
}
