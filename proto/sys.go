package proto

import (
	"github.com/davyxu/goobjfmt"
)

type Start struct {
}

func (m *Start) String() string { return goobjfmt.CompactTextString(m) }

type Stop struct {
}

func (m *Stop) String() string { return goobjfmt.CompactTextString(m) }

type Serialize struct {
	Hello int32
	Ser   interface{} `obj:"-"`
}

func (m *Serialize) String() string { return goobjfmt.CompactTextString(m) }
