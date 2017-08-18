package chatproto

import (
	"github.com/davyxu/actornet/proto"
)

// 登录到大厅
// client -> server
type LoginREQ struct {
}

// server -> client
type LoginACK struct {
	User proto.PID // 登录完成后, 要求客户端往自己的actor发消息
}

// 聊天消息
// client -> server
type ChatREQ struct {
	To      proto.PID // 玩家要发给目标
	Content string
}

// 聊天消息
// server -> client
type ChatACK struct {
	Name    string
	Content string
}

// 改名
// client <-> server
type RenameACK struct {
	NewName string
}
