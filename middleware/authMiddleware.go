// 鉴权中间件

package middleware

import (
	"github.com/gin-gonic/gin"
	"go-vue/common/constant"
	"go-vue/common/result"
	"strings"
)

// AuthMiddleware 鉴权
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			result.Failed(c, int(result.ApiCode.NOAUTH),
				result.ApiCode.GetMessage(result.ApiCode.NOAUTH))
			c.Abort() //终止当前请求处理，并阻止后续中间件或函数执行
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			result.Failed(c, int(result.ApiCode.AUTHFORMMATERROR),
				result.ApiCode.GetMessage(result.ApiCode.AUTHFORMMATERROR))
			c.Abort()
			return
		}
		// todo 检验token
		var token = "token"
		// 存用户信息
		c.Set(constant.ContextKeyUserObj, token)
		c.Next()
	}
}
