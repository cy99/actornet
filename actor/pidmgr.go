package actor

type PIDManager struct {
	seq int64
}

func (self *PIDManager) AllocID() int64 {

	self.seq++
	return self.seq
}

func NewPIDManager() *PIDManager {
	return &PIDManager{}
}


var localPIDManager = NewPIDManager()