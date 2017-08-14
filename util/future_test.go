package util

import (
	"testing"
	"time"
)

func TestFuture(t *testing.T) {

	f := NewFuture()

	go func() {

		t.Log("begin sleep 2 sec")
		time.Sleep(time.Second * 2)

		t.Log("send done")
		f.Done("done")

	}()

	t.Log("recv", f.Get())

}
