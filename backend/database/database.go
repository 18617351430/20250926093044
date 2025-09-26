package database

import (
	"anti-fake-system/config" // 导入项目配置包
	"context"                 // 导入context包，用于管理请求的生命周期
	"fmt"                     // 导入fmt包，用于格式化字符串
	"log"                     // 导入log包，用于日志输出
	"time"                    // 导入time包，用于时间相关的操作

	"github.com/redis/go-redis/v9" // 导入go-redis客户端库
	"gorm.io/driver/mysql"         // 导入GORM的MySQL驱动
	"gorm.io/gorm"                 // 导入GORM ORM库
)

// InitDB 函数用于初始化MySQL数据库连接。
// 它接收一个Config结构体指针，其中包含数据库连接信息。
// 返回一个*gorm.DB实例和可能发生的错误。
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	// 构建MySQL的DSN (Data Source Name) 字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,     // 数据库用户名
		cfg.Database.Password, // 数据库密码
		cfg.Database.Host,     // 数据库主机
		cfg.Database.Port,     // 数据库端口
		cfg.Database.Name,     // 数据库名称
	)

	// 使用GORM打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败: %v", err) // 如果连接失败，返回错误
	}

	// 获取通用数据库对象 sql.DB，用于配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置数据库连接池参数
	sqlDB.SetMaxIdleConns(10)    // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)   // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可复用的最长时间

	log.Println("数据库连接成功") // 记录数据库连接成功日志
	return db, nil             // 返回GORM数据库实例
}

// InitRedis 函数用于初始化Redis连接。
// 它接收一个Config结构体指针，其中包含Redis连接信息。
// 返回一个*redis.Client实例和可能发生的错误。
func InitRedis(cfg *config.Config) (*redis.Client, error) {
	// 使用go-redis库创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,     // Redis服务器地址
		Password: cfg.Redis.Password, // Redis密码
		DB:       cfg.Redis.DB,       // Redis数据库索引
	})

	// 测试Redis连接是否成功
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("Redis连接失败: %v", err) // 如果连接失败，返回错误
	}

	log.Println("Redis连接成功") // 记录Redis连接成功日志
	return client, nil         // 返回Redis客户端实例
}
