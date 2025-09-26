# FTP上传配置指南

## VS Code 配置 FTP 上传

### 方法1: 使用 SFTP 扩展 (推荐)

1. **安装扩展**
   - 打开 VS Code
   - 按 `Ctrl+Shift+X` 打开扩展面板
   - 搜索并安装 "SFTP" 扩展 (作者: Natizyskunk)

2. **配置 SFTP**
   - 按 `Ctrl+Shift+P` 打开命令面板
   - 输入 "SFTP: Config" 并选择
   - 会在项目根目录创建 `.vscode/sftp.json` 文件

3. **编辑配置文件**
```json
{
    "name": "防伪系统服务器",
    "host": "your_host",
    "protocol": "sftp",
    "port": 22,
    "username": "your_username",
    "password": "your_password",
    "remotePath": "/var/www/html/anti-fake-system",
    "uploadOnSave": true,
    "useTempFile": false,
    "openSsh": false,
    "ignore": [
        ".vscode",
        ".git",
        ".DS_Store",
        "node_modules",
        "*.log"
    ]
}
```

4. **使用方法**
   - 右键文件/文件夹 → "Upload" 上传
   - 右键文件/文件夹 → "Download" 下载
   - 设置 `"uploadOnSave": true` 可以保存时自动上传

### 方法2: 使用 FTP-Sync 扩展

1. **安装扩展**
   - 搜索并安装 "FTP-Sync" 扩展

2. **配置文件** (在项目根目录创建 `ftp-sync.json`)
```json
{
    "remotePath": "/var/www/html/anti-fake-system",
    "host": "your_host",
    "username": "your_username",
    "password": "your_password",
    "port": 21,
    "secure": false,
    "protocol": "ftp",
    "uploadOnSave": true,
    "passive": true,
    "debug": false,
    "privateKeyPath": null,
    "passphrase": null,
    "ignore": [
        ".vscode",
        ".git",
        ".DS_Store",
        "node_modules"
    ]
}
```

## WebStorm/PhpStorm 配置

1. **打开部署配置**
   - Tools → Deployment → Configuration

2. **添加服务器**
   - 点击 "+" 添加新服务器
   - 选择 SFTP 或 FTP

3. **配置连接**
   - Host: 8.138.244.31
   - Port: 22 (SFTP) 或 21 (FTP)
   - Username: your_username
   - Password: your_password
   - Root path: /var/www/html/anti-fake-system

4. **映射配置**
   - 在 "Mappings" 标签页设置本地路径和远程路径的映射

5. **自动上传**
   - Tools → Deployment → Automatic Upload (always)

## Sublime Text 配置

1. **安装 Package Control**
   - 按 `Ctrl+Shift+P`
   - 输入 "Install Package Control"

2. **安装 SFTP 插件**
   - `Ctrl+Shift+P` → "Package Control: Install Package"
   - 搜索并安装 "SFTP"

3. **配置 SFTP**
   - 右键项目文件夹 → "SFTP/FTP" → "Map to Remote..."
   - 编辑生成的 `sftp-config.json`

```json
{
    "type": "sftp",
    "host": "your_host",
    "user": "your_username",
    "password": "your_password",
    "port": "22",
    "remote_path": "/var/www/html/anti-fake-system",
    "upload_on_save": true,
    "ignore_regexes": [
        "\\.sublime-(project|workspace)",
        "sftp-config(-alt\\d?)?\\.json",
        "sftp-settings\\.json",
        "/venv/",
        "\\.svn/",
        "\\.hg/",
        "\\.git/",
        "\\.bzr",
        "_darcs",
        "CVS",
        "\\.DS_Store",
        "Thumbs\\.db",
        "desktop\\.ini"
    ]
}
```

## 通用 FTP 客户端推荐

### FileZilla (免费)
1. 下载安装 FileZilla
2. 站点管理器配置：
   - 主机: your_host
   - 端口: 21 (FTP) 或 22 (SFTP)
   - 协议: FTP 或 SFTP
   - 用户名和密码

### WinSCP (Windows, 免费)
1. 下载安装 WinSCP
2. 新建会话：
   - 文件协议: SFTP 或 FTP
   - 主机名: your_host
   - 端口: 22 或 21
   - 用户名和密码

## 安全建议

1. **使用 SFTP 而不是 FTP**
   - SFTP 加密传输，更安全
   - 端口通常是 22

2. **使用密钥认证**
```json
{
    "privateKeyPath": "/path/to/your/private/key",
    "passphrase": "your_passphrase"
}
```

3. **设置忽略文件**
   - 避免上传敏感文件 (.env, .git, node_modules等)

## 项目结构上传建议

```
服务器目录结构:
/var/www/html/anti-fake-system/
├── frontend/          # Vue前端文件
│   ├── dist/         # 构建后的文件
│   └── src/          # 源码
├── backend/          # Go后端文件
│   ├── main.go
│   └── ...
└── README.md
```

## 自动化部署脚本

创建 `deploy.sh` 脚本：
```bash
#!/bin/bash
# 前端构建
cd frontend
npm run build

# 上传到服务器
rsync -avz --delete dist/ user@8.138.244.31:/var/www/html/frontend/
rsync -avz --delete ../backend/ user@8.138.244.31:/var/www/html/backend/
```

选择适合您编辑器的配置方法，如果需要具体的服务器用户名和密码信息，请提供服务器的SSH/FTP访问凭据。