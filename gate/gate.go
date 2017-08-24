package gate

import (
	"fmt"
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
)

var receiptor *actor.PID

var acceptor cellnet.Peer

func Listen(address string) {

	acceptor = socket.NewAcceptor(nil)

	// 添加客户端消息侦听
	acceptor.AddChainRecv(
		cellnet.NewHandlerChain(
			newInboundHandler(),
		),
	)

	// 客户端断开
	cellnet.RegisterMessage(acceptor, "coredef.SessionClosed", func(ev *cellnet.Event) {

		pid := removeClient(ev.Ses)
		if pid != nil {
			actor.LocalPIDManager.Remove(pid)
		}
	})

	acceptor.Start(address)

	receiptor = actor.NewTemplate().WithID("gate_receiptor").WithFunc(func(c actor.Context) {
		switch msg := c.Msg().(type) {
		case *proto.BindClientACK:

			clientSes := acceptor.GetSession(msg.ClientSessionID)
			if clientSes != nil {

				backendPID := actor.NewPID(c.Source().Domain, msg.ID)

				outboundName := fmt.Sprintf("sid:%d", clientSes.ID())

				outboundPID := actor.NewTemplate().WithID(outboundName).WithCreator(newOutboundClient(clientSes)).Spawn()

				addClient(outboundPID, backendPID, clientSes)

				// 回应客户端
				clientSes.Send(&proto.BindClientACK{})

			} else {
				log.Warnln("BindClinet: client session not found: ", msg.ClientSessionID)
			}

		}
	}).Spawn()
}
