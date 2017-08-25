package tests

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/nexus"
	"github.com/davyxu/actornet/proto"
	"sync"
	"testing"
	"time"
)

func TestCrossProcessNotifyServer(t *testing.T) {

	actor.StartSystem()

	domain := actor.CreateDomain("server")

	nexus.Listen("127.0.0.1:8081", "server")

	var wg sync.WaitGroup

	wg.Add(1)

	domain.Spawn(actor.NewTemplate().WithID("echo").WithFunc(func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case *proto.TestMsgACK:
			t.Log(msg.Msg)

			if msg.Msg == "hello" {
				wg.Done()

				if c.Source() != nil {
					t.Log("send back")
					c.Reply(msg)
				}

			}
		}

	}))

	wg.Wait()

	// 等待发送完毕
	time.Sleep(time.Second)
}

func TestCrossProcessNotifyClient(t *testing.T) {

	actor.StartSystem()

	domain := actor.CreateDomain("client")

	nexus.ConnectSingleton("127.0.0.1:8081", "client")

	// 等待客户端连接上服务器
	nexus.WaitReady("server")

	var wg sync.WaitGroup

	wg.Add(1)

	client := domain.Spawn(actor.NewTemplate().WithID("client").WithFunc(func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case *proto.TestMsgACK:

			if msg.Msg == "hello" {
				wg.Done()
			}

		}

	}))

	target := actor.NewPID("server", "echo")
	target.TellBySender(proto.TestMsgACK{Msg: "hello"}, client)

	wg.Wait()
}
