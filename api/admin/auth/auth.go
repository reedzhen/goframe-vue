package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/model/entity"
)

type LoginReq struct {
	g.Meta      `path:"/auth/login" method:"post" summary:"登录" tags:"用户登录" noAuth:"true"`
	Username    string `json:"username" v:"required|length:3,12"`
	Password    string `json:"password" v:"required|length:6,16"`
	CaptchaKey  string `json:"captchaKey" v:"required"`
	CaptchaCode string `json:"captchaCode" v:"required"`
}

type LoginRes struct {
	AccessToken string `json:"accessToken" dc:"token"`
}

type LogoutReq struct {
	g.Meta `path:"/auth/logout" method:"post" summary:"登出" tags:"用户登录" noAuth:"true"`
}
type LogoutRes struct{}

type GetInfoReq struct {
	g.Meta `path:"/auth/user" method:"get" tags:"用户登录" summary:"获取当前用户信息" noAuth:"true"`
}
type GetInfoRes struct {
	UserItem
	Authorities []*entity.SysMenu `json:"authorities" dc:"权限"`
}

// UserItem 用户表包含密码等敏感数据，这里只返回非敏感数据
type UserItem struct {
	Id       int64  `json:"id" dc:"用户id"`
	Username string `json:"username"  dc:"账号"`
	Nickname string `json:"nickname"  dc:"昵称"`
	RealName string `json:"realName"  dc:"昵称"`
	Avatar   string `json:"avatar"  dc:"头像"`
}
