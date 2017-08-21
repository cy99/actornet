package mailbox

type MailReceiver interface {
	OnRecv(interface{})
}

type MailBox interface {
	Post(interface{})
	Start(MailReceiver)
	Hijack(func(interface{}) bool)
}
