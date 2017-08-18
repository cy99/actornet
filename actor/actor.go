package actor

// Actor模式接收消息主体
type Actor interface {
	OnRecv(c Context)
}

type ActorFunc func(c Context)

func (f ActorFunc) OnRecv(c Context) {
	f(c)
}
