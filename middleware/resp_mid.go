package middle

import (
	"TMS-GIN/common"
	"TMS-GIN/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Resp(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Fail("服务器异常"))
		}
	}()
	c.Next()
	for _, err := range c.Errors {
		switch spcificErr := err.Err.(type) {
		case errors.ServerError:
			c.JSON(http.StatusOK, resp.Fail(spcificErr.Error()))
		default:
			c.JSON(http.StatusOK, resp.Fail("服务器内部错误"))
		}
		c.Abort()
		return
	}

	if resp, ok := c.Get("resp"); ok {
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
}
