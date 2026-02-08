package router

import (
	"gin-demo/internal/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Setup(r *gin.Engine){
	v1 := r.Group("/api")
	{
	v1.GET("/ppt",handler.ListPPT)
	v1.POST("/ppt",handler.CreatePPT)
	v1.GET("/ppt/:id",handler.GetPPT)
	v1.PUT("/ppt/:id",handler.UpdatePPT)
	}
}