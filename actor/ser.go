package actor

import (
	"github.com/davyxu/actornet/proto"
)

type Serializer interface {
	IsLoading() bool

	Serialize(data interface{})

	Bytes() []byte
}

func Save(pid *PID) []byte {

	saverPID := SpawnByFunc("saver", func(ctx Context) {

	})

	ser := NewBinaryWriter()

	pid.Call(&proto.Serialize{
		Ser: ser,
	}, saverPID)

	return ser.Bytes()
}
