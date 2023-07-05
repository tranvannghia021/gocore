package singletons

import (
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/vars"
	"sync"
)

var lock = &sync.Mutex{}

func InstanceUser() *repositories.Core {
	if vars.User == nil {
		lock.Lock()
		defer lock.Unlock()
		vars.User = &repositories.Core{}
	}
	return vars.User
}

func InstancePayload() *vars.PayloadGenerate {
	if vars.Payload == nil {
		lock.Lock()
		defer lock.Unlock()
		vars.Payload = &vars.PayloadGenerate{}
	}
	return vars.Payload
}
