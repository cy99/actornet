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

func init() {

	cellnet.RegisterMessageMeta("binary", "proto.TestMsg", reflect.TypeOf((*TestMsg)(nil)).Elem(), util.StringHash("proto.TestMsg"))
}
