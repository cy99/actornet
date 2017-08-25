package actor

import (
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/actornet/util"
)

type PID struct {
	Domain string
	Id     string

	proc Process
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

	if self == nil {
		return nil
	}

	if self.proc != nil {
		return self.proc
	}

	dm := MustGetDomain(self.Domain)

	p := dm.GetByID(self.Id)
	if p != nil {
		self.proc = p
		return p
	}

	proc := RemoteProcessCreator(self, dm)

	if err := dm.Add(proc); err != nil {
		panic(err)
	}

	self.proc = proc

	return proc
}

func (self *PID) Tell(data interface{}) {

	self.TellBySender(data, nil)
}

func (self *PID) TellBySender(data interface{}, sender *PID) {

	if self == nil {
		return
	}

	self.ref().Tell(&Message{
		Data:      data,
		TargetPID: self,
		SourcePID: sender,
	})
}

func (self *PID) Call(data interface{}, sender *PID) interface{} {

	if self == nil {
		return nil
	}

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

	return &PID{
		Domain: domain,
		Id:     id,
	}
}

var RemoteProcessCreator func(*PID, *Domain) Process
