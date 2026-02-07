package main

import (
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
	dsn := "root:lyh985211@tcp(127.0.0.1:3306)/ai_education?charset=utf8mb4&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil{
		log.Fatal("数据库连接失败",err)
	}
	db.AutoMigrate(&model.PPT{})
	router.DB = db
	
	r := gin.Default()
	router.Setup(r)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("服务启动失败", err)
	}
}
