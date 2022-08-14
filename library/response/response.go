package response

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

var logger *glog.Logger

func init() {
	logger = g.Log("debug")
}

type Role struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Router string `json:"router"`
}

type User struct {
	Id       int    `json:"id"`
	Passport string `json:"passport"`
	Token    string `json:"token"`
	RoleInfo Role
	RoleId   int `json:"roleId"`
}
