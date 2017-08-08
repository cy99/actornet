package actor

type Context interface {
	Msg() interface{}
}
