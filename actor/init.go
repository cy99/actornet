package actor

import (
	"flag"
	"github.com/davyxu/actornet/util"
	"github.com/davyxu/golog"
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
		{"Text":"#notify","Color":"White"},
		{"Text":"#spawn","Color":"DarkGreen"}
	]
}
`

func StartSystem() {

	// 非测试环境时, 打开加色
	if flag.Lookup("test.v") == nil {
		golog.SetColorDefine("*", colorDefine)
		golog.EnableColorLogger("*", true)
	}

	OnReset.Invoke()

	exitSignal = make(chan int)

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
	inited = true

}
