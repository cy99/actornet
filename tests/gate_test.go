package tests

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/gate"
	"github.com/davyxu/actornet/nexus"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
	"sync"
	"testing"
	"time"
)

func TestLinkBackend(t *testing.T) {

	actor.EnableDebug = true

	actor.StartSystem()

	nexus.Connect("127.0.0.1:7111", "server")

	var wg sync.WaitGroup

	wg.Add(1)

	onRouteMsg := func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case *proto.TestMsg:

			log.Debugln("server recv", msg, c.Source())

			if msg.Msg == "hello" {
				wg.Done()

				if c.Source() != nil {
					log.Debugln("send back")

					c.Reply(msg)
				}

			}
		}

	}

	gate.StartBackend(func() *actor.PID {

		return actor.NewTemplate().WithFunc(onRouteMsg).Spawn()

	})

	wg.Wait()

	time.Sleep(time.Second)
}

func TestLinkGate(t *testing.T) {

	actor.EnableDebug = true

	actor.StartSystem()

	nexus.Listen("127.0.0.1:7111", "gate")

	gate.Listen("127.0.0.1:8031")

	actor.LoopSystem()
}

func TestLinkClient(t *testing.T) {

	peer := socket.NewConnector(nil)

	peer.Start("127.0.0.1:8031")

	var wg sync.WaitGroup
	wg.Add(1)

	// 客户端连接
	cellnet.RegisterMessage(peer, "coredef.SessionConnected", func(ev *cellnet.Event) {

		ev.Send(&proto.BindClientREQ{})

		time.Sleep(time.Second)

		ev.Send(&proto.TestMsg{"hello"})
	})

	cellnet.RegisterMessage(peer, "proto.TestMsg", func(ev *cellnet.Event) {

		msg := ev.Msg.(*proto.TestMsg)

		if msg.Msg == "hello" {
			wg.Done()
		}

	})

	wg.Wait()
}
