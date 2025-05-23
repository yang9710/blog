# Story-1
# 实现用户认证功能

**As a** 博客系统用户
**I want** 能够注册和登录系统
**so that** 我可以使用系统的所有功能

## Status

Approved

## Context

这是博客系统的第一个用户故事，实现基础的认证功能。根据架构文档，我们将使用 Go 语言和 Gin 框架实现后端 API，使用 JWT 进行身份认证。前端将使用 Next.js 实现用户界面。数据库采用 MySQL，使用 GORM 作为 ORM 框架。

## Security Issues

1. 登录请求中的密码以明文形式传输
   - 问题描述：在浏览器开发者工具的 Network 面板中可以看到登录请求中的密码是明文
   - 风险等级：高
   - 解决方案：
     - 前端使用 HTTPS 加密传输
     - 使用加密算法对密码进行加密后再传输
     - 后端验证加密后的密码

## Estimation

Story Points: 3

## Tasks

1. - [ ] 后端开发
   1. - [ ] 创建用户模型
   2. - [ ] 实现用户注册 API
   3. - [ ] 实现用户登录 API
   4. - [ ] 实现 JWT 认证中间件
   5. - [ ] 编写单元测试

2. - [ ] 前端开发
   1. - [ ] 创建注册页面
   2. - [ ] 创建登录页面
   3. - [ ] 实现表单验证
   4. - [ ] 实现 API 调用
   5. - [ ] 实现 token 存储

3. - [ ] 数据库开发
   1. - [ ] 创建用户表
   2. - [ ] 实现密码加密
   3. - [ ] 添加必要索引
   4. - [ ] 配置 MySQL 连接池

4. - [ ] 测试
   1. - [ ] 后端 API 测试
   2. - [ ] 前端功能测试
   3. - [ ] 集成测试
   4. - [ ] 性能测试
   5. - [ ] MySQL 性能测试

## Constraints

- 密码必须加密存储
- 使用 JWT 进行身份认证
- 前端表单必须进行验证
- 必须实现错误处理
- 必须记录日志
- 使用 MySQL 事务确保数据一致性
- 实现数据库连接池优化

## Data Models / Schema

### API 接口设计

1. 用户注册
```
POST /api/v1/auth/register
Request:
{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
}

Response (201 Created):
{
    "code": 201,
    "message": "注册成功",
    "data": {
        "id": 1,
        "username": "testuser",
        "email": "test@example.com",
        "created_at": "2024-04-13T12:00:00Z",
        "updated_at": "2024-04-13T12:00:00Z"
    }
}

Response (400 Bad Request):
{
    "code": 400,
    "message": "注册失败: email already exists",
    "data": null
}
```

2. 用户登录
```
POST /api/v1/auth/login
Request:
{
    "email": "test@example.com",
    "password": "password123"
}

Response (200 OK):
{
    "code": 200,
    "message": "登录成功",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
        "user": {
            "id": 1,
            "username": "testuser",
            "email": "test@example.com"
        }
    }
}

Response (401 Unauthorized):
{
    "code": 401,
    "message": "登录失败: invalid credentials",
    "data": null
}
```

### MySQL 用户表结构
```sql
CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    role VARCHAR(20) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email),
    INDEX idx_username (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

## Structure

```
├── /backend
│   ├── /internal
│   │   ├── /handler
│   │   │   └── auth.go
│   │   ├── /model
│   │   │   └── user.go
│   │   ├── /service
│   │   │   └── auth.go
│   │   └── /middleware
│   │       └── auth.go
│   ├── /config
│   │   └── database.go
├── /frontend
│   ├── /app
│   │   └── /auth
│   │       ├── /login
│   │       │   └── page.tsx
│   │       └── /register
│   │           └── page.tsx
│   ├── /components
│   │   └── /auth
│   │       ├── RegisterForm.tsx
│   │       └── LoginForm.tsx
│   ├── /services
│   │   ├── index.ts
│   │   ├── request.ts
│   │   └── url.ts

```

## Diagrams

### 用户注册流程
```mermaid
sequenceDiagram
    participant U as 用户
    participant F as 前端
    participant B as 后端
    participant D as MySQL

    U->>F: 填写注册表单
    F->>B: POST /api/v1/auth/register
    B->>D: 开始事务
    B->>D: 检查用户名和邮箱
    D-->>B: 返回检查结果
    B->>D: 创建用户记录
    D-->>B: 返回用户信息
    B->>D: 提交事务
    B-->>F: 返回成功响应
    F-->>U: 显示成功消息
```

### 用户登录流程
```mermaid
sequenceDiagram
    participant U as 用户
    participant F as 前端
    participant B as 后端
    participant D as MySQL

    U->>F: 填写登录表单
    F->>B: POST /api/v1/auth/login
    B->>D: 查询用户信息
    D-->>B: 返回用户数据
    B->>B: 验证密码
    B->>B: 生成 JWT token
    B-->>F: 返回 token 和用户信息
    F->>F: 存储 token
    F-->>U: 跳转到首页
```

## Dev Notes

- 使用 bcrypt 进行密码加密
- 使用 JWT 进行身份认证
- 实现 token 刷新机制
- 添加请求限流
- 实现错误处理中间件
- 配置 MySQL 连接池参数
- 使用事务确保数据一致性
- 添加适当的数据库索引

## Chat Command Log

- 开始实现用户认证功能
- 修改数据库为 MySQL