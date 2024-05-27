package handler

import (
	"TMS-GIN/config"
	"TMS-GIN/internal/common"
	"TMS-GIN/internal/errors"
	"TMS-GIN/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindQuery(&user); err != nil {
		fmt.Println(err)
		c.Error(errors.NewServerError("参数错误", err))
		return
	}
	c.Set("resp", resp.Success(user))
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.Error(errors.NewServerError("参数错误", err))
		return
	}
	fmt.Println("user:", user.Phone)
	db := config.DB.Create(&user)
	if err := db.Error; err != nil {
		c.Error(errors.NewServerError("新建用户失败", err))
		return
	}
	c.Set("resp", resp.Success(user))
}

//func (con *UserController) Login(c *gin.Context) {}
