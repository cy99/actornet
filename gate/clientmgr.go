package gate

import (
	"bytes"
	"fmt"
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/cellnet"
)

type sesBinding struct {
	outbound *actor.PID // 下发客户端的actor
	backend  *actor.PID // 绑定的后台actor
}

func GetSessionBinding(ses cellnet.Session) (backend, outbound *actor.PID) {

	b := rawSessionBinding(ses)

	backend = b.backend
	outbound = b.outbound

	return
}

func rawSessionBinding(ses cellnet.Session) (ret *sesBinding) {
	tag := ses.Tag()

	if tag == nil {
		ret = &sesBinding{}
		ses.SetTag(ret)
	} else {
		ret = tag.(*sesBinding)
	}

	return
}

func addClient(outbound, backend *actor.PID, ses cellnet.Session) {

	bind := rawSessionBinding(ses)
	bind.outbound = outbound
	bind.backend = backend

	log.Infof("client bind, sid: %d  outbound: %s <- backend: %s", ses.ID(), outbound.String(), backend.String())
}

func removeClient(ses cellnet.Session) *actor.PID {

	binding := rawSessionBinding(ses)

	if binding.outbound != nil {
		log.Infof("client erase: sid: %d pid: %s", ses.ID(), binding.outbound.String())
		ses.SetTag(nil)
		return binding.outbound
	}

	return nil
}

func Status() string {

	var buffer bytes.Buffer

	buffer.WriteString("\n=========Client Status=========\n")

	acceptor.VisitSession(func(ses cellnet.Session) bool {

		backendPID, outboundPID := GetSessionBinding(ses)

		buffer.WriteString(fmt.Sprintf("client sid:%s outbound: %s  backend: %s\n", ses.ID(), outboundPID.String(), backendPID.String()))

		return true
	})

	return buffer.String()
}
