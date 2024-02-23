package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// findEnvFile 向上遍历父级目录查找.env文件
func findEnvFile(startDir string) (string, bool) {
	dir := startDir
	for {
		// 检查当前目录下是否有.env文件
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath, true
		}

		// 到达文件系统的根目录时停止
		if dir == "/" {
			break
		}

		// 移动到上一级目录
		dir = filepath.Dir(dir)
	}

	return "", false
}

// InitEnv 从当前目录向上查找.env文件并加载
func InitEnv() {
	// 获取当前执行文件的目录
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory: ", err)
	}

	// 查找.env文件
	envPath, found := findEnvFile(currentDir)
	if !found {
		log.Panic(".env file not found in any parent directories")
	}

	// 加载找到的.env文件
	if err := godotenv.Load(envPath); err != nil {
		log.Panic("Error loading .env file: ", err)
	}
}
