package bootstrap

import (
	"blog/library/resource"
	"blog/router"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 必须进行初始化
func MustInit(g *gin.Engine) *gin.Engine {
	c := router.InitRouter(g)

	InitMysql(g)
	return c
}

// 初始化数据库
func InitMysql(g *gin.Engine) {
	err := InitConf(g, "mysql")
	if err != nil {
		panic(err)
	}
	mysqlConf := resource.Config.DB

	//连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.User,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	resource.MySQLClient = db

}

// 初始化redis
func InitRedis() {

}

// 获取config文件路径
func InitConf(g *gin.Engine, conf string) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	configPath := filepath.Join(path, "conf", "service", conf+".toml") // 使用filepath.Join来构建路径，这样可以确保路径在不同操作系统上都是有效的
	config := &resource.Config
	_, err = toml.DecodeFile(configPath, config)
	if err != nil {
		return err
	}
	return nil
}
