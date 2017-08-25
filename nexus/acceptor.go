package nexus

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
)

// 启动本机的listen
func Listen(address string, domain string) {

	peer := socket.NewAcceptor(nil)
	peer.Start(address)

	cellnet.RegisterMessage(peer, "coredef.SessionAccepted", func(ev *cellnet.Event) {

		sendDomains(ev.Ses)
	})

	shareInit(peer)

}

func sendDomains(ses cellnet.Session) {

	var msg proto.DomainSyncACK

	actor.VisitDomains(func(domain *actor.Domain) bool {

		msg.DomainNames = append(msg.DomainNames, domain.Name)

		return true
	})

	ses.Send(&msg)

}

func addDomains(domainNames []string, ses cellnet.Session) {

	for _, name := range domainNames {

		if actor.GetDomain(name) != nil {
			log.Errorf("Duplicate remote domain: %s", name)
			continue
		}

		domain := actor.CreateRemoteDomain(name)
		domain.RemoteContext = ses
	}

}

func removeDomains(ses cellnet.Session) {

	var removeNames []string
	actor.VisitDomains(func(domain *actor.Domain) bool {

		if domain.RemoteContext != nil {

			if domain.RemoteContext.(cellnet.Session) == ses {
				removeNames = append(removeNames, domain.Name)
			}

		}

		return true
	})

	for _, name := range removeNames {
		actor.DestroyDomain(name)
	}

}
