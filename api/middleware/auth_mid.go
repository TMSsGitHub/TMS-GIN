package middle

import (
	resp "TMS-GIN/internal/common"
	"TMS-GIN/internal/errors"
	"TMS-GIN/internal/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("token")

	if token == "" {
		c.Error(errors.SimpleErrorWithCode(resp.NeedToLogin, "请先登录"))
		c.Abort()
	}
	claim, err := utils.ValidateAccessToken(token)
	if err != nil { // access_token验证未通过
		msg := err.Error()
		// 过期
		if strings.Contains(msg, "token is expired") {
			c.JSON(resp.LoginExpired, resp.Fail("access已过期"))
			c.Abort()
			return
		}
		// 其他解析错误
		c.Error(errors.SimpleErrorWithCode(resp.NeedToLogin, "请重新登录"))
		c.Abort()
		return
	}
	c.Set("claim", claim) // todo 用途未定
	c.Next()
}
