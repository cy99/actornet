package actor

type Relation interface {
	AddChild(*PID)

	ParentPID() *PID

	SetParentPID(ppid *PID)
}

type RelationImplement struct {
	ppid *PID // 父级

	childs []*PID // 子级

	proc Process
}

func (self *RelationImplement) ParentPID() *PID {
	return self.ppid
}

func (self *RelationImplement) SetParentPID(ppid *PID) {
	self.ppid = ppid
}

func (self *RelationImplement) AddChild(pid *PID) {

	childProc := pid.ref()

	if childProc == nil {
		panic("child can not be nil when add child")
	}

	childProc.SetParentPID(self.proc.PID())

	self.childs = append(self.childs, pid)
}

func NewRelation(proc Process) *RelationImplement {
	return &RelationImplement{
		proc: proc,
	}
}
