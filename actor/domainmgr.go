package actor

import (
	"github.com/davyxu/actornet/proto"
	"sync"
)

var (
	domainByName      = map[string]*Domain{}
	domainByNameGuard sync.Mutex
)

func VisitDomains(callback func(*Domain) bool) {

	for _, domain := range domainByName {
		if !callback(domain) {
			return
		}
	}

}

func MustGetDomain(name string) *Domain {

	domain := GetDomain(name)

	if domain == nil {
		panic("domain not exists: " + name)
	}

	return domain
}

// 找到对应地址的远程pid管理器
func GetDomain(name string) *Domain {

	domainByNameGuard.Lock()

	defer domainByNameGuard.Unlock()

	if dm, ok := domainByName[name]; ok {
		return dm
	}

	return nil
}

func CreateDomain(name string) *Domain {

	return rawCreate(name, true)
}

func rawCreate(name string, isLocal bool) *Domain {
	if GetDomain(name) != nil {
		log.Errorf("Duplicate domain name: %s", name)
	}

	if isLocal {
		log.Debugf("Domain create: %s", name)
	} else {
		log.Debugf("Remote Domain create: %s", name)
	}

	domain := newDomain(name)

	if isLocal {
		domain.Spawn(NewTemplate().WithID("system").WithFunc(func(c Context) {

			switch msg := c.Msg().(type) {
			case *proto.SystemExit:
				Exit(int(msg.Code))
			}

		}))
	}

	domainByNameGuard.Lock()
	domainByName[name] = domain
	domainByNameGuard.Unlock()

	return domain
}

func CreateRemoteDomain(name string) *Domain {

	return rawCreate(name, false)
}

func DestroyDomain(name string) {

	log.Debugf("Domain destroy: %s", name)

	domainByNameGuard.Lock()
	delete(domainByName, name)
	domainByNameGuard.Unlock()
}
