package nexus

import (
	"bytes"
	"fmt"
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/cellnet"
	"runtime"
	"sync"
)

// =============================================
// 管理进程内通过svcid标示的,到各服务器的连接
// =============================================
var (
	sesByDomain      map[string]cellnet.Session
	sesByDomainGuard sync.RWMutex
)

// 对一个服务器进程来说, 连到其他服务的, 只有1个

// 通过给定远方的ServiceID, 来获取其session
func resolveSessionByDomain(domain string) cellnet.Session {

	sesByDomainGuard.RLock()

	defer sesByDomainGuard.RUnlock()

	if ses, ok := sesByDomain[domain]; ok {
		return ses
	}

	return nil
}

func getDomainBySession(ses cellnet.Session) string {

	if tag := ses.Tag(); tag != nil {
		return tag.(string)
	}

	return ""
}

func sendToDomain(domain string, msg interface{}) {

	ses := resolveSessionByDomain(domain)
	if ses == nil {

		log.Errorln("domain not exists:", domain, Status())
		return
	}

	ses.Send(msg)
}

func addServiceSession(domain string, ses cellnet.Session) {

	if preSes := resolveSessionByDomain(domain); preSes != nil {
		log.Warnf("Duplicate Domain, domain: %s, pre sesid:%d", domain, preSes.ID())
	}

	sesByDomainGuard.Lock()
	sesByDomain[domain] = ses
	ses.SetTag(domain)
	sesByDomainGuard.Unlock()

	log.Infof("Domain attach, domain: %s  sid: %d", domain, ses.ID())
}

func removeServiceSession(ses cellnet.Session) {

	if domain := getDomainBySession(ses); domain != "" {

		log.Infof("Domain detach, domain: %s sid: %d", domain, ses.ID())

		delete(sesByDomain, domain)
	}
}

func WaitReady(domain string) {

	// spin lock 自旋锁
	for {

		if resolveSessionByDomain(domain) != nil {
			break
		}

		runtime.Gosched()
	}

}

func Status() string {

	var buffer bytes.Buffer

	buffer.WriteString("=========Link Status=========\n")

	for domain, ses := range sesByDomain {

		buffer.WriteString(fmt.Sprintf("domain: %s sid: %d \n", domain, ses.ID()))
	}

	return buffer.String()
}

func init() {

	actor.OnReset.Add(func(...interface{}) error {

		sesByDomain = make(map[string]cellnet.Session)

		return nil
	})

}
