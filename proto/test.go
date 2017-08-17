package proto

type TestMsgACK struct {
	Msg string
}

// 客户端请求后台服务器绑定
type BindClientREQ struct {
	ClientSessionID int64 // 网关上的id
}

type BindClientACK struct {
	ClientSessionID int64 // 网关上的id
	ID              string
}
