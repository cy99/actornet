package nexus

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
)

func ConnectSingleton(address string, domain string) {

	con(address, domain)
}

// 启动本机的listen
func con(address string, domain string) {

	peer := socket.NewConnector(nil)
	peer.Start(address)

	cellnet.RegisterMessage(peer, "coredef.SessionConnected", func(ev *cellnet.Event) {

		sendDomains(ev.Ses)
	})

	shareInit(peer)
}
