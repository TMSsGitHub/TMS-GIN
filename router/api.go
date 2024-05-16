package router

import (
	"TMS-GIN/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(c *gin.Engine) {
	userRouter := c.Group("user")
	userController := controller.GetUserInstance()
	{
		userRouter.GET("/:username/:password", userController.GetUserInfo)
		userRouter.GET()
	}
}
