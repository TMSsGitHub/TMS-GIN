package main

import (
	"TMS-GIN/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.UserRouter(r)
	r.Run(":8080")
}
