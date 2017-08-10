package actor

import "github.com/davyxu/actornet/util"

var OnReset util.Delegate

func StartSystem() {

	OnReset.Invoke()
}
