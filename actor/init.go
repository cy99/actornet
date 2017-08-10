package actor

import "github.com/davyxu/actornet/util"

var OnReset util.Delegate

func StartSystem() {

	OnReset.Invoke()
}

var inited bool

func init() {

	OnReset.Add(func(...interface{}) error {

		inited = true

		return nil
	})

}
