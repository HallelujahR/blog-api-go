package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Test(g *gin.Context) {
	fmt.Println("user controller")
}
