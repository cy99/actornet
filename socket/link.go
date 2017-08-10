package socket

import (
	"bytes"
	"fmt"
	"github.com/davyxu/cellnet"
	"sync"
)

// =============================================
// 管理进程内通过svcid标示的,到各服务器的连接
// =============================================
var (
	sesByAddress = make(map[string]cellnet.Session)
	addressBySes = make(map[cellnet.Session]string)
	svclinkGuard sync.RWMutex
)

// 对一个服务器进程来说, 连到其他服务的, 只有1个

// 通过给定远方的ServiceID, 来获取其session
func ServiceSessionByServiceID(addr string) cellnet.Session {

	svclinkGuard.RLock()
	defer svclinkGuard.RUnlock()

	if ses, ok := sesByAddress[addr]; ok {
		return ses
	}

	return nil
}

func AddressBySession(ses cellnet.Session) (string, bool) {

	svclinkGuard.RLock()
	defer svclinkGuard.RUnlock()

	if addr, ok := addressBySes[ses]; ok {
		return addr, true
	}

	return "", false
}

func SendToSession(addr string, msg interface{}) {

	ses := ServiceSessionByServiceID(addr)
	if ses == nil {

		log.Errorln("service not ready:", addr, Status())
		return
	}

	ses.Send(msg)
}

func addServiceSession(address string, ses cellnet.Session) {

	svclinkGuard.Lock()
	if pre, ok := sesByAddress[address]; ok {
		log.Warnf("duplicate svc session: %v, pre sesid:%d", address, pre.ID())
	}

	sesByAddress[address] = ses
	addressBySes[ses] = address
	svclinkGuard.Unlock()

	log.Infof("svc session attach: %v  sid: %d", address, ses.ID())
}

func removeServiceSession(ses cellnet.Session) string {

	svclinkGuard.Lock()
	defer svclinkGuard.Unlock()

	if id, ok := addressBySes[ses]; ok {

		log.Infof("svc session detach: %v sid: %d", id, ses.ID())

		delete(addressBySes, ses)
		delete(sesByAddress, id)

		return id
	}

	return ""
}

func Status() string {

	var buffer bytes.Buffer

	buffer.WriteString("=========Link Status=========\n")

	for address, ses := range sesByAddress {

		buffer.WriteString(fmt.Sprintf("svcid:%s id:%d \n", address, ses.ID()))

	}

	return buffer.String()
}
