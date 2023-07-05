package socials

import (
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/src/service"
	"github.com/tranvannghia021/gocore/vars"
)

var pinterest = "pinterest"

type sPinterest struct {
	http *service.SHttpRequest
}

func (s *sPinterest) loadConfig() {
	//TODO implement me
	panic("implement me")
}

func (s *sPinterest) getToken(code string) vars.ResReq {
	//TODO implement me
	panic("implement me")
}

func (s *sPinterest) profile(token string) repositories.Core {
	//TODO implement me
	panic("implement me")
}
