package config

import (
	"os"      // 导入os包，用于获取环境变量
	"strconv" // 导入strconv包，用于字符串和数字之间的转换
)

// Config 结构体定义了整个应用程序的配置信息，聚合了服务器、数据库、Redis和JWT的配置。
type Config struct {
	Server   ServerConfig   // 服务器配置
	Database DatabaseConfig // 数据库配置
	Redis    RedisConfig    // Redis缓存配置
	JWT      JWTConfig      // JWT认证配置
}

// ServerConfig 结构体定义了服务器相关的配置，如端口和运行模式。
type ServerConfig struct {
	Port string // 服务器监听端口
	Mode string // 服务器运行模式 (e.g., "debug", "release")
}

// DatabaseConfig 结构体定义了数据库连接相关的配置。
type DatabaseConfig struct {
	Host     string // 数据库主机地址
	Port     string // 数据库端口
	User     string // 数据库用户名
	Password string // 数据库密码
	Name     string // 数据库名称
}

// RedisConfig 结构体定义了Redis连接相关的配置。
type RedisConfig struct {
	Addr     string // Redis服务器地址 (e.g., "localhost:6379")
	Password string // Redis密码
	DB       int    // Redis数据库索引
}

// JWTConfig 结构体定义了JWT认证相关的配置。
type JWTConfig struct {
	Secret string // JWT签名密钥
	Expire int    // JWT过期时间 (小时)
}

// Load 函数用于从环境变量或使用默认值加载所有配置。
// 返回一个指向Config结构体的指针。
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),  // 从环境变量SERVER_PORT获取端口，默认8080
			Mode: getEnv("SERVER_MODE", "debug"), // 从环境变量SERVER_MODE获取模式，默认debug
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),        // 从环境变量DB_HOST获取数据库主机，默认localhost
			Port:     getEnv("DB_PORT", "3306"),             // 从环境变量DB_PORT获取数据库端口，默认3306
			User:     getEnv("DB_USER", "root"),             // 从环境变量DB_USER获取数据库用户，默认root
			Password: getEnv("DB_PASSWORD", ""),             // 从环境变量DB_PASSWORD获取数据库密码，默认空
			Name:     getEnv("DB_NAME", "anti_fake_system"), // 从环境变量DB_NAME获取数据库名，默认anti_fake_system
		},
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"), // 从环境变量REDIS_ADDR获取Redis地址，默认localhost:6379
			Password: getEnv("REDIS_PASSWORD", ""),           // 从环境变量REDIS_PASSWORD获取Redis密码，默认空
			DB:       getEnvInt("REDIS_DB", 0),               // 从环境变量REDIS_DB获取Redis数据库索引，默认0

		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "anti-fake-system-secret"), // 从环境变量JWT_SECRET获取JWT密钥，默认"anti-fake-system-secret"
			Expire: getEnvInt("JWT_EXPIRE", 24),                     // 从环境变量JWT_EXPIRE获取JWT过期时间，默认24小时
		},
	}
}

// getEnv 是一个辅助函数，用于从环境变量中获取指定键的值。
// 如果环境变量不存在或为空，则返回提供的默认值。
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvInt 是一个辅助函数，用于从环境变量中获取指定键的整数值。
// 如果环境变量不存在、为空或无法转换为整数，则返回提供的默认值。
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
