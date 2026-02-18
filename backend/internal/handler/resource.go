package handler

import (
	"fmt"
	"gin-demo/internal/model"
	"gin-demo/pkg/utils"
	"net/http"
	"os"
	"path/filepath"
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
	fmt.Println("收到的 title:", c.PostForm("title"))
	fmt.Println("收到的 category:", c.PostForm("category"))
	// 手动获取表单字段
	ppt.Title = c.PostForm("title")
	ppt.Category = c.PostForm("category")
	ppt.Description = c.PostForm("description")
	// 验证必填字段
	if ppt.Title == "" {
		utils.Error(c, http.StatusBadRequest, "标题为必填项")
		return
	}

	// 先创建数据库记录获取 ID
	if err := db.Create(&ppt).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "创建PPT失败")
		return
	}
	//设置统一路径
	baseStoragePath := "./storage"
	pptSaveDir := filepath.Join(baseStoragePath, "ppt")
	coverSaveDir := filepath.Join(baseStoragePath, "covers")
	// 自动创建目录 (如果不存在)
	os.MkdirAll(pptSaveDir, 0755)
	os.MkdirAll(coverSaveDir, 0755)
	// 处理封面文件
	coverFile, err := c.FormFile("coverFile")
	//获取封面文件名
	coverFileName := ""
	//获取封面文件完整路径
	if err == nil && coverFile != nil {
		coverFileName = utils.GenerateUniqueFilename(coverFile.Filename)
		coverFilePath := filepath.Join(coverSaveDir, coverFileName)
		ppt.CoverURL = coverFilePath
		if err := c.SaveUploadedFile(coverFile, coverFilePath); err != nil {
			fmt.Println("保存封面失败:", err)
		}
	}

	// 处理 PPT 文件
	pptFile, err := c.FormFile("pptFile")
	pptFileName := ""
	if err == nil && pptFile != nil {
		//获取ppt文件名
		pptFileName = utils.GenerateUniqueFilename(pptFile.Filename)
		//获取ppt文件完整路径
		pptFilePath := filepath.Join(pptSaveDir, pptFileName)
		ppt.FileURL = pptFilePath
		if err := c.SaveUploadedFile(pptFile, pptFilePath); err != nil {
			fmt.Println("保存PPT失败:", err)
		} 
	}
	fmt.Println(ppt.CoverURL)
	// 更新文件路径到数据库
	db.Save(&ppt)

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
	utils.Success(c, ppt)
}
func deletePPT(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}
	if err := db.Delete(&model.PPT{}, id); err != nil {
		utils.Error(c, http.StatusInternalServerError, "删除PPT失败")
		return
	}
	utils.Success(c, gin.H{"message": "PPT已删除"})
}
