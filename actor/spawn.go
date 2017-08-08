package actor

func SpawnFromFunc(f ActorFunc) *PID {

	proc := NewLocalProcess(f)

	pid := &PID{
		ID:   localPIDManager.AllocID(),
		proc: proc,
	}

	proc.Stop()

	return pid

}
