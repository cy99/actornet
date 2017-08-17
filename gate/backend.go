package gate

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
)

func StartBackend(clientActorCreator func() *actor.PID) {

	// 网关助手, 放置于每个后台服务器
	// 未指定目标的消息,可以汇总到这里
	actor.NewTemplate().WithID("gate_assit").WithFunc(func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case *proto.BindClientREQ:

			pid := clientActorCreator()

			c.Reply(&proto.BindClientACK{
				ClientSessionID: msg.ClientSessionID,
				ID:              pid.Id,
			})
		}

	}).Spawn()
}
