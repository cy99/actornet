package actor

import (
	"sync"
	"testing"
)

func TestHelloWorld(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	pid := SpawnFromFunc(func(c Context) {

		switch data := c.Msg().(type) {
		case string:
			t.Log(data)
			wg.Done()
		}

	})

	pid.Send("hello")

	wg.Wait()
}
