package middle

import (
	"TMS-GIN/internal/errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func auth(c *gin.Context) {
	token := c.Request.Header.Get("token")
	fmt.Println("token", token)
	if token == "" {
		c.Error(errors.SimpleErrorWithCode(444, "请先登录"))
		c.Abort()
	}
	c.Next()
}
