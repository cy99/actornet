package tests

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/nexus"
	"github.com/davyxu/actornet/proto"
	"sync"
	"testing"
	"time"
)

func TestCrossProcessCallServer(t *testing.T) {

	actor.StartSystem()

	domain := actor.CreateDomain("server")

	nexus.Listen("127.0.0.1:8081", "server")

	var wg sync.WaitGroup

	wg.Add(1)

	domain.Spawn(actor.NewTemplate().WithID("echo").WithFunc(func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case *proto.TestMsgACK:

			if msg.Msg == "hello" {

				log.Debugln("reply to client")
				c.Reply(msg)
				wg.Done()
			}
		}

	}))

	wg.Wait()

	// 等待发送完毕
	time.Sleep(time.Second)
}

func TestCrossProcessCallClient(t *testing.T) {

	domain := actor.CreateDomain("client")
	actor.StartSystem()

	nexus.ConnectSingleton("127.0.0.1:8081", "client")

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
	reply := target.Call(proto.TestMsgACK{Msg: "hello"}, client)

	if msg := reply.(*proto.TestMsgACK).Msg; msg == "hello" {
		log.Debugln("recved reply", msg)
		wg.Done()
	}

	wg.Wait()
}
