package tests

import (
	"sync"
	"testing"

	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
)

func TestHelloWorld(t *testing.T) {

	actor.StartSystem()

	var wg sync.WaitGroup

	wg.Add(1)

	pid := actor.NewTemplate().WithID("hello").WithFunc(func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case string:
			log.Debugln(msg)
			wg.Done()
		}

	}).Spawn()

	pid.Notify("hello")

	wg.Wait()
}

func TestEcho(t *testing.T) {

	actor.StartSystem()

	echoFunc := func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case string:
			log.Debugln("server recv", msg)

			if c.Source() != nil {
				c.Source().NotifyBySender(msg, c.Self())
			}

		}

	}

	server := actor.NewTemplate().WithID("server").WithFunc(echoFunc).Spawn()

	var wg sync.WaitGroup

	wg.Add(1)

	actor.NewTemplate().WithID("client").WithFunc(func(c actor.Context) {

		switch data := c.Msg().(type) {
		case *proto.Start:
			log.Debugln("client start")

			server.NotifyBySender("hello", c.Self())
		case string:
			log.Debugln("client recv", data)

			wg.Done()
		}

	}).Spawn()

	wg.Wait()
}

func TestRPC(t *testing.T) {

	actor.StartSystem()

	rpcFunc := func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case string:
			log.Debugln("server recv", msg, c.Source())

			c.Reply(msg)
		}

	}

	server := actor.NewTemplate().WithID("server").WithFunc(rpcFunc).Spawn()

	var wg sync.WaitGroup

	wg.Add(1)

	actor.NewTemplate().WithID("client").WithFunc(func(c actor.Context) {

		switch c.Msg().(type) {
		case *proto.Start:

			log.Debugln("client call")

			reply := server.Call("hello", c.Self())

			log.Debugln("client recv reply", reply)

			wg.Done()

		}

	}).Spawn()

	wg.Wait()
}

type myActor struct {
	hp       int
	nameList []string
}

func (self *myActor) OnRecv(c actor.Context) {

	switch c.Msg().(type) {
	case *proto.Start:
		self.hp = 123
		self.nameList = []string{"genji", "mei"}
	}
}

func (self *myActor) OnSerialize(ser actor.Serializer) {
	ser.Serialize(&self.hp)
	ser.Serialize(&self.nameList)
}

func TestSerialize(t *testing.T) {

	actor.StartSystem()

	pid := actor.NewTemplate().WithID("hibernate").WithInstance(new(myActor)).Spawn()

	t.Log(actor.Save(pid))

}
