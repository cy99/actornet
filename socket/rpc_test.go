package socket

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"sync"
	"testing"
	"time"
)

func TestRPCServer(t *testing.T) {

	actor.StartSystem()

	Listen("127.0.0.1:8081", "server")

	var wg sync.WaitGroup

	wg.Add(1)

	actor.Spawn("server", func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case *proto.TestMsg:

			if msg.Msg == "hello" {

				log.Debugln("reply to client")
				c.Reply(msg)
				wg.Done()
			}
		}

	})

	wg.Wait()
}

func TestRPCClient(t *testing.T) {

	actor.StartSystem()

	Connect("127.0.0.1:8081", "client")

	// 等待客户端连接上服务器
	time.Sleep(time.Second)

	var wg sync.WaitGroup

	wg.Add(1)

	client := actor.Spawn("client", func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case *proto.TestMsg:

			if msg.Msg == "hello" {
				wg.Done()
			}

		}

	})

	target := actor.NewPID("127.0.0.1:8081", "server")
	reply := target.Call(proto.TestMsg{Msg: "hello"}, client)

	if msg := reply.(*proto.TestMsg).Msg; msg == "hello" {
		log.Debugln("recved reply", msg)
		wg.Done()
	}

	wg.Wait()
}
