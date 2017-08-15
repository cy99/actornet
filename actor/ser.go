package actor

import (
	"bytes"
)

type Serializer interface {
	IsLoading() bool

	Serialize(data interface{})

	Bytes() []byte
}

func Save(pid *PID) []byte {

	proc := LocalPIDManager.Get(pid)

	ser := NewBinaryWriter()

	proc.(interface {
		Serialize(ser Serializer)
	}).Serialize(ser)

	return ser.Bytes()
}

func Load(pid *PID, data []byte) {

	proc := LocalPIDManager.Get(pid)

	ser := NewBinaryReader(bytes.NewReader(data))

	proc.(interface {
		Serialize(ser Serializer)
	}).Serialize(ser)

}
