package router

import (
	"gin-demo/internal/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Setup(r *gin.Engine){
	r.GET("/ppt",handler.ListPPT)
	r.POST("/ppt",handler.CreatePPT)
	r.GET("/ppt:id",handler.GetPPT)
	r.PUT("/ppt:id",handler.UpdatePPT)
}