package socket

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
)

// 启动本机的listen
func Listen(address string, id string) {

	peer := socket.NewAcceptor(nil)
	peer.Start(address)

	actor.LocalPIDManager.Address = id

	cellnet.RegisterMessage(peer, "proto.ServiceIdentify", func(ev *cellnet.Event) {
		msg := ev.Msg.(*proto.ServiceIdentify)

		addServiceSession(msg.Address, ev.Ses)
	})

	cellnet.RegisterMessage(peer, "coredef.SessionClosed", func(ev *cellnet.Event) {

		// 其他服务器断开
		removeServiceSession(ev.Ses)

	})

	cellnet.RegisterMessage(peer, "proto.Route", onRouter)

}
