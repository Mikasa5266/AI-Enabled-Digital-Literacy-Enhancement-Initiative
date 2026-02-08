package handler

import (
	"fmt"
	"gin-demo/internal/model"
	"gin-demo/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
}

func ListPPT(c *gin.Context) {
	var ppts []model.PPT
	if err := db.Find(&ppts).Error; err != nil {
		fmt.Println("数据库查询错误:", err) 
		utils.Error(c, http.StatusInternalServerError, "查询ppt失败")
		return
	}
	utils.Success(c, ppts)
}
func CreatePPT(c *gin.Context) {
	var ppt model.PPT
	if err := c.ShouldBindJSON(&ppt); err != nil {
		utils.Error(c, http.StatusBadRequest, "请求参数错误")
		return
	}
	if err := db.Create(&ppt).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "创建PPT失败")
		return
	}
	utils.Success(c, ppt)
}
func GetPPT(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}
	var ppt model.PPT
	if err := db.First(&ppt, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "PPT未找到")
		return
	}
	utils.Success(c, ppt)
}
func UpdatePPT(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}
	var ppt model.PPT
	if err := c.ShouldBindJSON(&ppt); err != nil {
		utils.Error(c, http.StatusBadRequest, "请求参数错误")
		return
	}
	if err := db.Model(&model.PPT{}).Where("id = ?", id).Updates(ppt).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "更新PPT失败")
		return
	}
	utils.Success(c,ppt)
}
func deletePPT(c *gin.Context){
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		utils.Error(c,http.StatusBadRequest,"无效的ID")
		return
	}
	if err := db.Delete(&model.PPT{},id);err != nil{
		utils.Error(c,http.StatusInternalServerError,"删除PPT失败")
		return
	}
	utils.Success(c,gin.H{"message":"PPT已删除"})
}
