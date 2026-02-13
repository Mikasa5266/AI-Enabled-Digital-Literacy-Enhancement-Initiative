package main

import (
	"fmt"
	"log"

	"gin-demo/internal/handler"
	"gin-demo/internal/model"
	"gin-demo/internal/router"
	"gin-demo/pkg/config"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("配置加载失败", err)
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	handler.InitDB(db)
	db.AutoMigrate(&model.PPT{})
	router.DB = db

	r := gin.Default()
	router.Setup(r)
	log.Printf("服务启动在端口 %s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("服务启动失败", err)
	}
}
