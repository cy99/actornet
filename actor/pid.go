package actor

type PID struct {
	Address string
	Id      string

	proc Process
}

func (self *PID) IsLocal() bool {
	return LocalPIDManager.Address == self.Address
}

func (self *PID) ref() Process {

	if self.proc != nil {
		return self.proc
	}

	if self.IsLocal() {

		p := LocalPIDManager.Get(self.Id)
		if p != nil {
			self.proc = p
			return p
		}

	} else if RemoteProcessCreator != nil {

		mgr := remotePIDManager(self.Address)

		proc := mgr.Get(self.Id)

		if proc == nil {
			proc = RemoteProcessCreator(self)

			if err := mgr.Add(proc); err != nil {
				panic(err)
			}
		}

		self.proc = proc

		return proc
	}

	panic("invalid pid to create process")

	return nil
}

func (self *PID) Notify(data interface{}) {

	self.ref().Notify(data)
}

func (self *PID) NotifyData(data interface{}) {

	self.ref().Notify(&Message{
		Data:      data,
		TargetPID: self,
	})
}

func (self *PID) NotifyDataBySender(data interface{}, sender *PID) {

	self.ref().Notify(&Message{
		Data:      data,
		TargetPID: self,
		SourcePID: sender,
	})
}

type CallHijack interface {
	BeginHijack(waitCallID int64) chan *Message

	EndHijack(reply chan *Message) *Message
}

func (self *PID) Call(data interface{}, sender *PID) interface{} {

	proc := sender.ref().(CallHijack)

	callid := AllocRPCSeq()

	reply := proc.BeginHijack(callid)

	self.Notify(&Message{
		Data:      data,
		TargetPID: self,
		SourcePID: sender,
		CallID:    callid,
	})

	return proc.EndHijack(reply).Data
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
		Address: LocalPIDManager.Address,
		Id:      id,
	}
}

var RemoteProcessCreator func(*PID) Process
