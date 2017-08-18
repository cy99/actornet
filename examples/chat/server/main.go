package main

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/examples/chat/proto"
	"github.com/davyxu/actornet/nexus"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/golog"
)

var log *golog.Logger = golog.New("main")

type lobby struct {
	users   []*actor.PID
	selfPID *actor.PID
}

func (self *lobby) Broardcast(data interface{}) {
	for _, u := range self.users {

		u.TellBySender(data, self.selfPID)
	}
}

func (self *lobby) OnRecv(c actor.Context) {
	switch msg := c.Msg().(type) {
	case *proto.Start:
		self.selfPID = c.Self()
	case *chatproto.LoginREQ:

		// 生成服务器对象pid
		serverUserPID := actor.NewTemplate().WithInstance(newUser(c.Source())).Spawn()

		self.users = append(self.users, serverUserPID)

		self.Broardcast(&chatproto.LoginACK{User: serverUserPID.ToProto()})

	case *chatproto.ChatREQ:
		self.Broardcast(msg)
	}
}

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
	case *chatproto.ChatREQ:

		self.clientpid.TellBySender(&chatproto.ChatACK{
			Name:    self.name,
			Content: msg.Content,
		}, c.Self())
	case *chatproto.LoginACK:
		self.clientpid.TellBySender(msg, c.Self())
	case *chatproto.RenameACK:

		log.Debugln("%s rename '%s' -> '%s'", c.Self().String(), self.name, msg.NewName)

		self.name = msg.NewName

		// TODO 广播给大厅所有人
	}
}

func newUser(clientpid *actor.PID) actor.Actor {
	return &user{
		clientpid: clientpid,
	}
}

func main() {

	actor.StartSystem()

	nexus.Listen("127.0.0.1:8081", "server")

	actor.NewTemplate().WithID("lobby").WithInstance(new(lobby)).Spawn()

	actor.LoopSystem()
}
