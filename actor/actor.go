package actor

// Actor模式接收消息主体
type Actor interface {
	OnRecv(c Context)
}

type ActorCreator func() Actor

type funcActor struct {
	LocalProcess
	f func(c Context)
}

func (self *funcActor) OnRecv(c Context) {
	self.f(c)
}

func newFuncActor(f func(c Context)) func() Actor {
	return func() Actor {
		return &funcActor{
			f: f,
		}
	}

}
