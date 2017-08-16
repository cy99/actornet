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

	nexus.Listen("127.0.0.1:8081", "server")

	var wg sync.WaitGroup

	wg.Add(1)

	actor.NewTemplate().WithName("server").WithFunc(func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case *proto.TestMsg:
			t.Log(msg.Msg)

			if msg.Msg == "hello" {
				wg.Done()

				if c.Source() != nil {
					t.Log("send back")
					c.Source().NotifyDataBySender(msg, c.Self())
				}

			}
		}

	}).Spawn()

	wg.Wait()
}

func TestCrossProcessNotifyClient(t *testing.T) {

	actor.StartSystem()

	nexus.Connect("127.0.0.1:8081", "client")

	time.Sleep(time.Second)

	var wg sync.WaitGroup

	wg.Add(1)

	client := actor.NewTemplate().WithName("client").WithFunc(func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case *proto.TestMsg:

			if msg.Msg == "hello" {
				wg.Done()
			}

		}

	}).Spawn()

	target := actor.NewPID("127.0.0.1:8081", "server")
	target.NotifyDataBySender(proto.TestMsg{Msg: "hello"}, client)

	wg.Wait()
}
