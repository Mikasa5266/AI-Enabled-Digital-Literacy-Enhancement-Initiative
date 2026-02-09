package main

import (
	"gin-demo/internal/handler"
	"gin-demo/internal/model"
	"gin-demo/internal/router"
	"gin-demo/pkg/config"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("配置加载失败", err)
	}
	dsn := "root:cl200683@tcp(127.0.0.1:3306)/ai_education?charset=utf8mb4&parseTime=True&loc=Local"
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
