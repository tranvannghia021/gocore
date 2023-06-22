package socials

func newFacebook() {
	coreConfig.UsePKCE = false
	coreConfig.Scopes = []string{}
	coreConfig.Separator = ","
}
