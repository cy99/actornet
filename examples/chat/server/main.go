package main

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/nexus"
	"github.com/davyxu/golog"
)

var log *golog.Logger = golog.New("main")

func main() {

	actor.StartSystem()

	nexus.Listen("127.0.0.1:8081", "server")

	actor.NewTemplate().WithID("lobby").WithCreator(newLobby).Spawn()

	actor.LoopSystem()
}
