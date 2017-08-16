package actor

import "github.com/davyxu/actornet/util"

type PID struct {
	Domain string
	Id     string

	proc Process
}

func (self *PID) IsLocal() bool {
	return LocalPIDManager.Domain == self.Domain
}

func (self *PID) ref() Process {

	if self.proc != nil {
		return self.proc
	}

	if self.IsLocal() {

		p := LocalPIDManager.GetByID(self.Id)
		if p != nil {
			self.proc = p
			return p
		}

	} else if RemoteProcessCreator != nil {

		mgr := remotePIDManager(self.Domain)

		proc := mgr.GetByID(self.Id)

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

func (self *PID) Call(data interface{}, sender *PID) interface{} {

	return self.CallFuture(data, sender).Get().(*Message).Data
}

type rpcCreator interface {
	CreateRPC(waitCallID int64) *util.Future
}

func (self *PID) CallFuture(data interface{}, sender *PID) *util.Future {

	proc := sender.ref().(rpcCreator)

	callid := AllocRPCSeq()

	f := proc.CreateRPC(callid)

	self.ref().Notify(&Message{
		Data:      data,
		TargetPID: self,
		SourcePID: sender,
		CallID:    callid,
	})

	return f
}

func (self *PID) String() string {
	if self == nil {
		return "nil"
	}
	return self.Domain + "/" + self.Id
}

func NewPID(domain, id string) *PID {
	return &PID{
		Domain: domain,
		Id:     id,
	}
}

func NewLocalPID(id string) *PID {
	return &PID{
		Domain: LocalPIDManager.Domain,
		Id:     id,
	}
}

var RemoteProcessCreator func(*PID) Process
