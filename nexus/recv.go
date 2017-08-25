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

	msg := ev.Msg.(*proto.RouteACK)

	userMsg, err := cellnet.DecodeMessage(msg.MsgID, msg.MsgData)
	if err != nil {
		log.Errorln(err)
		return
	}

	domain := actor.GetDomain(msg.Target.Domain)

	if domain == nil {
		log.Errorf("domain not found: %s", msg.Target.Domain)
		return
	}

	tgtProc := domain.GetByID(msg.Target.Id)

	if tgtProc != nil {

		m := &actor.Message{
			Data:      userMsg,
			TargetPID: tgtProc.PID(),
			CallID:    msg.CallID,
		}

		if msg.Source.IsValid() {
			m.SourcePID = &actor.PID{Domain: msg.Source.Domain, Id: msg.Source.Id}
		}

		if checkHijack(m) {
			tgtProc.Tell(m)
		}

	} else {
		log.Errorf("target not found,  id: %s", msg.Target.String())
	}

}

func init() {

	hijackList = list.New()

}
