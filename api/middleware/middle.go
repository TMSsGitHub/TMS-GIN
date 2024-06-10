package middle

import "github.com/gin-gonic/gin"

func InitMiddleware(r *gin.Engine) {
	// 注册跨域、统一响应体、认证中间件
	// todo 把认证中间件放到具体路由里更具灵活性
	r.Use(CorsMiddle(), Resp, auth)
}
