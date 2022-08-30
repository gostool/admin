package cmd

import (
	"admin/internal/model"
	"admin/internal/service"
	"context"
	"errors"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	DefaultUserGuest     = "guest"
	DefaultGuestPassword = "guest"
	DefaultGuestNickname = "guest"
)

func ToRoleIds(roleId int) string {
	roleIds := []int{roleId}
	return gconv.String(roleIds)
}

func userSave(ctx context.Context, name, password, nickname, roleIds string) {
	password, err := gmd5.EncryptString(password)
	if err != nil {
		logger.Fatal(ctx, err)
	}
	data, err := service.User().Save(ctx, name, password, nickname, guestRoleId, roleIds)
	if err != nil {
		logger.Fatal(ctx, err)
	}
	logger.Debugf(ctx, "user passport:%v ", name)
	logger.Debugf(ctx, "user passwd:%v ", password)
	logger.Debugf(ctx, "user init ok:%v ", data)
}

func userGuestInit(ctx context.Context) {
	type User struct {
		Passport string
		Password string
		Nickname string
		RoleIds  string
	}

	userList := []User{
		{DefaultUserGuest, DefaultGuestPassword, DefaultGuestNickname, ToRoleIds(guestRoleId)},
	}
	for _, user := range userList {
		userSave(ctx, user.Passport, user.Password, user.Nickname, user.RoleIds)
	}
}

func userCreateAdmin(ctx context.Context, passport, password, nickname string) {
	var err error
	if len(password) < 6 {
		err = errors.New("password must >=6")
		logger.Fatal(ctx, err)
	}
	in := model.UserCreateInput{}
	in.Passport = passport
	in.Password = password
	in.Nickname = nickname
	in.RoleId = adminRoleId
	in.RoleIds = ToRoleIds(in.RoleId)
	password, err = gmd5.EncryptString(in.Password)
	if err != nil {
		logger.Fatal(ctx, err)
	}
	in.Password = password
	id, err := service.User().Create(ctx, in)
	if err != nil {
		logger.Fatal(ctx, err)
	}
	logger.Debugf(ctx, "user create ok:%v", id)
}

func createAdmin(ctx context.Context) {
	name := gcmd.Scan("> What's your super user name?\n")
	password := gcmd.Scan("> What's your super user password?\n")
	nickname := gcmd.Scan("> What's your super user nickname?\n")
	userCreateAdmin(ctx, name, password, nickname)
}
