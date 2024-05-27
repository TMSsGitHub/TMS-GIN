package router

import (
	"TMS-GIN/api/handler"
	"github.com/gin-gonic/gin"
)

func InitAPIRouter(r *gin.Engine) {
	userRouter(r)
}

func userRouter(c *gin.Engine) {
	userRouter := c.Group("user")
	{
		userRouter.GET("/:username/:password", handler.GetUserInfo)
		userRouter.POST("", handler.CreateUser)
	}
}
