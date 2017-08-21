package nexus

import "github.com/davyxu/actornet/actor"

var (
	watchList []*actor.PID
)

func Watch(pid *actor.PID) {

	watchList = append(watchList, pid)
}

func broardCast(data interface{}) {

	for _, pid := range watchList {
		pid.Tell(data)
	}
}
