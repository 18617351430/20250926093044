package main

import (
	"log"

	"anti-fake-system/config"
	"anti-fake-system/database"
	"anti-fake-system/models"
	"anti-fake-system/routes"

	"github.com/joho/godotenv"
)

func main() {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		log.Println("未找到.env文件，使用默认配置")
	}

	// 加载配置
	cfg := config.Load()

	// 初始化数据库
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatal("数据库初始化失败:", err)
	}

	// 自动迁移表结构
	if err := models.AutoMigrate(db); err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	// 初始化Redis
	redisClient, err := database.InitRedis(cfg)
	if err != nil {
		log.Fatal("Redis初始化失败:", err)
	}

	// 设置路由
	router := routes.SetupRouter(db, redisClient, cfg)

	// 启动服务器
	log.Printf("服务器启动在端口 %s", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
