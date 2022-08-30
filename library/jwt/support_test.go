package jwt

import (
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestHeader(t *testing.T) {
	base64HS256 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	base64HS384 := "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9"
	base64HS512 := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9"
	base64Hmac := []string{base64HS256, base64HS384, base64HS512}
	for _, v := range base64Hmac {
		ans := IsSupport(v)
		g.Dump(ans)
	}
}
