package actor

import (
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/actornet/util"
)

type PID struct {
	Domain string
	Id     string

	proc Process `binary:"-" text:"-"`
}

func (self *PID) IsLocal() bool {
	return LocalPIDManager.Domain == self.Domain
}

func (self *PID) ToProto() proto.PID {
	return proto.PID{
		Domain: self.Domain,
		Id:     self.Id,
	}
}

func (self *PID) FromProto(pid proto.PID) {
	self.Domain = pid.Domain
	self.Id = pid.Id
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

func (self *PID) Tell(data interface{}) {

	self.TellBySender(data, nil)
}

func (self *PID) TellBySender(data interface{}, sender *PID) {

	self.ref().Tell(&Message{
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

	self.ref().Tell(&Message{
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

	// 是本地pid时, 直接取已经存在的进程, 避免同地址pid指针不同
	if domain == LocalPIDManager.Domain {
		return NewLocalPID(id)
	}

	return &PID{
		Domain: domain,
		Id:     id,
	}
}

func NewLocalPID(id string) *PID {

	if proc := LocalPIDManager.GetByID(id); proc != nil {
		return proc.PID()
	}

	return &PID{
		Domain: LocalPIDManager.Domain,
		Id:     id,
	}
}

var RemoteProcessCreator func(*PID) Process
