package actor



type Actor interface {
	Receive(c Context)
}

// The ActorFunc type is an adapter to allow the use of ordinary functions as actors to process messages
type ActorFunc func(c Context)

// Receive calls f(c)
func (f ActorFunc) Receive(c Context) {
	f(c)
}
