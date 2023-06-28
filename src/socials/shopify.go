package socials

import (
	"fmt"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/vars"
)

var shopify = "shopify"

type sShopify struct {
}

var (
	scopeSp []string
	domain  string
)

func (s sShopify) loadConfig() {
	coreConfig.Separator = ","
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"unauthenticated_read_product_listings",
	}, scopeSp...))
	urlAuth = fmt.Sprintf("https://%s/admin/oauth/authorize", vars.Payload.Domain)

}

func (s sShopify) getToken(code string) vars.ResReq {
	//TODO implement me
	panic("implement me")
}

func (s sShopify) profile(token string) repositories.Core {
	//TODO implement me
	panic("implement me")
}

func (s sShopify) setParameter(rawDomain string) {
	domain = rawDomain
	vars.EndPoint = fmt.Sprintf("https://%s/admin/api/%s", rawDomain, vars.Version)
}
