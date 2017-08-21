package proto

import (
	_ "github.com/davyxu/cellnet/codec/binary"
)

type PID struct {
	Domain string
	Id     string
}
// ============================================
// 系统事件
// ============================================

// 一个Actor启动时
// any -> any
type Start struct {
}

// 一个Actor停止时
// any -> any
type Stop struct {
}

// 整个物理进程退出
// any -> localhost/system
type SystemExit struct {
	Code int32
}

// 进程互联通道打开
// nexus -> any
type NexusOpen struct {
	Domain string
}

// 进程互联通道关闭
// nexus -> any
type NexusClose struct {
	Domain string
}

// ============================================
// 组件通信消息
// ============================================

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
	Domain    string
	Singleton bool
}
