package nexus

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/cellnet"
	"runtime"
)

func sendToDomain(domainName string, msg interface{}) {

	domain := actor.GetDomain(domainName)
	if domain == nil {
		log.Errorf("Remote domain not found: %s", domainName)
		return
	}

	if domain.RemoteContext == nil {
		log.Errorf("NOT Remote domain: %s", domainName)
		return
	}

	ses := domain.RemoteContext.(cellnet.Session)
	ses.Send(msg)
}

// 等待服务器连接上
func WaitReady(domainName string) {

	for {

		domain := actor.GetDomain(domainName)
		if domain != nil {
			break
		}

		runtime.Gosched()
	}

}
