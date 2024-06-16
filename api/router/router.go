package router

import (
	"TMS-GIN/api/handler"
	middle "TMS-GIN/api/middleware"
	log "TMS-GIN/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 在这里注册路由
func InitRouter(r *gin.Engine) {
	// 找不到路由，重定向
	r.NoRoute(func(c *gin.Context) {
		log.Logger.Info("??")
		c.Redirect(http.StatusFound, "https://www.baidu.com") // fixme 重定向地址需要更改
	})
	notAuthRouter(r)
	// 需要通过登录认证的使用authRouter创建路由组
	auth := r.Group("")
	auth.Use(middle.Auth)
	authRouter(auth)
}

func authRouter(r *gin.RouterGroup) {
	userRouter := r.Group("user")
	{
		userRouter.GET("", handler.GetUserInfo)
	}

	fileRouter := r.Group("file")
	{
		fileRouter.POST("upload", handler.FileUpload)
		fileRouter.GET("download/:url", handler.FileDownload)
	}
}

func notAuthRouter(r *gin.Engine) {
	accountRouter := r.Group("account")
	{
		accountRouter.POST("login", handler.Login)
		accountRouter.POST("register", handler.Register)
	}
}
