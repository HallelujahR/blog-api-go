package middleware

import (
	"blog/library/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// 如果header中不存在token，则返回错误
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "登录状态验证失败"})
			c.Abort()
			return
		}

		// 验证token是否有效
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "登录状态已过期！"})
			c.Abort()
			return
		}

		// 在context中设置用户信息，后续的处理函数可以使用
		c.Set("phonenumber", claims.Phonenumber)
		c.Next()
	}
}
