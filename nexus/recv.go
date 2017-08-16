package nexus

import (
	"container/list"
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
	"sync"
)

var (
	hijackList      *list.List
	hijackListGuard sync.Mutex
)

func addHijack(proc *nexusProcess) {

	hijackListGuard.Lock()
	hijackList.PushBack(proc)
	hijackListGuard.Unlock()
}

func checkHijack(m *actor.Message) bool {

	for {

		hijackListGuard.Lock()
		elem := hijackList.Front()
		hijackListGuard.Unlock()

		if elem == nil {
			break
		}

		proc := elem.Value.(*nexusProcess)

		if proc.hijack(m) {
			hijackListGuard.Lock()
			hijackList.Remove(elem)
			hijackListGuard.Unlock()
			return false
		}
	}

	// 没有需要处理的rpc
	return true
}

func onRouter(ev *cellnet.Event) {

	msg := ev.Msg.(*proto.Route)

	userMsg, err := cellnet.DecodeMessage(msg.MsgID, msg.MsgData)
	if err != nil {
		log.Errorln(err)
		return
	}

	address := getDomainBySession(ev.Ses)

	tgtProc := actor.LocalPIDManager.GetByID(msg.TargetID)

	if tgtProc != nil {

		m := &actor.Message{
			Data:      userMsg,
			TargetPID: tgtProc.PID(),
			CallID:    msg.CallID,
		}

		if msg.SourceID != "" {
			m.SourcePID = actor.NewPID(address, msg.SourceID)
		}

		if checkHijack(m) {
			tgtProc.Notify(m)
		}

	} else {
		log.Errorln("node not found:", msg.TargetID)
	}

}

func init() {

	actor.OnReset.Add(func(...interface{}) error {

		hijackList = list.New()

		return nil
	})

}
