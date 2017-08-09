package actor

type Actor interface {
	OnRecv(c Context)
}

// The ActorFunc type is an adapter to allow the use of ordinary functions as actors to process messages
type ActorFunc func(c Context)

// Receive calls f(c)
func (f ActorFunc) OnRecv(c Context) {
	f(c)
}
