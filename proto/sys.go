package proto

import (
	_ "github.com/davyxu/cellnet/codec/binary"
)

type PID struct {
	Domain string
	Id     string
}

// 一个Actor启动时
type Start struct {
}

// 一个Actor停止时
type Stop struct {
}

// 整个物理进程退出
type SystemExit struct {
	Code int32
}

// 路由到另外一个进程
type RouteACK struct {
	SourceID string
	TargetID string
	MsgID    uint32
	MsgData  []byte
	CallID   int64
}

// 领域标识
type DomainIdentifyACK struct {
	Domain string
}
