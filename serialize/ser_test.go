package serialize

import (
	"github.com/davyxu/goobjfmt"
	"testing"
)

type foo struct {
	key interface{}
}

func TestSer(t *testing.T) {

	data, err := goobjfmt.BinaryWrite(foo{key: "hello"})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(data)

	var my foo
	err = goobjfmt.BinaryRead(data, &my)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(my)

}
