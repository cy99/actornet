package gate

import (
	"fmt"
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
)

var receiptor *actor.PID

func Listen(address string) {

	peer := socket.NewAcceptor(nil)

	// 添加客户端消息侦听
	peer.AddChainRecv(
		cellnet.NewHandlerChain(
			newInboundHandler(),
		),
	)

	// 客户端连接
	cellnet.RegisterMessage(peer, "coredef.SessionAccepted", func(ev *cellnet.Event) {

		name := fmt.Sprintf("sid:%d", ev.Ses.ID())

		clientPID := actor.NewTemplate().WithID(name).WithInstance(newOutboundClient(ev.Ses)).Spawn()

		addClient(clientPID, ev.Ses)
	})

	// 客户端断开
	cellnet.RegisterMessage(peer, "coredef.SessionClosed", func(ev *cellnet.Event) {

		pid := removeClient(ev.Ses)
		if pid != nil {
			actor.LocalPIDManager.Remove(pid)
		}
	})

	peer.Start(address)

	receiptor = actor.NewTemplate().WithID("gate_receiptor").WithFunc(func(c actor.Context) {
		switch msg := c.Msg().(type) {
		case *proto.BindClientACK:

			clientSes := peer.GetSession(msg.ClientSessionID)
			if clientSes != nil {

				log.Debugln("bind client, sesid: %d --> pid: %s", msg.ClientSessionID, c.Source())

				pid := actor.NewPID(c.Source().Domain, msg.ID)

				clientSes.SetTag(pid)
			} else {
				log.Warnln("BindClinet: client session not found: ", msg.ClientSessionID)
			}

		}
	}).Spawn()
}
