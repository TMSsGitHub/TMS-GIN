package main

import (
	"TMS-GIN/api/router"
	"TMS-GIN/config"
	"TMS-GIN/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middle.Resp)
	router.UserRouter(r)
	r.Run(fmt.Sprintf(":%s", config.Cfg.App.Port))
}
