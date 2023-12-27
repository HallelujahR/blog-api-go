package dao

import (
	"blog/library/resource"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var CliMysqlMain *gorm.DB
var ormMainOnce sync.Once

func OrmLazyInit() {
	ormMainOnce.Do(func() {
		CliMysqlMain = resource.MySQLClient
	})
}

// NewDbDao 新建Dao
func NewDbDao(g *gin.Context) *gorm.DB {
	OrmLazyInit()
	return CliMysqlMain.WithContext(g)
}
