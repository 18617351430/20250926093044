-- 防伪系统初始数据
-- 创建默认平台管理员用户
INSERT INTO users (username, password, email, user_type, status, created_at, updated_at) VALUES 
('admin', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'admin@antifake.com', 1, 1, NOW(), NOW());
-- 密码是: password

-- 创建默认商户
INSERT INTO merchants (name, contact_person, phone, email, address, status, created_at, updated_at) VALUES 
('测试商户', '张三', '13800138000', 'test@merchant.com', '北京市朝阳区测试地址', 1, NOW(), NOW());

-- 创建默认商户管理员用户
INSERT INTO users (username, password, email, user_type, merchant_id, status, created_at, updated_at) VALUES 
('merchant', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'merchant@test.com', 2, 1, 1, NOW(), NOW());
-- 密码是: password

-- 创建默认商品
INSERT INTO products (merchant_id, name, description, category, price, status, created_at, updated_at) VALUES 
(1, '测试商品', '这是一个测试商品', '电子产品', 99.99, 1, NOW(), NOW());

-- 创建默认批次
INSERT INTO product_batches (merchant_id, product_id, batch_code, production_date, expire_date, quantity, status, created_at, updated_at) VALUES 
(1, 1, 'BATCH001', '2024-01-01', '2025-12-31', 1000, 1, NOW(), NOW());