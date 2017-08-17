package proto

import (
	_ "github.com/davyxu/cellnet/codec/binary"
)

type Start struct {
}

type Stop struct {
}

// 路由到另外一个进程
type Route struct {
	SourceID string
	TargetID string
	MsgID    uint32
	MsgData  []byte
	CallID   int64
}

type ServiceIdentify struct {
	Domain string
}

type Serialize struct {
	Hello int32
	Ser   interface{} `obj:"-"`
}
