package proto

type TestMsgACK struct {
	Msg string
}

// 客户端请求后台服务器绑定
// client -> gate -> gate_assit
type BindClientREQ struct {
	ClientSessionID int64 // 网关上的id (透传)

}

// gate_assit -> gate_receiptor -> client
type BindClientACK struct {
	ClientSessionID int64  // 网关上的id (透传)
	ID              string // 用户后台的网关用户pid.ID
}
