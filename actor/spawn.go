package actor

func Spawn(name string, a ActorFunc) *PID {

	pid := PID{
		Address: LocalPIDManager.Address,
		Id:      name,
	}

	proc := NewLocalProcess(a, pid)

	if err := LocalPIDManager.Add(proc); err != nil {
		panic(err)
	}

	return &proc.pid

}
