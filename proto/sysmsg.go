package proto

import (
	"github.com/davyxu/cellnet"
	_ "github.com/davyxu/cellnet/codec/binary"
	"github.com/davyxu/cellnet/util"
	"github.com/davyxu/goobjfmt"
	"reflect"
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
}

func (m *Route) String() string { return goobjfmt.CompactTextString(m) }

type ServiceIdentify struct {
	Address string
}

func (m *ServiceIdentify) String() string { return goobjfmt.CompactTextString(m) }

type TestMsg struct {
	Msg string
}

func (m *TestMsg) String() string { return goobjfmt.CompactTextString(m) }

func init() {

	cellnet.RegisterMessageMeta("binary", "proto.ServiceIdentify", reflect.TypeOf((*ServiceIdentify)(nil)).Elem(), util.StringHash("proto.ServiceIdentify"))
	cellnet.RegisterMessageMeta("binary", "proto.Route", reflect.TypeOf((*Route)(nil)).Elem(), util.StringHash("proto.Route"))
	cellnet.RegisterMessageMeta("binary", "proto.TestMsg", reflect.TypeOf((*TestMsg)(nil)).Elem(), util.StringHash("proto.TestMsg"))
}
