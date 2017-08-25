package main

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/golog"
)

var log *golog.Logger = golog.New("main")

func main() {

	actor.StartSystem()

	domain := actor.CreateDomain("local")

	pid := domain.Spawn(actor.NewTemplate().WithID("hello").WithFunc(func(c actor.Context) {

		switch msg := c.Msg().(type) {
		case string:
			log.Debugln(msg)

			actor.Exit(0)
		}

	}))

	pid.Tell("hello world")

	actor.LoopSystem()

}
