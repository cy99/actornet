package actor

func SpawnByFunc(f ActorFunc) *PID {

	return &PID{
		ID: localPIDManager.AllocID(),
		p:  NewLocalProcess(f),
	}

}
