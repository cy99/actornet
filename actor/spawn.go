package actor

func SpawnFromFunc(name string, f ActorFunc) *PID {

	pid := PID{
		Address: localPIDManager.Address,
		Id:      name,
	}

	proc := NewLocalProcess(f, pid)

	if err := localPIDManager.Add(proc); err != nil {
		panic(err)
	}

	return &proc.pid

}
