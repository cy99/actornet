package actor

import (
	"github.com/davyxu/actornet/proto"
	"sync"
	"testing"
)

func TestHelloWorld(t *testing.T) {

	StartSystem()

	var wg sync.WaitGroup

	wg.Add(1)

	pid := Spawn("hello", func(c Context) {

		switch msg := c.Msg().(type) {
		case string:
			log.Debugln(msg)
			wg.Done()
		}

	})

	pid.NotifyData("hello")

	wg.Wait()
}

func TestEcho(t *testing.T) {

	StartSystem()

	echoFunc := func(c Context) {

		switch msg := c.Msg().(type) {
		case string:
			log.Debugln("server recv", msg)

			if c.Source() != nil {
				c.Source().NotifyDataBySender(msg, c.Self())
			}

		}

	}

	server := Spawn("server", echoFunc)

	var wg sync.WaitGroup

	wg.Add(1)

	Spawn("client", func(c Context) {

		switch data := c.Msg().(type) {
		case *proto.Start:
			log.Debugln("client start")

			server.NotifyDataBySender("hello", c.Self())
		case string:
			log.Debugln("client recv", data)

			wg.Done()
		}

	})

	wg.Wait()
}

func TestRPC(t *testing.T) {

	StartSystem()

	rpcFunc := func(c Context) {

		switch msg := c.Msg().(type) {
		case string:
			log.Debugln("server recv", msg, c.Source())

			c.Reply(msg)
		}

	}

	server := Spawn("server", rpcFunc)

	var wg sync.WaitGroup

	wg.Add(1)

	Spawn("client", func(c Context) {

		switch c.Msg().(type) {
		case *proto.Start:

			log.Debugln("client call")

			reply := server.Call("hello", c.Self())

			log.Debugln("client recv reply", reply)

			wg.Done()

		}

	})

	wg.Wait()
}
