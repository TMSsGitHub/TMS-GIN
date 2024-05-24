package handler

import (
	"TMS-GIN/common"
	"TMS-GIN/errors"
	"TMS-GIN/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindQuery(&user); err != nil {
		fmt.Println(err)
		c.Error(errors.NewServerError(500, "参数错误", err))
		return
	}
	c.Set("resp", resp.Success(user))
}

//func (con *UserController) Login(c *gin.Context) {}
