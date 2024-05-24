package router

import (
	"TMS-GIN/api/handler"
	"github.com/gin-gonic/gin"
)

func UserRouter(c *gin.Engine) {
	userRouter := c.Group("user")
	{
		userRouter.GET("/:username/:password", handler.GetUserInfo)
	}
}
