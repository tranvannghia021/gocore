package socials

import "github.com/tranvannghia021/gocore/src/repositories"

func newFacebook(coreConfig *repositories.ConfigSocial) {
	coreConfig.UsePKCE = false
	coreConfig.Scopes = []string{}
	coreConfig.Separator = ","
}
