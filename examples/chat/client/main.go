package main

import (
	"bufio"
	"github.com/davyxu/actornet/examples/chat/proto"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
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

func main() {

	queue := cellnet.NewEventQueue()

	peer := socket.NewConnector(queue).Start("127.0.0.1:8081")
	peer.SetName("client")

	cellnet.RegisterMessage(peer, "coredef.SessionConnected", func(ev *cellnet.Event) {

		ev.Send(&proto.BindClientREQ{})
	})

	cellnet.RegisterMessage(peer, "chatproto.ChatACK", func(ev *cellnet.Event) {
		msg := ev.Msg.(*chatproto.ChatACK)

		log.Infof("%s(%s) say: %s", msg.Name, msg.User.String(), msg.Content)
	})

	queue.StartLoop()

	ReadConsole(func(str string) {

		if str[0] == '/' {
			strlist := strings.Split(str, " ")

			if len(strlist) > 0 {

				switch strlist[0] {
				case "/rename":

					peer.(socket.Connector).DefaultSession().Send(&chatproto.RenameACK{
						NewName: strlist[1],
					})

					return
				}

			}
		}

		peer.(socket.Connector).DefaultSession().Send(&chatproto.ChatREQ{
			Content: str,
		})

	})
}
