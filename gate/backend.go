package gate

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
)

func StartBackend(clientActorCreator func() *actor.PID) {

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
