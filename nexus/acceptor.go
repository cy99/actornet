package nexus

import (
	"fmt"
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/proto"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
	"sync/atomic"
)

var multiInsDomainSeq int64

// 连接上来具备多实例时, 按序号命名
func multiInstanceDomain(domain string) string {

	id := atomic.AddInt64(&multiInsDomainSeq, 1)

	return fmt.Sprintf("%s_%d", domain, id)
}

// 启动本机的listen
func Listen(address string, domain string) {

	peer := socket.NewAcceptor(nil)
	peer.Start(address)

	actor.LocalPIDManager.Domain = domain

	cellnet.RegisterMessage(peer, "proto.DomainIdentifyACK", func(ev *cellnet.Event) {
		msg := ev.Msg.(*proto.DomainIdentifyACK)

		ev.Send(&proto.DomainIdentifyACK{
			Domain: domain,
		})

		var remoteDomain string
		if msg.Singleton {
			remoteDomain = msg.Domain
		} else {
			remoteDomain = multiInstanceDomain(msg.Domain)
		}

		broardCast(&proto.NexusOpen{
			Domain: remoteDomain,
		})

		addServiceSession(remoteDomain, ev.Ses)
	})

	shareInit(peer)

}
