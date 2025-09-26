package models

import (
    "time"
    "gorm.io/gorm"
)

// 商户表
type Merchant struct {
    ID          uint      `gorm:"primaryKey"`
    Code        string    `gorm:"size:6;uniqueIndex;not null"` // 商户标识码(2-6位)
    Name        string    `gorm:"size:100;not null"`
    Status      int       `gorm:"default:1"` // 1-启用, 0-禁用
    Config      string    `gorm:"type:json"` // 商户配置参数
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// 用户表
type User struct {
    ID          uint      `gorm:"primaryKey"`
    Username    string    `gorm:"size:50;uniqueIndex;not null"`
    Password    string    `gorm:"size:255;not null"`
    Email       string    `gorm:"size:100"`
    Phone       string    `gorm:"size:20"`
    UserType    int       `gorm:"not null"` // 1-平台用户, 2-商户用户
    MerchantID  *uint     // 商户用户关联的商户ID
    Status      int       `gorm:"default:1"` // 1-启用, 0-禁用
    LastLoginAt *time.Time
    CreatedAt   time.Time
    UpdatedAt   time.Time
    
    Merchant    Merchant `gorm:"foreignKey:MerchantID"`
}

// 角色表
type Role struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string    `gorm:"size:50;not null"`
    RoleType    int       `gorm:"not null"` // 1-平台级角色, 2-商户级角色
    MerchantID  *uint     // 商户级角色关联的商户ID
    Description string    `gorm:"size:200"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    
    Merchant    Merchant `gorm:"foreignKey:MerchantID"`
    Permissions []Permission `gorm:"many2many:role_permissions;"`
}

// 权限表
type Permission struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string    `gorm:"size:50;not null"`
    Code        string    `gorm:"size:100;uniqueIndex;not null"` // 权限标识
    Type        int       `gorm:"not null"` // 1-菜单权限, 2-按钮权限, 3-数据权限
    ParentID    *uint
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// 用户角色关联表
type UserRole struct {
    ID     uint `gorm:"primaryKey"`
    UserID uint `gorm:"not null"`
    RoleID uint `gorm:"not null"`
    CreatedAt time.Time
    
    User User `gorm:"foreignKey:UserID"`
    Role Role `gorm:"foreignKey:RoleID"`
}

// 防伪码规则表
type SecurityCodeRule struct {
    ID          uint      `gorm:"primaryKey"`
    MerchantID  uint      `gorm:"not null"`
    Name        string    `gorm:"size:100;not null"`
    RuleConfig  string    `gorm:"type:json;not null"` // 规则配置(加密存储)
    Status      int       `gorm:"default:1"` // 1-启用, 0-禁用
    CreatedAt   time.Time
    UpdatedAt   time.Time
    
    Merchant    Merchant `gorm:"foreignKey:MerchantID"`
}

// 商品表
type Product struct {
    ID          uint      `gorm:"primaryKey"`
    MerchantID  uint      `gorm:"not null"`
    Name        string    `gorm:"size:200;not null"`
    Category    string    `gorm:"size:100"`
    Description string    `gorm:"type:text"`
    Images      string    `gorm:"type:json"` // 商品图片数组
    Status      int       `gorm:"default:1"` // 1-启用, 0-禁用
    CreatedAt   time.Time
    UpdatedAt   time.Time
    
    Merchant    Merchant `gorm:"foreignKey:MerchantID"`
}

// 商品批次表
type ProductBatch struct {
    ID          uint      `gorm:"primaryKey"`
    ProductID   uint      `gorm:"not null"`
    MerchantID  uint      `gorm:"not null"`
    BatchCode   string    `gorm:"size:8;not null"` // 批次标识(2-8位)
    ProductionDate time.Time
    Quantity    int       `gorm:"not null"`
    Status      int       `gorm:"default:1"` // 1-启用, 0-禁用
    CreatedAt   time.Time
    UpdatedAt   time.Time
    
    Product     Product  `gorm:"foreignKey:ProductID"`
    Merchant    Merchant `gorm:"foreignKey:MerchantID"`
}

// 溯源信息表
type TraceabilityInfo struct {
    ID          uint      `gorm:"primaryKey"`
    BatchID     uint      `gorm:"not null"`
    MerchantID  uint      `gorm:"not null"`
    StageType   int       `gorm:"not null"` // 1-生产, 2-仓储, 3-物流, 4-销售
    Content     string    `gorm:"type:json;not null"` // 溯源信息内容
    Attachments string    `gorm:"type:json"` // 附件信息
    CreatedAt   time.Time
    UpdatedAt   time.Time
    
    Batch       ProductBatch `gorm:"foreignKey:BatchID"`
    Merchant    Merchant     `gorm:"foreignKey:MerchantID"`
}

// 验证记录表
type VerificationRecord struct {
    ID          uint      `gorm:"primaryKey"`
    SecurityCode string   `gorm:"size:32;not null"`
    MerchantID  uint      `gorm:"not null"`
    VerifyTime  time.Time `gorm:"not null"`
    IPAddress   string    `gorm:"size:45"`
    Result      int       `gorm:"not null"` // 1-真品, 0-伪品, 2-无效码
    UserAgent   string    `gorm:"size:500"`
    CreatedAt   time.Time
    
    Merchant    Merchant `gorm:"foreignKey:MerchantID"`
}

// 初始化数据库表
func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &Merchant{},
        &User{},
        &Role{},
        &Permission{},
        &UserRole{},
        &SecurityCodeRule{},
        &Product{},
        &ProductBatch{},
        &TraceabilityInfo{},
        &VerificationRecord{},
    )
}