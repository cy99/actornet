package socket

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"sync"
	"testing"
	"time"
)

func TestServer(t *testing.T) {

	Listen("127.0.0.1:8081", "server")

	var wg sync.WaitGroup

	wg.Add(1)

	actor.Spawn("server", func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case *proto.TestMsg:
			t.Log(msg.Msg)

			if msg.Msg == "hello" {
				wg.Done()

				if c.Source() != nil {
					t.Log("send back")
					c.Source().Notify(msg, c.Self())
				}

			}
		}

	})

	wg.Wait()
}

func TestClient(t *testing.T) {

	Connect("127.0.0.1:8081", "client")

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
	target.Notify(proto.TestMsg{Msg: "hello"}, client)

	wg.Wait()
}
