package handler

import (
	"TMS-GIN/internal/common"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	c.Set(resp.RES, resp.Success("ok"))
	//var user model.User
	//if err := c.ShouldBindQuery(&user); err != nil {
	//	fmt.Println(err)
	//	c.Error(errors.NewServerError("参数错误", err))
	//	return
	//}
	//c.Set("res", resp.Success(user))
}

//func CreateUser(c *gin.Context) { // todo 删去
//	var user model.User
//	if err := c.ShouldBindJSON(&user); err != nil {
//		fmt.Println(err)
//		c.Error(errors.NewServerError("参数错误", err))
//		return
//	}
//	fmt.Println("user:", user.Phone)
//	res := datastore.DB.Create(&user)
//	if err := res.Error; err != nil {
//		c.Error(errors.NewServerError("新建用户失败", err))
//		return
//	}
//	c.Set(resp.RES, resp.Success(user))
//}

//func (con *UserController) Login(c *gin.Context) {}
