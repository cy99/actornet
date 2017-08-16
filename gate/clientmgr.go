package gate

import (
	"bytes"
	"fmt"
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/util"
	"github.com/davyxu/cellnet"
)

var (
	sesLinkPID *util.DuplexMap
)

func ServiceSessionByPID(pid *actor.PID) cellnet.Session {

	if raw, ok := sesLinkPID.MainBySlave(pid); ok {
		return raw.(cellnet.Session)
	}

	return nil
}

func PIDBySession(ses cellnet.Session) *actor.PID {

	if raw, ok := sesLinkPID.SlaveByMain(ses); ok {
		return raw.(*actor.PID)
	}

	return nil
}

func SendToClient(pid *actor.PID, msg interface{}) {

	ses := ServiceSessionByPID(pid)
	if ses == nil {

		log.Errorln("client not exists:", pid.String(), Status())
		return
	}

	ses.Send(msg)
}

func BroardcastToClients(msg interface{}) {

	sesLinkPID.Visit(func(main, slave interface{}) bool {

		main.(cellnet.Session).Send(msg)

		return true
	})
}

func BroardcastToActor(msg interface{}) {

	sesLinkPID.Visit(func(main, slave interface{}) bool {

		slave.(*actor.PID).NotifyData(msg)

		return true
	})
}

func addClient(pid *actor.PID, ses cellnet.Session) {

	sesLinkPID.Add(ses, pid)

	log.Infof("client attach, sid: %d  pid: %s", ses.ID(), pid.String())
}

func removeClient(ses cellnet.Session) *actor.PID {

	if raw, err := sesLinkPID.RemoveByMain(ses); err == nil {

		pid := raw.(*actor.PID)
		log.Infof("client detach: sid: %d pid: %s", ses.ID(), pid.String())
		return pid
	}

	return nil
}

func Status() string {

	var buffer bytes.Buffer

	buffer.WriteString("=========Link Status=========\n")

	sesLinkPID.Visit(func(main, slave interface{}) bool {

		ses := main.(cellnet.Session)
		pid := slave.(*actor.PID)

		buffer.WriteString(fmt.Sprintf("client sid:%s pid: %s \n", ses.ID(), pid.String()))

		return true
	})

	return buffer.String()
}

func init() {

	actor.OnReset.Add(func(...interface{}) error {

		sesLinkPID = util.NewDuplexMap()

		return nil
	})

}
