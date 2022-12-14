package jwt

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

type Claims struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}

/*
new secret by uid. it is unique for everyone
*/

func SecSecret(uid, salt string) string {
	return ToMd5(fmt.Sprintf("%v:+-*/:%v", salt, uid))
}

func ToMd5(text string) (decode string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(text))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

/*
before CreateToken:
	secret = SecSecret(uid, salt)
then:
	CreateToken(uid, secret, expireToken)
*/

func CreateToken(uid, secret string, expireToken time.Time) (string, error) {
	now := time.Now()
	if expireToken.Before(now) {
		expireToken = now.Add(time.Hour * 1)
	}
	// 1. Create payload
	claims := Claims{
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireToken),
			Issuer:    uid,
		},
	}

	// 2. Create the tokens using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. Signs the tokens with a secret.
	signedToken, err := token.SignedString([]byte(secret))

	return signedToken, err
}

/*
before AuthToken:
	uid, err := GetUid(signedToken)
	if err != nil {
		return "", err
	}
	secret = SecSecret(uid, secret)
then AuthToken(signedToken, secret)
*/

func AuthToken(signedToken string, secret interface{}) (string, error) {
	// 1.tpl decode by tokens & secret
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return "", fmt.Errorf("AuthToken src.err:%v sercret:%v signedToken:%v", err, secret, signedToken)
	}
	// 2. tokens valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Uid, nil
	}
	return "", err
}

func SafeAuthToken(signedToken string, header string, secret string) (string, error) {
	if !IsSupport(header) {
		return "", errors.New("token error, header alg only support hmac")
	}
	return AuthToken(signedToken, []byte(secret))
}

func GetHeaderAndClaims(token string) (header string, ob Claims, err error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", ob, jwt.NewValidationError("tokens contains an invalid number of segments", jwt.ValidationErrorMalformed)
	}

	payload := parts[1]
	data, err := DecodeSegment(payload)
	if err != nil {
		return "", ob, err
	}
	ob = Claims{}
	err = json.Unmarshal(data, &ob)
	if err != nil {
		return "", ob, err
	}
	return parts[0], ob, nil
}

func DecodeSegment(seg string) ([]byte, error) {
	if jwt.DecodePaddingAllowed {
		if l := len(seg) % 4; l > 0 {
			seg += strings.Repeat("=", 4-l)
		}
		return base64.URLEncoding.DecodeString(seg)
	}

	return base64.RawURLEncoding.DecodeString(seg)
}

/*
tpl:
header: part[0]
payload: part[1]
sign: part[2]
*/

func GetHeaderAndUid(signedToken string) (header string, uid string, err error) {
	header, ob, err := GetHeaderAndClaims(signedToken)
	if err != nil {
		return header, uid, err
	}
	return header, ob.Uid, nil
}
