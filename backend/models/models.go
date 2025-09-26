package models

import (
	"time" // 导入time包，用于处理时间字段

	"gorm.io/gorm" // 导入GORM库，用于数据库操作和模型定义
)

// Merchant 结构体定义了商户表的数据模型。
// 对应数据库中的 `merchants` 表。
type Merchant struct {
	ID        uint      `gorm:"primaryKey"`                  // 主键ID，无符号整型
	Code      string    `gorm:"size:6;uniqueIndex;not null"` // 商户标识码，长度6，唯一索引，非空
	Name      string    `gorm:"size:100;not null"`           // 商户名称，长度100，非空
	Status    int       `gorm:"default:1"`                   // 商户状态：1-启用, 0-禁用，默认1
	Config    string    `gorm:"type:json"`                   // 商户配置参数，以JSON字符串形式存储
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}

// User 结构体定义了用户表的数据模型。
// 对应数据库中的 `users` 表。
type User struct {
	ID          uint       `gorm:"primaryKey"`                   // 主键ID
	Username    string     `gorm:"size:50;uniqueIndex;not null"` // 用户名，长度50，唯一索引，非空
	Password    string     `gorm:"size:255;not null"`            // 密码（通常是哈希值），长度255，非空
	Email       string     `gorm:"size:100"`                     // 邮箱，长度100
	Phone       string     `gorm:"size:20"`                      // 手机号，长度20
	UserType    int        `gorm:"not null"`                     // 用户类型：1-平台用户, 2-商户用户，非空
	MerchantID  *uint      // 商户用户关联的商户ID，可为空
	Status      int        `gorm:"default:1"` // 用户状态：1-启用, 0-禁用，默认1
	LastLoginAt *time.Time // 最后登录时间，可为空
	CreatedAt   time.Time  // 创建时间
	UpdatedAt   time.Time  // 更新时间

	Merchant Merchant `gorm:"foreignKey:MerchantID"` // 关联的商户信息，通过MerchantID外键关联
}

// Role 结构体定义了角色表的数据模型。
// 对应数据库中的 `roles` 表。
type Role struct {
	ID          uint      `gorm:"primaryKey"`       // 主键ID
	Name        string    `gorm:"size:50;not null"` // 角色名称，长度50，非空
	RoleType    int       `gorm:"not null"`         // 角色类型：1-平台级角色, 2-商户级角色，非空
	MerchantID  *uint     // 商户级角色关联的商户ID，可为空
	Description string    `gorm:"size:200"` // 角色描述，长度200
	CreatedAt   time.Time // 创建时间
	UpdatedAt   time.Time // 更新时间

	Merchant    Merchant     `gorm:"foreignKey:MerchantID"`       // 关联的商户信息，通过MerchantID外键关联
	Permissions []Permission `gorm:"many2many:role_permissions;"` // 角色拥有的权限，多对多关系
}

// Permission 结构体定义了权限表的数据模型。
// 对应数据库中的 `permissions` 表。
type Permission struct {
	ID        uint      `gorm:"primaryKey"`                    // 主键ID
	Name      string    `gorm:"size:50;not null"`              // 权限名称，长度50，非空
	Code      string    `gorm:"size:100;uniqueIndex;not null"` // 权限标识码，长度100，唯一索引，非空
	Type      int       `gorm:"not null"`                      // 权限类型：1-菜单权限, 2-按钮权限, 3-数据权限，非空
	ParentID  *uint     // 父权限ID，用于构建权限树，可为空
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}

// UserRole 结构体定义了用户与角色关联表的数据模型。
// 对应数据库中的 `user_roles` 表。
type UserRole struct {
	ID        uint      `gorm:"primaryKey"` // 主键ID
	UserID    uint      `gorm:"not null"`   // 用户ID，非空
	RoleID    uint      `gorm:"not null"`   // 角色ID，非空
	CreatedAt time.Time // 创建时间

	User User `gorm:"foreignKey:UserID"` // 关联的用户信息
	Role Role `gorm:"foreignKey:RoleID"` // 关联的角色信息
}

// SecurityCodeRule 结构体定义了防伪码规则表的数据模型。
// 对应数据库中的 `security_code_rules` 表。
type SecurityCodeRule struct {
	ID         uint      `gorm:"primaryKey"`         // 主键ID
	MerchantID uint      `gorm:"not null"`           // 商户ID，非空
	Name       string    `gorm:"size:100;not null"`  // 规则名称，长度100，非空
	RuleConfig string    `gorm:"type:json;not null"` // 规则配置，以JSON字符串形式存储，非空
	Status     int       `gorm:"default:1"`          // 规则状态：1-启用, 0-禁用，默认1
	CreatedAt  time.Time // 创建时间
	UpdatedAt  time.Time // 更新时间

	Merchant Merchant `gorm:"foreignKey:MerchantID"` // 关联的商户信息
}

// Product 结构体定义了商品表的数据模型。
// 对应数据库中的 `products` 表。
type Product struct {
	ID          uint      `gorm:"primaryKey"`        // 主键ID
	MerchantID  uint      `gorm:"not null"`          // 商户ID，非空
	Name        string    `gorm:"size:200;not null"` // 商品名称，长度200，非空
	Category    string    `gorm:"size:100"`          // 商品类别，长度100
	Description string    `gorm:"type:text"`         // 商品描述，长文本
	Images      string    `gorm:"type:json"`         // 商品图片URL数组，以JSON字符串形式存储
	Status      int       `gorm:"default:1"`         // 商品状态：1-启用, 0-禁用，默认1
	CreatedAt   time.Time // 创建时间
	UpdatedAt   time.Time // 更新时间

	Merchant Merchant `gorm:"foreignKey:MerchantID"` // 关联的商户信息
}

// ProductBatch 结构体定义了商品批次表的数据模型。
// 对应数据库中的 `product_batches` 表。
type ProductBatch struct {
	ID             uint      `gorm:"primaryKey"`      // 主键ID
	ProductID      uint      `gorm:"not null"`        // 商品ID，非空
	MerchantID     uint      `gorm:"not null"`        // 商户ID，非空
	BatchCode      string    `gorm:"size:8;not null"` // 批次标识码，长度8，非空
	ProductionDate time.Time // 生产日期
	Quantity       int       `gorm:"not null"`  // 批次数量，非空
	Status         int       `gorm:"default:1"` // 批次状态：1-启用, 0-禁用，默认1
	CreatedAt      time.Time // 创建时间
	UpdatedAt      time.Time // 更新时间

	Product  Product  `gorm:"foreignKey:ProductID"`  // 关联的商品信息
	Merchant Merchant `gorm:"foreignKey:MerchantID"` // 关联的商户信息
}

// TraceabilityInfo 结构体定义了溯源信息表的数据模型。
// 对应数据库中的 `traceability_infos` 表。
type TraceabilityInfo struct {
	ID          uint      `gorm:"primaryKey"`         // 主键ID
	BatchID     uint      `gorm:"not null"`           // 商品批次ID，非空
	MerchantID  uint      `gorm:"not null"`           // 商户ID，非空
	StageType   int       `gorm:"not null"`           // 溯源阶段类型：1-生产, 2-仓储, 3-物流, 4-销售，非空
	Content     string    `gorm:"type:json;not null"` // 溯源信息内容，以JSON字符串形式存储，非空
	Attachments string    `gorm:"type:json"`          // 附件信息，以JSON字符串形式存储
	CreatedAt   time.Time // 创建时间
	UpdatedAt   time.Time // 更新时间

	Batch    ProductBatch `gorm:"foreignKey:BatchID"`    // 关联的商品批次信息
	Merchant Merchant     `gorm:"foreignKey:MerchantID"` // 关联的商户信息
}

// VerificationRecord 结构体定义了防伪验证记录表的数据模型。
// 对应数据库中的 `verification_records` 表。
type VerificationRecord struct {
	ID           uint      `gorm:"primaryKey"`       // 主键ID
	SecurityCode string    `gorm:"size:32;not null"` // 防伪码，长度32，非空
	MerchantID   uint      `gorm:"not null"`         // 商户ID，非空
	VerifyTime   time.Time `gorm:"not null"`         // 验证时间，非空
	IPAddress    string    `gorm:"size:45"`          // 验证IP地址，长度45
	Result       int       `gorm:"not null"`         // 验证结果：1-真品, 0-伪品, 2-无效码，非空
	UserAgent    string    `gorm:"size:500"`         // 用户代理（浏览器信息），长度500
	CreatedAt    time.Time // 创建时间

	Merchant Merchant `gorm:"foreignKey:MerchantID"` // 关联的商户信息
}

// AutoMigrate 函数用于自动迁移数据库表结构。
// 它接收一个GORM数据库实例，并根据定义的模型创建或更新数据库表。
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Merchant{},           // 迁移商户表
		&User{},               // 迁移用户表
		&Role{},               // 迁移角色表
		&Permission{},         // 迁移权限表
		&UserRole{},           // 迁移用户角色关联表
		&SecurityCodeRule{},   // 迁移防伪码规则表
		&Product{},            // 迁移商品表
		&ProductBatch{},       // 迁移商品批次表
		&TraceabilityInfo{},   // 迁移溯源信息表
		&VerificationRecord{}, // 迁移验证记录表
	)
}
