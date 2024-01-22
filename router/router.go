package router

import (
	securityquestion "blog/controller/security_question"
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
	// 用户注册密码相关路由
	u := g.Group("/user")
	{
		u.POST("/login", user.Login)                                                 //登陆
		u.POST("/register", user.Register)                                           //注册
		u.GET("/security_question", securityquestion.GetSecurityQuestion)            //密保问题获取
		u.POST("/reset_password", user.ResetPassword)                                //重制密码
		u.POST("/get_security_question", securityquestion.GetSecurityQuestionByUser) //重制密码获取密保问题
		u.POST("/check_security_question", securityquestion.CheckSecurityQuestion)   //验证密保问题

	}
	//主页

	return g

}
