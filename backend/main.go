package main

import (
	"log" // 导入log包，用于日志输出

<<<<<<< HEAD
	"anti-fake-system/config"   // 导入项目配置包
	"anti-fake-system/database" // 导入数据库初始化包
	"anti-fake-system/models"   // 导入数据模型包，用于数据库迁移
	"anti-fake-system/routes"   // 导入路由设置包

	"github.com/joho/godotenv" // 导入godotenv包，用于加载.env文件
=======
	"anti-fake-system/config"
	"anti-fake-system/database"
	"anti-fake-system/models"
	"anti-fake-system/routes"

	"github.com/joho/godotenv"
>>>>>>> ccdb57742d35244c5b94196a200ad37653fe0f0a
)

// main 函数是Go应用程序的入口点。
// 它负责初始化配置、数据库、Redis、设置路由并启动HTTP服务器。
func main() {
<<<<<<< HEAD
	// 尝试加载.env文件中的环境变量。
	// 如果文件不存在或加载失败，则打印警告信息，但程序会继续使用系统环境变量或代码中定义的默认值。
	if err := godotenv.Load(); err != nil {
		log.Println("未找到.env文件，使用默认配置或系统环境变量")
	}

	// 加载应用程序配置。
	// config.Load() 函数会从环境变量中读取配置，如果环境变量不存在，则使用预设的默认值。
=======
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		log.Println("未找到.env文件，使用默认配置")
	}

	// 加载配置
>>>>>>> ccdb57742d35244c5b94196a200ad37653fe0f0a
	cfg := config.Load()

	// 初始化数据库连接。
	// 调用database.InitDB函数，传入加载的配置，获取GORM数据库实例。
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatal("数据库初始化失败:", err) // 如果数据库初始化失败，记录致命错误并退出程序
	}

	// 自动迁移数据库表结构。
	// models.AutoMigrate函数会根据定义的模型自动创建或更新数据库表。
	if err := models.AutoMigrate(db); err != nil {
		log.Fatal("数据库迁移失败:", err) // 如果数据库迁移失败，记录致命错误并退出程序
	}

	// 初始化Redis连接。
	// 调用database.InitRedis函数，传入加载的配置，获取Redis客户端实例。
	redisClient, err := database.InitRedis(cfg)
	if err != nil {
		log.Fatal("Redis初始化失败:", err) // 如果Redis初始化失败，记录致命错误并退出程序
	}

	// 设置HTTP路由。
	// routes.SetupRouter函数会配置所有API路由，并注入数据库和Redis客户端以及配置信息。
	router := routes.SetupRouter(db, redisClient, cfg)

	// 启动HTTP服务器。
	// 服务器将监听配置中指定的端口。
	log.Printf("服务器启动在端口 %s", cfg.Server.Port) // 打印服务器启动信息
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("服务器启动失败:", err) // 如果服务器启动失败，记录致命错误并退出程序
	}
}
