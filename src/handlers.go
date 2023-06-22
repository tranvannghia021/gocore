package src

import "github.com/tranvannghia021/gocore/src/socials"

func Social(platform string) string {
	return socials.New(platform)
}
