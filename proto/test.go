package proto

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/util"
	"github.com/davyxu/goobjfmt"
	"reflect"
)

type TestMsg struct {
	Msg string
}

func (m *TestMsg) String() string { return goobjfmt.CompactTextString(m) }

// 客户端请求后台服务器绑定
type BindClientREQ struct {
	ClientSessionID int64 // 网关上的id
}

func (m *BindClientREQ) String() string { return goobjfmt.CompactTextString(m) }

type BindClientACK struct {
	ClientSessionID int64 // 网关上的id
	ID              string
}

func (m *BindClientACK) String() string { return goobjfmt.CompactTextString(m) }

func init() {

	cellnet.RegisterMessageMeta("binary", "proto.TestMsg", reflect.TypeOf((*TestMsg)(nil)).Elem(), util.StringHash("proto.TestMsg"))
	cellnet.RegisterMessageMeta("binary", "proto.BindClientREQ", reflect.TypeOf((*BindClientREQ)(nil)).Elem(), util.StringHash("proto.BindClientREQ"))
	cellnet.RegisterMessageMeta("binary", "proto.BindClientACK", reflect.TypeOf((*BindClientACK)(nil)).Elem(), util.StringHash("proto.BindClientACK"))
}
