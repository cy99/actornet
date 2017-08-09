package actor

import (
	"github.com/davyxu/actornet/proto"
	"sync"
	"testing"
)

func TestHelloWorld(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	pid := Spawn("hello", func(c Context) {

		switch msg := c.Msg().(type) {
		case string:
			t.Log(msg)
			wg.Done()
		}

	})

	Root.Send(pid, "hello")

	wg.Wait()
}

func Test2ActorCommunicate(t *testing.T) {

	echoFunc := func(c Context) {

		switch msg := c.Msg().(type) {
		case string:
			t.Log("server recv", msg)
			c.Self().Send(c.Source(), msg)
		}

	}

	server := Spawn("server", echoFunc)

	var wg sync.WaitGroup

	wg.Add(1)

	Spawn("client", func(c Context) {

		switch data := c.Msg().(type) {
		case *proto.Start:
			t.Log("client start", c.Self().String())
			c.Self().Send(server, "hello")
		case string:
			t.Log("client recv", data)

			wg.Done()
		}

	})

	wg.Wait()
}
