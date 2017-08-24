package actor

// Actor模式接收消息主体
type Actor interface {
	OnRecv(c Context)
}

type ActorCreator func() Actor
