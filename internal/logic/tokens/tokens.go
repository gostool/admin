package tokens

import (
	"admin/internal/model"
	"admin/internal/service"
	"admin/library/jwt"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"time"
)

type sToken struct {
}

var logger *glog.Logger

const jwtSalt = "test"
const jwtExp = 604800 // 单位s(60*60*24*7)

func init() {
	logger = g.Log("debug")
	instance := New()
	fmt.Println(instance)
	service.RegisterToken(instance)
}

func New() *sToken {
	return &sToken{}
}

func (t *sToken) GenAccessToken(ctx context.Context, r *model.TokenServiceGenTokenReq) (token string, err error) {
	secret := jwt.SecSecret(r.Uid, jwtSalt)
	logger.Debugf(ctx, "uid:%v secret:%v exp:%v", r.Uid, secret, r.Exp)
	token, err = jwt.CreateToken(r.Uid, secret, r.Exp)
	if err != nil {
		logger.Errorf(ctx, "req:%v err:%v token:%v", r, err, token)
		return "", err
	}
	return token, nil
}

// exp: 小于0时，使用配置中的默认时间
func (t *sToken) GenToken(ctx context.Context, uid string, exp int64) (token string, err error) {
	if exp <= 0 {
		exp = time.Now().Add(time.Duration(jwtExp) * time.Second).Unix()
	}
	return t.GenAccessToken(ctx, &model.TokenServiceGenTokenReq{
		Uid: uid,
		Exp: exp,
	})
}

func (t *sToken) CheckToken(ctx context.Context, token string) (uid string, err error) {
	uid, err = jwt.GetUid(token)
	if err != nil {
		logger.Error(ctx, err)
		return uid, errors.New("token格式错误")
	}
	secret := jwt.SecSecret(uid, jwtSalt)
	// 覆盖uid
	uid, err = jwt.AuthToken(token, secret)
	if err != nil {
		logger.Errorf(ctx, "req:%v err:%v uid:%v", token, err, uid)
		return uid, errors.New("token认证失败, 请重新登陆")
	}
	return uid, nil
}
