package controller

import (
	"TMS-GIN/dao"
	"github.com/gin-gonic/gin"
	"sync"
)

type UserController struct {
	BaseController
}

var (
	instance UserController
	once     sync.Once
)

func GetUserInstance() *UserController {
	once.Do(func() {
		instance = UserController{}
	})
	return &instance
}

func (con *UserController) GetUserInfo(c *gin.Context) {
	var user dao.User
	if c.ShouldBindQuery(&user) != nil {
		con.fail(c, "error")
		return
	}
	//fmt.Println(user)
	con.success(c, user)
}

//func (con *UserController) Login(c *gin.Context) {}
