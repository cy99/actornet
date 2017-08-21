package main

import (
    "github.com/davyxu/actornet/examples/chat/proto"
    "github.com/davyxu/actornet/actor"
    "github.com/davyxu/actornet/proto"
)

type user struct {
    name      string
    selfPID   *actor.PID
    clientpid *actor.PID
}

func (self *user) OnRecv(c actor.Context) {

    switch msg := c.Msg().(type) {
    case *proto.Start:
        self.name = "noname"
        self.selfPID = c.Self()
    case *chatproto.GetName:
        msg.Name = self.name
        c.Reply(msg)
    case *chatproto.ChatREQ:

        self.clientpid.Tell(&chatproto.ChatACK{
            User:    self.selfPID.ToProto(),
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

func newUser(clientpid *actor.PID) actor.Actor {
    return &user{
        clientpid: clientpid,
    }
}
