package actor

import (
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/actornet/util"
)

var (
	OnReset    util.Delegate
	exitSignal chan int
)

const colorDefine = `
{
	"Rule":[
		{"Text":"panic:","Color":"Red"},
		{"Text":"[DB]","Color":"Green"},
		{"Text":"#recv","Color":"Blue"},
		{"Text":"#send","Color":"Purple"},
		{"Text":"#connected","Color":"Blue"},
		{"Text":"#listen","Color":"Blue"},
		{"Text":"#accepted","Color":"Blue"},
		{"Text":"#closed","Color":"Blue"},
		{"Text":"#notify","Color":"White"}
	]
}
`

func StartSystem() {

	//golog.SetColorDefine("*", colorDefine)
	//golog.EnableColorLogger("*", true)

	OnReset.Invoke()

	exitSignal = make(chan int)

	NewTemplate().WithID("system").WithFunc(func(c Context) {

		switch msg := c.Msg().(type) {
		case *proto.SystemExit:
			Exit(int(msg.Code))
		}

	}).Spawn()
}

// 退出
func Exit(exitcode int) {
	exitSignal <- exitcode
}

func LoopSystem() int {

	return <-exitSignal
}

var inited bool

func init() {

	OnReset.Add(func(...interface{}) error {

		inited = true

		return nil
	})

}
