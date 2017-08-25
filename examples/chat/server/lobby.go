package main

import (
	"github.com/davyxu/actornet/actor"
	"github.com/davyxu/actornet/examples/chat/proto"
)

type lobby struct {
	actor.LocalProcess

	userByDomain map[string]*actor.PID
}

func (self *lobby) Broardcast(data interface{}) {

	self.BroardcastBySender(data, self.PID())
}

func (self *lobby) BroardcastBySender(data interface{}, sender *actor.PID) {
	for _, u := range self.userByDomain {

		u.TellBySender(data, sender)
	}
}

func (self *lobby) addUser(pid *actor.PID, sourceDomain string) {

	self.userByDomain[sourceDomain] = pid
}

func (self *lobby) removeUser(sourceDomain string) {

	delete(self.userByDomain, sourceDomain)
}
func (self *lobby) getUser(sourceDomain string) *actor.PID {

	if u, ok := self.userByDomain[sourceDomain]; ok {
		return u
	}

	return nil
}

func (self *lobby) OnRecv(c actor.Context) {
	switch msg := c.Msg().(type) {

	case *chatproto.ChatREQ:

		serverUserPID := self.getUser(c.Source().Domain)

		// 通过rpc获取来源名字
		name := serverUserPID.Call(&chatproto.GetName{}, self.PID()).(*chatproto.GetName).Name

		self.Broardcast(&chatproto.ChatACK{
			User:    serverUserPID.ToProto(),
			Name:    name,
			Content: msg.Content,
		})
	}
}

func newLobby() actor.Actor {
	return &lobby{
		userByDomain: make(map[string]*actor.PID),
	}
}
