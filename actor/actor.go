package actor

// Actor模式接收消息主体
type Actor interface {
	OnRecv(c Context)
}

// 信息上下文
type Context interface {

	// 消息本体
	Msg() interface{}

	// 消息来源方，可能为空
	Source() *PID

	// Actor本体的PID
	Self() *PID

	// 当对方用Call调用时， 需要用Reply回应
	Reply(data interface{})
}

type ActorFunc func(c Context)

func (f ActorFunc) OnRecv(c Context) {
	f(c)
}

