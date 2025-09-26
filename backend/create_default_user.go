package main

import (
	"anti-fake-system/config"
	"anti-fake-system/database"
	"anti-fake-system/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateDefaultUser() {
	// 加载配置
	cfg := config.Load()

	// 初始化数据库
	db, err := database.InitDB(cfg)
	if err != nil {
		fmt.Printf("数据库初始化失败: %v\n", err)
		return
	}

	// 创建默认管理员用户
	createDefaultAdmin(db)
}

func createDefaultAdmin(db *gorm.DB) {
	// 检查是否已存在管理员用户
	var existingUser models.User
	if err := db.Where("username = ?", "admin").First(&existingUser).Error; err == nil {
		fmt.Printf("默认管理员用户已存在，用户名: admin\n")
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("密码加密失败: %v\n", err)
		return
	}

	// 创建管理员用户
	adminUser := models.User{
		Username: "admin",
		Password: string(hashedPassword),
		Email:    "admin@example.com",
		UserType: 1, // 平台用户
		Status:   1, // 启用
	}

	if err := db.Create(&adminUser).Error; err != nil {
		fmt.Printf("创建管理员用户失败: %v\n", err)
		return
	}

	fmt.Printf("默认管理员用户创建成功！\n")
	fmt.Printf("用户名: admin\n")
	fmt.Printf("密码: admin123\n")
	fmt.Printf("用户类型: 平台管理员\n")
}
