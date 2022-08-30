// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OperationLog is the golang structure for table operation_log.
type OperationLog struct {
	Id        int         `json:"id"        ` //
	UserId    int         `json:"userId"    ` //
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 修改时间
	DeletedAt *gtime.Time `json:"deletedAt" ` // 删除时间
	IsDeleted int         `json:"isDeleted" ` // 数据的逻辑删除
	Status    int         `json:"status"    ` // 状态
	Extra     string      `json:"extra"     ` //
	Ip        string      `json:"ip"        ` // 请求ip
	Path      string      `json:"path"      ` // 请求路径
	Method    string      `json:"method"    ` // 请求方法
	Request   string      `json:"request"   ` // 请求参数
	Response  string      `json:"response"  ` // 响应内容
	Agent     string      `json:"agent"     ` // 代理
	Latency   int         `json:"latency"   ` // 延迟
	ErrMsg    string      `json:"errMsg"    ` // 错误信息
	UserName  string      `json:"userName"  ` // 用户名
}
