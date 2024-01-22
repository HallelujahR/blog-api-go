package bootstrap

import (
	"blog/library/resource"
	"blog/router"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 必须进行初始化
func MustInit(g *gin.Engine) *gin.Engine {
	c := router.InitRouter(g)

	InitMysql(g)
	InitRedis(g)
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
func InitRedis(g *gin.Engine) {
	err := InitConf(g, "redis")
	if err != nil {
		panic(err)
	}
	redisConf := resource.Config.Redis
	dns := fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     dns,                // Redis 地址
		Password: redisConf.Password, // Redis 密码，没有则留空
		DB:       redisConf.DB,       // 使用默认 DB
	})
	// 检查连接
	err = rdb.Ping().Err()
	if err != nil {
		panic(err)
	}
	resource.RedisClient = rdb

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
