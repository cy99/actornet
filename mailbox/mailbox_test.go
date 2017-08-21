package mailbox

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type testMailreceiver struct {
	wg    *sync.WaitGroup
	count int
	max   int
}

func (self *testMailreceiver) OnRecv(data interface{}) {
	self.count++

	if self.count == self.max {
		self.wg.Done()
	} else if self.count > self.max {
		log.Errorln("Unexpected data")
	}

}

func TestUnbounded(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(1)
	const max = 1000000
	c := 100

	mailbox := NewUnbouned()
	mailbox.Start(&testMailreceiver{
		wg:  &wg,
		max: max,
	})

	for j := 0; j < c; j++ {
		cmax := max / c
		go func() {
			for i := 0; i < cmax; i++ {
				if rand.Intn(10) == 0 {
					time.Sleep(time.Duration(rand.Intn(1000)))
				}

				mailbox.Post(fmt.Sprintf("%v %v", j, i))
			}
		}()
	}
	wg.Wait()
	time.Sleep(1 * time.Second)
}
