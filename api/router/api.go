package router

import (
	"TMS-GIN/api/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 在这里注册路由
func InitAPIRouter(r *gin.Engine) {
	// 找不到路由，重定向
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://www.baidu.com")
	})
	userRouter(r)
	fileRouter(r)
	accountRouter(r)
}

func userRouter(c *gin.Engine) {
	userRouter := c.Group("user")
	{
		userRouter.GET("", handler.GetUserInfo)
		userRouter.POST("", handler.CreateUser)
	}
}

func fileRouter(c *gin.Engine) {
	fileRouter := c.Group("file")
	{
		fileRouter.POST("upload", handler.FileUpload)
		fileRouter.GET("download/:url", handler.FileDownload)
	}
}

func accountRouter(c *gin.Engine) {
	accountRouter := c.Group("account")
	{
		accountRouter.POST("login", handler.Login)
	}
}
