package gate

import (
	"fmt"
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
)

var (
	receiptorPID *actor.PID

	acceptorPeer cellnet.Peer

	// 后台的辅助actor
	backendAssitPID *actor.PID

	gateDomain *actor.Domain
)

func Listen(address string, backendAssit *actor.PID) {

	gateDomain = actor.CreateDomain("gate")

	backendAssitPID = backendAssit

	acceptorPeer = socket.NewAcceptor(nil)

	// 添加客户端消息侦听
	acceptorPeer.AddChainRecv(
		cellnet.NewHandlerChain(
			newInboundHandler(),
		),
	)

	// 客户端断开
	cellnet.RegisterMessage(acceptorPeer, "coredef.SessionClosed", func(ev *cellnet.Event) {

		pid := removeClient(ev.Ses)
		if pid != nil {
			gateDomain.Kill(pid)
		}
	})

	acceptorPeer.Start(address)

	// 前台接收 后面的服务器发过来的消息
	receiptorPID = gateDomain.Spawn(actor.NewTemplate().WithID("gate_receiptor").WithFunc(func(c actor.Context) {
		switch msg := c.Msg().(type) {
		case *proto.BindClientACK:

			clientSes := acceptorPeer.GetSession(msg.ClientSessionID)
			if clientSes != nil {

				backendPID := actor.NewPID(c.Source().Domain, msg.ID)

				outboundID := MakeOutboundID(clientSes.ID())

				outboundPID := gateDomain.Spawn(actor.NewTemplate().WithID(outboundID).WithCreator(newOutboundClient(clientSes)))

				addClient(outboundPID, backendPID, clientSes)

				// 回应客户端
				clientSes.Send(&proto.BindClientACK{})

			} else {
				log.Warnln("BindClinet: client session not found: ", msg.ClientSessionID)
			}

		}
	}))
}

func MakeOutboundID(clientSessionID int64) string {
	return fmt.Sprintf("sid:%d", clientSessionID)
}
