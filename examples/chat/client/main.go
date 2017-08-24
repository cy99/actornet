package main

import (
	"bufio"
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/examples/chat/proto"
	"github.com/davyxu/actornet/nexus"
	"github.com/davyxu/golog"
	"os"
	"strings"
)

var log *golog.Logger = golog.New("main")

func ReadConsole(callback func(string)) {

	for {
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			break
		}
		text = strings.TrimRight(text, "\n\r ")

		text = strings.TrimLeft(text, " ")

		callback(text)
	}
}

type user struct {
	actor.LocalProcess
	target *actor.PID
}

func (self *user) Send(msg interface{}) {
	if self.target != nil {
		self.target.TellBySender(msg, self.PID())
	} else {
		log.Errorln("target not link")
	}
}

func (self *user) SendToLobby(msg interface{}) {

	if self.ParentPID() == nil {
		log.Errorln("lobby not link")
	} else {
		self.ParentPID().TellBySender(msg, self.PID())
	}
}

func (self *user) PublicChat(text string) {

	self.SendToLobby(&chatproto.ChatREQ{
		Content: text,
	})
}

func (self *user) Rename(newName string) {
	self.Send(&chatproto.RenameACK{
		NewName: newName,
	})
}

func (self *user) OnRecv(c actor.Context) {
	switch msg := c.Msg().(type) {
	case *chatproto.LoginACK:
		self.target = actor.NewPID(msg.User.Domain, msg.User.Id)
	case *chatproto.ChatACK:
		log.Infof("%s(%s) say: %s", msg.Name, msg.User.String(), msg.Content)
	}
}

func main() {
	actor.StartSystem()

	nexus.ConnectMulti("127.0.0.1:8081", "client")

	nexus.WaitReady("server")

	lobby := actor.NewPID("server", "lobby")

	thisUser := new(user)
	speaker := actor.NewTemplate().WithID("speaker").WithInstance(thisUser).WithParent(lobby).Spawn()

	lobby.TellBySender(&chatproto.LoginREQ{}, speaker)

	ReadConsole(func(str string) {

		if str[0] == '/' {
			strlist := strings.Split(str, " ")

			if len(strlist) > 0 {

				switch strlist[0] {
				case "/rename":
					thisUser.Rename(strlist[1])
					return
				}

			}
		}

		thisUser.PublicChat(str)

	})
}
