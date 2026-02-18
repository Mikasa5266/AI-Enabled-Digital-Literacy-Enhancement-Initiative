package utils

import (
	"math/rand"
	"fmt"
	"path/filepath"
	"time"
)

func GenerateUniqueFilename(originalName string) string {
	ext := filepath.Ext(originalName)  //获取拓展名
	timestamp := time.Now().Unix()
	randomNum := rand.Intn(1000)
	return fmt.Sprintf("%d_%d%s",timestamp,randomNum,ext)
}