package actor

import (
	"github.com/davyxu/actornet/proto"
	"sync"
	"testing"
)

func TestHelloWorld(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	pid := SpawnFromFunc("hello", func(c Context) {

		switch data := c.Msg().(type) {
		case string:
			t.Log(data)
			wg.Done()
		}

	})

	Root.Send(pid, "hello")

	wg.Wait()
}

func Test2ActorCommunicate(t *testing.T) {

	echoFunc := func(c Context) {

		switch data := c.Msg().(type) {
		case string:
			t.Log("server recv", data)
			c.Self().Send(c.Source(), data)
		}

	}

	server := SpawnFromFunc("server", echoFunc)

	var wg sync.WaitGroup

	wg.Add(1)

	SpawnFromFunc("client", func(c Context) {

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
