# Go URL Shortener  
一个基于**Golang+Gin+MySQL**实现的短链接服务，用于将长链接转换为短链接，并在访问时进行重定向和访问次数统计。  

---
## 功能特点
- 生成短链接
- 通过短链接重定向
- 统计短链接访问次数
- 数据持久化到MySQL数据库
- 支持后端服务独立部署
---
## 技术栈
- Golang
- Gin Web Framework
- MySQL
- GORM ORM
- Linux/云服务器环境

---
## 项目结构

```text
URLShortener/
├── API/                       # 接口层（HTTP 请求入口）
│   ├── createShortURLAPI.go   # 创建短链接接口
│   └── redirectAPI.go         # 短链接重定向接口
│
├── Config/                    # 配置相关
│   └── initDB.go              # 数据库初始化与连接
│
├── Models/                    # 数据模型层
│   └── url.go                 # URL 数据结构，与数据库表映射
│
├── Routes/                    # 路由定义
│   └── router.go              # API 路由注册与分发
│
├── Service/                   # 业务逻辑层
│   └── increasesVisitCount.go # 访问次数自增逻辑
│
├── Utils/                     # 工具函数
│   └── generateShortCode.go   # 短链接编码生成（Base62）
│
├── main.go                    # 程序入口，初始化并启动服务
├── go.mod                     # Go 模块依赖管理
└── README.md                  # 项目说明文档
```
---
## 部署与运行