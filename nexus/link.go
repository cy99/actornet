package nexus

import (
	"bytes"
	"fmt"
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/util"
	"github.com/davyxu/cellnet"
)

// =============================================
// 管理进程内通过svcid标示的,到各服务器的连接
// =============================================
var (
	sesLinkAddress *util.DuplexMap
)

// 对一个服务器进程来说, 连到其他服务的, 只有1个

// 通过给定远方的ServiceID, 来获取其session
func ResolveAddressToSession(addr string) cellnet.Session {

	if raw, ok := sesLinkAddress.MainBySlave(addr); ok {
		return raw.(cellnet.Session)
	}

	return nil
}

func AddressBySession(ses cellnet.Session) string {

	if raw, ok := sesLinkAddress.SlaveByMain(ses); ok {
		return raw.(string)
	}

	return ""
}

func SendToSession(addr string, msg interface{}) {

	ses := ResolveAddressToSession(addr)
	if ses == nil {

		log.Errorln("service not ready:", addr, Status())
		return
	}

	ses.Send(msg)
}

func Broadcast(msg interface{}) {

	sesLinkAddress.Visit(func(main, slave interface{}) bool {

		main.(cellnet.Session).Send(msg)

		return true
	})
}

func addServiceSession(address string, ses cellnet.Session) {

	if pre, ok := sesLinkAddress.MainBySlave(address); ok {
		log.Warnf("duplicate svc session: %v, pre sesid:%d", address, pre.(cellnet.Session).ID())
	}

	sesLinkAddress.Add(ses, address)

	log.Infof("svc session attach: %v  sid: %d", address, ses.ID())
}

func removeServiceSession(ses cellnet.Session) {

	if raw, err := sesLinkAddress.RemoveByMain(ses); err == nil {
		log.Infof("svc session detach: %v sid: %d", raw.(string), ses.ID())
	}
}

func Status() string {

	var buffer bytes.Buffer

	buffer.WriteString("=========Link Status=========\n")

	sesLinkAddress.Visit(func(main, slave interface{}) bool {

		ses := main.(cellnet.Session)
		address := slave.(string)

		buffer.WriteString(fmt.Sprintf("svcid:%s id:%d \n", address, ses.ID()))

		return true
	})

	return buffer.String()
}

func init() {

	actor.OnReset.Add(func(...interface{}) error {

		sesLinkAddress = util.NewDuplexMap()

		return nil
	})

}
