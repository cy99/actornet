package main

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/examples/chat/proto"
	"github.com/davyxu/actornet/gate"
	"github.com/davyxu/actornet/proto"
)

type user struct {
	actor.LocalProcess
	name      string
	clientpid *actor.PID
}

func (self *user) OnRecv(c actor.Context) {

	switch msg := c.Msg().(type) {
	case *proto.Start:
		self.name = "noname"
	case *chatproto.GetName:
		msg.Name = self.name
		c.Reply(msg)
	case *chatproto.ChatREQ:

		log.Debugln("chat", c.Source())

		self.clientpid.Tell(&chatproto.ChatACK{
			User:    self.PID().ToProto(),
			Name:    self.name,
			Content: msg.Content,
		})

	case *chatproto.RenameACK:

		log.Debugf("[%s] rename '%s' -> '%s'", c.Self().String(), self.name, msg.NewName)

		self.name = msg.NewName

	case *chatproto.ChatACK, *chatproto.LoginACK:
		self.clientpid.Tell(msg)
	}
}

func newUser(clientSesID int64) actor.ActorCreator {
	return func() actor.Actor {
		return &user{
			clientpid: actor.NewPID("gate", gate.MakeOutboundID(clientSesID)),
		}
	}

}
