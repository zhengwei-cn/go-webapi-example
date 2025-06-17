# WebAPI - Go RESTful API

基于 Gin、GORM 和 PostgreSQL 的标准 RESTful API 项目。

## 项目结构

```
webapi/
├── main.go                 # 应用入口点
├── go.mod                  # Go 模块文件
├── go.sum                  # Go 依赖锁定文件
├── .env                    # 环境变量配置
├── docker-compose.yml      # PostgreSQL 数据库容器配置
├── config/                 # 配置模块
│   └── config.go
├── database/               # 数据库连接和迁移
│   └── database.go
├── models/                 # 数据模型层
│   └── models.go
├── services/               # 业务逻辑层
│   ├── user_service.go
│   └── product_service.go
├── controllers/            # API 控制器层
│   ├── user_controller.go
│   └── product_controller.go
├── routes/                 # 路由配置
│   └── routes.go
└── docs/                   # Swagger 生成的文档（运行后生成）
```

## 功能特性

- ✅ 使用 Gin 框架构建 RESTful API
- ✅ 使用 GORM 进行数据库操作
- ✅ PostgreSQL 数据库支持
- ✅ 自动数据库迁移
- ✅ Swagger API 文档
- ✅ 分层架构（API层、服务层、模型层）
- ✅ 环境变量配置
- ✅ CORS 支持
- ✅ 数据验证

## API 端点

### 用户管理
- `POST /api/v1/users` - 创建用户
- `GET /api/v1/users` - 获取所有用户
- `GET /api/v1/users/:id` - 获取指定用户
- `PUT /api/v1/users/:id` - 更新用户
- `DELETE /api/v1/users/:id` - 删除用户

### 产品管理
- `POST /api/v1/products` - 创建产品
- `GET /api/v1/products` - 获取所有产品
- `GET /api/v1/products/:id` - 获取指定产品
- `PUT /api/v1/products/:id` - 更新产品
- `DELETE /api/v1/products/:id` - 删除产品

### 其他
- `GET /health` - 健康检查
- `GET /swagger/index.html` - Swagger API 文档

## 快速开始

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 启动 PostgreSQL 数据库

```bash
docker-compose up -d
```

### 3. 生成 Swagger 文档

```bash
# 安装 swag CLI
go install github.com/swaggo/swag/cmd/swag@latest

# 生成文档
swag init
```

### 4. 运行应用

```bash
go run main.go
```

应用将在 `http://localhost:8080` 启动。

## 环境变量

在 `.env` 文件中配置以下变量：

```
DATABASE_URL=postgres://postgres:password@localhost:5432/webapi?sslmode=disable
PORT=8080
ENVIRONMENT=development
```

## API 测试示例

### 创建用户

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三",
    "email": "zhangsan@example.com",
    "age": 25
  }'
```

### 创建产品

```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPhone 15",
    "description": "最新款苹果手机",
    "price": 6999.99,
    "stock": 100,
    "user_id": 1
  }'
```

### 获取所有用户

```bash
curl http://localhost:8080/api/v1/users
```

## 数据库架构

### Users 表
- id (主键)
- name (用户名)
- email (邮箱，唯一)
- age (年龄)
- created_at, updated_at, deleted_at (时间戳)

### Products 表
- id (主键)
- name (产品名称)
- description (产品描述)
- price (价格)
- stock (库存)
- user_id (关联用户外键)
- created_at, updated_at, deleted_at (时间戳)

## 技术栈

- **Go 1.24.4** - 编程语言
- **Gin** - HTTP Web 框架
- **GORM** - ORM 库
- **PostgreSQL** - 数据库
- **Swagger** - API 文档
- **Docker** - 容器化数据库
