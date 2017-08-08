package actor

type Context interface {
	Message() interface{}
}

type localContext struct {
}

func (self *localContext) Message() interface{} {

	return nil
}
