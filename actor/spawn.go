package actor

func Spawn(name string, a ActorFunc) *PID {

	pid := NewLocalPID(name)

	proc := newLocalProcess(a, pid)

	if err := LocalPIDManager.Add(proc); err != nil {
		panic(err)
	}

	pid.proc = proc

	return proc.pid

}
