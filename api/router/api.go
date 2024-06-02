package router

import (
	"TMS-GIN/api/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitAPIRouter(r *gin.Engine) {
	// 找不到路由，重定向
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://www.baidu.com")
	})
	userRouter(r)
}

func userRouter(c *gin.Engine) {
	userRouter := c.Group("user")
	{
		userRouter.GET("", handler.GetUserInfo)
		userRouter.POST("", handler.CreateUser)
	}
}
