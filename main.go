package main

import (
	"blog/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	g := bootstrap.MustInit(r)

	g.Run("0.0.0.0:8888")
}
