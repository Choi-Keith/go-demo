package user

import (
	"demo01/pkg/common/errors"
	"demo01/pkg/jwtx"
)

type LoginParamsDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginOAuthParamsDto struct {
	Code  string `form:"code"  binding:"required"`
	State string `form:"state"  binding:"required"`
}

type LoginRespDto struct {
	User      *User           `json:"user"`
	TokenInfo *jwtx.TokenInfo `json:"tokenInfo"`
}

type ListParamsDto struct {
	PageNum  int
	PageSize int
	Username string
	ID       string
}

type ListRespDto struct {
	//Code    int    `json:"code"`
	//Message string `json:"msg"`
	*errors.CodeMessage
	Data *Item `json:"data"`
}

type Item struct {
	Total int                `json:"total"`
	Items []*CreateParamsDto `json:"items"`
}

type CreateParamsDto struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	RePassword  string `json:"rePassword" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Phone       string `json:"phone"`
	WeChatID    string `json:"weChatId"`
	Website     string `json:"website"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Level       int    `json:"level"`
}

type CreateRespDto struct {
	*errors.CodeMessage
	Data string `json:"data"`
}
