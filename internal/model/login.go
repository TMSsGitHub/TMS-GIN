package model

type Account struct {
	Phone string `json:"phone" binding:"required"`
	Pwd   string `json:"pwd"`
	Code  string `json:"code"`
	Mode  int    `json:"mode"` // 暂定传入1为密码登录
}
