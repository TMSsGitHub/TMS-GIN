package middle

import (
	resp "TMS-GIN/internal/common"
	"TMS-GIN/internal/errors"
	"TMS-GIN/internal/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	login    = "/account/login"
	register = "/account/register"
)

// 跳过认证的请求
var skip = map[string]struct{}{
	login:    {},
	register: {},
}

func auth(c *gin.Context) {
	// 跳过不需认证的请求
	token := c.Request.Header.Get("token")
	url := c.Request.URL.Path
	if _, exist := skip[url]; exist {
		c.Next()
		return
	}

	if token == "" {
		c.Error(errors.SimpleErrorWithCode(444, "请先登录"))
		c.Abort()
	}
	claim, err := utils.ValidateAccessToken(token)
	if err != nil { // access_token验证未通过
		msg := err.Error()
		// 过期
		if strings.Contains(msg, "token is expired") {
			c.JSON(409, resp.Fail("access已过期"))
			c.Abort()
			return
		}
		// 其他解析错误
		c.Error(errors.SimpleErrorWithCode(444, "请重新登录"))
		c.Abort()
		return
	}
	c.Set("claim", claim) // todo 用途未定
	c.Next()
}
