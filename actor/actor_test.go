package actor

import "testing"

func TestHelloWorld(t *testing.T) {

	pid := SpawnByFunc(func(c Context) {

		switch data := c.(type) {
		case string:
			t.Log(data)
		}

	})

	pid.Send("hello")

}
