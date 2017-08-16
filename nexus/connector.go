package nexus

import (
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
)

// 启动本机的listen
func Connect(address string, id string) {

	peer := socket.NewConnector(nil)
	peer.Start(address)

	cellnet.RegisterMessage(peer, "coredef.SessionConnected", func(ev *cellnet.Event) {

		ev.Send(&proto.ServiceIdentify{
			Address: id,
		})

		addServiceSession(address, ev.Ses)
	})

	cellnet.RegisterMessage(peer, "coredef.SessionClosed", func(ev *cellnet.Event) {

		removeServiceSession(ev.Ses)

	})

	cellnet.RegisterMessage(peer, "proto.Route", onRouter)
}
