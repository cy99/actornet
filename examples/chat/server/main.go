package main

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/gate"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/golog"
)

var log *golog.Logger = golog.New("main")

func main() {

	actor.StartSystem()

	domain := actor.CreateDomain("server")

	lobbyPID := domain.Spawn(actor.NewTemplate().WithID("lobby").WithFunc(func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case *proto.BindClientREQ:

			log.Debugln("bind", c.Source())

			pid := domain.Spawn(actor.NewTemplate().WithCreator(newUser(msg.ClientSessionID)).WithParent(c.Parent()))

			c.Reply(&proto.BindClientACK{
				ClientSessionID: msg.ClientSessionID,
				ID:              pid.Id,
			})

		}

	}))

	gate.Listen("127.0.0.1:8081", lobbyPID)

	actor.LoopSystem()
}
