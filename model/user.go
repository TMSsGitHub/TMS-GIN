package model

type User struct {
	Username string `form:"username" binding:"required"`
	Password int32  `form:"password" binding:"required"`
}
