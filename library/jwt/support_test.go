package jwt

import (
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestHeader(t *testing.T) {
	base64header := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	ans := IsSupport(base64header)
	g.Dump(ans)
}
