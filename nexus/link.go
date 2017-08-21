package nexus

import (
	"bytes"
	"fmt"
	"github.com/davyxu/cellnet"
	"runtime"
	"sync"
)

var (
	sesByDomain      map[string]cellnet.Session
	sesByDomainGuard sync.RWMutex
)

// 根据域名, 解析其对应的会话
func resolveSessionByDomain(domain string) cellnet.Session {

	sesByDomainGuard.RLock()

	defer sesByDomainGuard.RUnlock()

	if ses, ok := sesByDomain[domain]; ok {
		return ses
	}

	return nil
}

// 根据会话, 取其域名
func getDomainBySession(ses cellnet.Session) string {

	if tag := ses.Tag(); tag != nil {
		return tag.(string)
	}

	return ""
}

// 给域名服务器发送消息
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

func removeServiceSession(ses cellnet.Session) string {

	if domain := getDomainBySession(ses); domain != "" {

		log.Infof("Domain detach, domain: %s sid: %d", domain, ses.ID())

		delete(sesByDomain, domain)

		return domain
	}

	return ""
}

// 等待服务器连接上
func WaitReady(domain string) {

	for {

		if resolveSessionByDomain(domain) != nil {
			break
		}

		runtime.Gosched()
	}

}

// 当前连接状态
func Status() string {

	var buffer bytes.Buffer

	buffer.WriteString("\n=========Link Status=========\n")

	for domain, ses := range sesByDomain {

		buffer.WriteString(fmt.Sprintf("domain: %s sid: %d \n", domain, ses.ID()))
	}

	return buffer.String()
}

func init() {

	sesByDomain = make(map[string]cellnet.Session)

}
