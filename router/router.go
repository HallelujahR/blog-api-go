package router

import (
	"blog/controller/user"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter(g *gin.Engine) *gin.Engine {
	g = InitUserRouter(g)
	return g
}

// 用户相关路由
func InitUserRouter(g *gin.Engine) *gin.Engine {
	u := g.Group("/user")
	{
		u.GET("/register", user.Test)
	}
	return g

}
