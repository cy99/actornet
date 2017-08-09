package actor

type PID struct {
	Address string
	Id      string

	proc Process
}

func (self *PID) raw() PID {

	return PID{
		Address: self.Address,
		Id:      self.Id,
	}
}

func (self *PID) Send(target *PID, data interface{}) {

	if target != nil {
		target.proc.Send(self, data)
	} else {
		panic("empty target")
	}

}

func (self *PID) String() string {
	if self == nil {
		return "nil"
	}
	return self.Address + "/" + self.Id
}

func NewPID(address, id string) *PID {
	return &PID{
		Address: address,
		Id:      id,
	}
}

func NewLocalPID(id string) *PID {
	return &PID{
		Address: localPIDManager.Address,
		Id:      id,
	}
}

var Root = NewLocalPID("Root")
