package jwt

const base64AlgHS256 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
const base64AlgHS384 = "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9"
const base64AlgHS512 = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9"

var supportHeaderAlg = map[string]bool{
	base64AlgHS256: true,
	base64AlgHS384: true,
	base64AlgHS512: true,
}

func IsSupport(header string) bool {
	if _, ok := supportHeaderAlg[header]; ok {
		return true
	}
	return false
}
