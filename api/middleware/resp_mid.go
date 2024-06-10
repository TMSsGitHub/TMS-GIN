package middle

import (
	"TMS-GIN/internal/common"
	cusErr "TMS-GIN/internal/errors"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Resp(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Fail("服务器内部错误"))
		}
	}()
	c.Next()
	for _, err := range c.Errors {
		var serverErr cusErr.ServerError
		switch {
		case errors.As(err.Err, &serverErr):
			c.JSON(http.StatusOK, resp.FailWithCode(serverErr.Code, serverErr.Msg))
		default:
			c.JSON(http.StatusOK, resp.Fail("服务器异常"))
		}
		c.Abort()
		return
	}

	if res, ok := c.Get(resp.RES); ok {
		c.JSON(http.StatusOK, res)
	} else {
		//c.JSON(http.StatusInternalServerError, gin.H{})
	}
}
