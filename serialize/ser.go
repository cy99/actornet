package serialize

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
)

type Serializer interface {
	IsLoading() bool

	Serialize(data interface{})
}

func Save(pid *actor.PID) {

	bc := &struct {
		actor.Message
		Serializer
	}{
		Message: actor.Message{
			TargetPID: pid,
			Data:      &proto.Serialize{},
		},

		Serializer: NewBinaryWriter(),
	}

	pid.Notify(bc)
}
