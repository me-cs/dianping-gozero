# 🍔 DianPing Clone - Go 微服务实战项目

<div align="center">

![CI](https://github.com/me-cs/dianping-gozero/actions/workflows/ci.yml/badge.svg)
![Docker Build](https://github.com/me-cs/dianping-gozero/actions/workflows/docker.yml/badge.svg)
![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![go-zero](https://img.shields.io/badge/go--zero-1.9.4-7C3AED?style=flat)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)
![MySQL](https://img.shields.io/badge/MySQL-8.0-4479A1?style=flat&logo=mysql)
![Redis](https://img.shields.io/badge/Redis-7.0-DC382D?style=flat&logo=redis)
![License](https://img.shields.io/badge/License-MIT-green?style=flat)
![AI Powered](https://img.shields.io/badge/90%25_Built_by-Claude_AI-FF6B6B?style=flat)

**🎓 一个完整的 Go 微服务学习项目 | 黑马点评克隆版后端**

[English](./README_EN.md) | [快速开始](#-快速开始) | [在线文档](./docs) | [贡献指南](#-贡献指南)

**⭐ 如果这个项目对你有帮助，请给个 Star！**

</div>

---

## 📚 目录

- [项目简介](#项目简介)
- [技术栈](#技术栈)
- [架构设计](#架构设计)
- [数据库设计](#数据库设计)
- [环境准备（新手必读）](#环境准备新手必读)
- [快速开始](#快速开始)
- [部署指南](#部署指南)
- [开发指南](#开发指南)
- [调试与故障排查](#调试与故障排查)
- [学习路径](#学习路径)
- [贡献指南](#贡献指南)
- [法律声明](#法律声明)
- [致谢](#致谢)
- [支持项目](#支持项目)
- [联系我们](#联系我们)
- [路线图](#路线图)
- [项目统计](#项目统计)
- [许可证](#许可证)

---

## 🎯 项目简介

**DianPing Clone** 是一个基于 **go-zero** 微服务框架构建的黑马点评克隆版后端系统，专为学习和实践现代微服务架构而设计。

### ✨ 项目亮点

- 🤖 **AI 驱动开发**：90% 的代码由 **Claude AI** 协助完成，展示了 AI 在软件工程中的强大能力
- 🏗️ **企业级架构**：完整的微服务架构设计，包含 API 网关、RPC 服务、缓存层、消息队列等
- 📊 **可观测性**：集成 Jaeger 分布式追踪、Prometheus 监控、Grafana 可视化
- 🐳 **容器化部署**：完整的 Docker Compose 配置，一键启动所有服务
- 📖 **详细文档**：从零开始的完整教程，适合新手学习
- 🎓 **最佳实践**：遵循 Go 社区最佳实践和微服务设计模式

### 🎯 核心功能

- 👤 **用户系统**：手机验证码登录、用户信息管理、Session 管理
- 🏪 **商铺系统**：商铺信息查询、分类浏览、地理位置搜索
- 🎫 **优惠券系统**：普通优惠券、秒杀券（高并发场景）
- 📝 **笔记/博客系统**：发布探店笔记、点赞、评论
- 🛒 **订单系统**：优惠券下单、订单管理

### 🎓 适合人群

- Go 语言初学者，想要学习微服务架构
- 有一定基础，想要了解 go-zero 框架
- 想要学习分布式系统设计和高并发处理
- 对 AI 辅助编程感兴趣的开发者
- 准备面试，需要一个完整的项目经验

---

## 🛠️ 技术栈

### 后端技术

| 技术 | 版本 | 用途 |
|------|------|------|
| **Go** | 1.23+ | 主要编程语言 |
| **go-zero** | 1.9.4 | 微服务框架 |
| **gRPC** | 1.78.0 | RPC 通信协议 |
| **Protocol Buffers** | 3.0 | 服务接口定义 |

### 数据存储

| 技术 | 版本 | 用途 |
|------|------|------|
| **MySQL** | 8.0 | 主数据库 |
| **Redis** | 7.0 | 缓存、Session、分布式锁 |
| **etcd** | 3.5 | 服务发现与注册 |

### 监控与追踪

| 技术 | 版本 | 用途 |
|------|------|------|
| **Jaeger** | latest | 分布式链路追踪 |
| **Prometheus** | latest | 指标收集 |
| **Grafana** | latest | 可视化监控面板 |

### 基础设施

| 技术 | 版本 | 用途 |
|------|------|------|
| **Docker** | 20.10+ | 容器化 |
| **Docker Compose** | 2.0+ | 容器编排 |
| **Nginx** | latest | 反向代理（可选） |

### 开发工具

- **GoLand / VS Code**：推荐 IDE
- **Postman / Apifox**：API 测试
- **TablePlus / Navicat**：数据库管理
- **Claude AI**：代码生成和问题解决

---

## 🏗️ 架构设计

### 整体架构

```
┌─────────────────────────────────────────────────────────────────┐
│                         Frontend (Future)                        │
│                    Vue 3 / React / Next.js                       │
└──────────────────────────────┬──────────────────────────────────┘
                               │ HTTP/REST
                               ↓
┌─────────────────────────────────────────────────────────────────┐
│                        API Gateway (8081)                        │
│                         go-zero API                              │
├─────────────────────────────────────────────────────────────────┤
│  • 统一入口                                                       │
│  • 身份认证                                                       │
│  • 请求路由                                                       │
│  • 负载均衡                                                       │
└──────────────────────────────┬──────────────────────────────────┘
                               │ gRPC
                ┌──────────────┼──────────────┐
                ↓              ↓              ↓
    ┌──────────────┐  ┌──────────────┐  ┌──────────────┐
    │  User RPC    │  │  Shop RPC    │  │ Voucher RPC  │
    │   (8001)     │  │   (8002)     │  │   (8003)     │
    └──────┬───────┘  └──────┬───────┘  └──────┬───────┘
           │                 │                  │
    ┌──────────────┐  ┌──────────────┐
    │  Order RPC   │  │  Blog RPC    │
    │   (8004)     │  │   (8005)     │
    └──────┬───────┘  └──────┬───────┘
           │                 │
           └─────────┬───────┘
                     │
    ┌────────────────┴────────────────┐
    │                                  │
    ↓                                  ↓
┌─────────────┐              ┌──────────────┐
│   MySQL     │              │    Redis     │
│   (3306)    │              │    (6379)    │
├─────────────┤              ├──────────────┤
│ • 用户表     │              │ • 缓存        │
│ • 商铺表     │              │ • Session    │
│ • 订单表     │              │ • 分布式锁    │
│ • 优惠券表   │              │ • 排行榜      │
└─────────────┘              └──────────────┘

┌─────────────────────────────────────────────────────────────────┐
│                    Infrastructure Services                       │
├─────────────────┬─────────────────┬──────────────┬──────────────┤
│  etcd (2379)    │ Jaeger (16686)  │ Prometheus   │  Grafana     │
│  服务发现        │  链路追踪        │  (9090)      │  (3000)      │
│                 │                 │  监控指标     │  可视化       │
└─────────────────┴─────────────────┴──────────────┴──────────────┘
```

### 服务划分

#### 1. API Gateway (端口: 8081)
- **职责**：统一入口、鉴权、路由
- **特点**：RESTful API、JWT 认证
- **对外接口**：HTTP/JSON

#### 2. User RPC (端口: 8001)
- **职责**：用户注册、登录、信息管理
- **数据**：用户表、Session
- **缓存策略**：用户信息缓存、验证码存储

#### 3. Shop RPC (端口: 8002)
- **职责**：商铺 CRUD、分类管理、地理位置搜索
- **数据**：商铺表、商铺类型表
- **缓存策略**：热点商铺缓存、分类列表缓存

#### 4. Voucher RPC (端口: 8003)
- **职责**：优惠券管理、秒杀券库存控制
- **数据**：优惠券表、秒杀券表
- **高并发方案**：Redis 预扣库存、Lua 脚本保证原子性

#### 5. Order RPC (端口: 8004)
- **职责**：订单创建、查询、状态管理
- **数据**：订单表
- **事务处理**：分布式事务、幂等性保证

#### 6. Blog RPC (端口: 8005)
- **职责**：笔记发布、点赞、评论
- **数据**：博客表、点赞表
- **缓存策略**：点赞数缓存、热门笔记排行榜

### 数据流转示例

**用户登录流程**：
```
1. 前端 → API Gateway (POST /user/login)
2. API Gateway → User RPC (gRPC Login)
3. User RPC → Redis (验证验证码)
4. User RPC → MySQL (查询/创建用户)
5. User RPC → Redis (创建 Session)
6. User RPC ← MySQL (用户信息)
7. API Gateway ← User RPC (返回 Token)
8. 前端 ← API Gateway (返回用户信息)
```

**秒杀优惠券流程**：
```
1. 前端 → API Gateway (POST /voucher/seckill/{id})
2. API Gateway → Order RPC (gRPC CreateSeckillOrder)
3. Order RPC → Redis (Lua 脚本检查并扣减库存)
4. Order RPC → Voucher RPC (gRPC CheckVoucherStock)
5. Order RPC → MySQL (创建订单记录)
6. Order RPC → Redis (缓存订单信息)
7. API Gateway ← Order RPC (返回订单 ID)
8. 前端 ← API Gateway (下单成功)
```

---

## 💾 数据库设计

### ER 图

```
┌─────────────┐         ┌─────────────┐         ┌─────────────┐
│   tb_user   │         │   tb_shop   │         │ tb_voucher  │
├─────────────┤         ├─────────────┤         ├─────────────┤
│ id (PK)     │         │ id (PK)     │         │ id (PK)     │
│ phone       │         │ name        │         │ shop_id(FK) │
│ password    │         │ type_id(FK) │         │ title       │
│ nick_name   │         │ images      │         │ pay_value   │
│ icon        │         │ area        │         │ actual_value│
│ create_time │         │ address     │         │ type        │
│ update_time │         │ x, y        │         │ status      │
└─────────────┘         │ avg_price   │         │ create_time │
                        │ sold        │         └──────┬──────┘
                        │ comments    │                │
                        │ score       │         ┌──────┴──────┐
                        │ create_time │         │tb_seckill_  │
                        └──────┬──────┘         │  voucher    │
                               │                ├─────────────┤
                        ┌──────┴──────┐         │voucher_id PK│
                        │tb_shop_type │         │ stock       │
                        ├─────────────┤         │ begin_time  │
                        │ id (PK)     │         │ end_time    │
                        │ name        │         │ create_time │
                        │ icon        │         └─────────────┘
                        │ sort_order  │
                        └─────────────┘

┌─────────────┐         ┌─────────────┐         ┌─────────────┐
│ tb_voucher_ │         │   tb_blog   │         │  tb_follow  │
│   order     │         ├─────────────┤         ├─────────────┤
├─────────────┤         │ id (PK)     │         │ id (PK)     │
│ id (PK)     │         │ shop_id(FK) │         │ user_id(FK) │
│ user_id(FK) │────┐    │ user_id(FK) │         │follow_uid   │
│voucher_id FK│    └───→│ title       │         │ create_time │
│ pay_type    │         │ images      │         └─────────────┘
│ status      │         │ content     │
│ create_time │         │ liked       │
│ pay_time    │         │ comments    │
│ use_time    │         │ create_time │
└─────────────┘         └─────────────┘
```

### 核心表设计

#### 1. tb_user - 用户表
```sql
CREATE TABLE `tb_user` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `phone` VARCHAR(11) NOT NULL COMMENT '手机号码',
  `password` VARCHAR(128) DEFAULT '' COMMENT '密码，加密存储',
  `nick_name` VARCHAR(32) DEFAULT '' COMMENT '昵称，默认是随机字符',
  `icon` VARCHAR(255) DEFAULT '' COMMENT '人物头像',
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uniq_phone` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
```

**设计考虑**：
- 使用手机号作为唯一标识，便于登录和找回密码
- 密码字段预留，支持密码登录扩展
- nick_name 和 icon 支持用户个性化
- 时间戳自动维护

#### 2. tb_shop - 商铺表
```sql
CREATE TABLE `tb_shop` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` VARCHAR(128) NOT NULL COMMENT '商铺名称',
  `type_id` BIGINT UNSIGNED NOT NULL COMMENT '商铺类型的id',
  `images` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '商铺图片，多个图片以,隔开',
  `area` VARCHAR(128) DEFAULT '' COMMENT '商圈，例如陆家嘴',
  `address` VARCHAR(255) NOT NULL COMMENT '地址',
  `x` DOUBLE NOT NULL COMMENT '经度',
  `y` DOUBLE NOT NULL COMMENT '纬度',
  `avg_price` BIGINT UNSIGNED DEFAULT '0' COMMENT '均价，取整数',
  `sold` INT UNSIGNED DEFAULT '0' COMMENT '销量',
  `comments` INT UNSIGNED DEFAULT '0' COMMENT '评论数量',
  `score` INT UNSIGNED DEFAULT '50' COMMENT '评分，1~50分，乘10保存，避免小数',
  `open_hours` VARCHAR(32) DEFAULT '' COMMENT '营业时间，例如 10:00-22:00',
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_type_id` (`type_id`) USING BTREE,
  KEY `idx_area` (`area`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商铺表';
```

**设计考虑**：
- x, y 字段存储经纬度，支持地理位置搜索（可扩展 GeoHash）
- score 乘以 10 存储，避免浮点数精度问题
- 冗余 sold、comments 字段，避免频繁关联查询
- 索引优化：type_id 和 area 是常用筛选条件

#### 3. tb_voucher - 优惠券表
```sql
CREATE TABLE `tb_voucher` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `shop_id` BIGINT UNSIGNED NOT NULL COMMENT '商铺id',
  `title` VARCHAR(255) NOT NULL COMMENT '代金券标题',
  `sub_title` VARCHAR(255) DEFAULT '' COMMENT '副标题',
  `rules` VARCHAR(1024) DEFAULT '' COMMENT '使用规则',
  `pay_value` BIGINT UNSIGNED NOT NULL COMMENT '支付金额，单位是分',
  `actual_value` BIGINT UNSIGNED NOT NULL COMMENT '抵扣金额，单位是分',
  `type` TINYINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '0,普通券；1,秒杀券',
  `status` TINYINT UNSIGNED NOT NULL DEFAULT '1' COMMENT '1,上架; 2,下架; 3,过期',
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_shop_id` (`shop_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='优惠券表';
```

**设计考虑**：
- 金额以分为单位存储，避免浮点数问题
- type 字段区分普通券和秒杀券
- status 字段支持券的生命周期管理

#### 4. tb_seckill_voucher - 秒杀券表
```sql
CREATE TABLE `tb_seckill_voucher` (
  `voucher_id` BIGINT UNSIGNED NOT NULL COMMENT '关联的优惠券的id',
  `stock` INT NOT NULL COMMENT '库存',
  `begin_time` TIMESTAMP NOT NULL COMMENT '生效时间',
  `end_time` TIMESTAMP NOT NULL COMMENT '失效时间',
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`voucher_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='秒杀优惠券表，与优惠券是一对一关系';
```

**设计考虑**：
- 一对一关系，使用 voucher_id 作为主键
- stock 字段实时更新，但真正的库存控制在 Redis 中
- begin_time 和 end_time 控制秒杀时间窗口

#### 5. tb_voucher_order - 优惠券订单表
```sql
CREATE TABLE `tb_voucher_order` (
  `id` BIGINT NOT NULL COMMENT '主键',
  `user_id` BIGINT UNSIGNED NOT NULL COMMENT '下单的用户id',
  `voucher_id` BIGINT UNSIGNED NOT NULL COMMENT '购买的代金券id',
  `pay_type` TINYINT UNSIGNED NOT NULL DEFAULT '1' COMMENT '支付方式 1：余额支付；2：支付宝；3：微信',
  `status` TINYINT UNSIGNED NOT NULL DEFAULT '1' COMMENT '订单状态，1：未支付；2：已支付；3：已核销；4：已取消；5：退款中；6：已退款',
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '下单时间',
  `pay_time` TIMESTAMP NULL DEFAULT NULL COMMENT '支付时间',
  `use_time` TIMESTAMP NULL DEFAULT NULL COMMENT '核销时间',
  `refund_time` TIMESTAMP NULL DEFAULT NULL COMMENT '退款时间',
  `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_voucher_id` (`voucher_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='优惠券订单表';
```

**设计考虑**：
- 使用分布式 ID（雪花算法）作为主键
- 订单状态支持完整的生命周期
- 时间字段支持订单各个阶段的追踪
- 一人一单限制在业务层实现（Redis + Lua）

### 表关系说明

1. **用户 ↔ 订单**：一对多关系
   - 一个用户可以下多个订单
   - 订单通过 user_id 关联用户

2. **商铺 ↔ 优惠券**：一对多关系
   - 一个商铺可以发布多张优惠券
   - 优惠券通过 shop_id 关联商铺

3. **优惠券 ↔ 秒杀券**：一对一关系
   - 秒杀券是优惠券的扩展信息
   - 使用 voucher_id 作为主键实现一对一

4. **优惠券 ↔ 订单**：一对多关系
   - 一张优惠券可以被多个用户购买（库存允许）
   - 订单通过 voucher_id 关联优惠券

5. **用户 ↔ 博客**：一对多关系
   - 一个用户可以发布多篇博客
   - 博客通过 user_id 关联用户

---

## 🔧 环境准备（新手必读）

> 如果你已经有完整的开发环境，可以跳过这一章节，直接看 [快速开始](#-快速开始)

本章节将手把手教你从零搭建开发环境。即使你是完全的新手，只要按照步骤操作，也能顺利完成。

### 1. 安装 Go 环境

#### 1.1 下载 Go

访问 Go 官网下载页面：https://go.dev/dl/

**推荐版本**：Go 1.23 或更高版本

**选择对应的安装包**：
- Linux (x86_64): `go1.23.5.linux-amd64.tar.gz`
- macOS (Intel): `go1.23.5.darwin-amd64.pkg`
- macOS (Apple Silicon): `go1.23.5.darwin-arm64.pkg`
- Windows: `go1.23.5.windows-amd64.msi`

#### 1.2 安装 Go

**Linux 安装步骤**：

```bash
# 1. 下载（可以换成最新版本号）
wget https://go.dev/dl/go1.23.5.linux-amd64.tar.gz

# 2. 解压到 /usr/local 目录
# -C 指定解压目录，-z 表示 gzip 压缩，-x 表示解压，-v 显示详细信息，-f 指定文件
sudo tar -C /usr/local/ -zxvf go1.23.5.linux-amd64.tar.gz

# 3. 验证解压
ls /usr/local/go/bin/go  # 应该能看到 go 可执行文件
```

**macOS 安装步骤**：

```bash
# 方式一：使用 pkg 安装包（推荐，会自动配置环境变量）
# 下载 .pkg 文件后双击安装即可

# 方式二：使用 tar.gz 手动安装（和 Linux 相同）
wget https://go.dev/dl/go1.23.5.darwin-amd64.tar.gz
sudo tar -C /usr/local/ -zxvf go1.23.5.darwin-amd64.tar.gz
```

**Windows 安装步骤**：

```bash
# 下载 .msi 文件后双击安装即可
# 安装程序会自动配置环境变量
```

#### 1.3 配置环境变量

**Linux / macOS (bash)**：

```bash
# 编辑 ~/.bashrc 文件
vim ~/.bashrc
# 或使用 nano（更简单）
nano ~/.bashrc

# 在文件末尾添加以下内容：
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

# 保存后，让配置生效
source ~/.bashrc
```

**macOS (zsh - 默认 shell)**：

```bash
# 编辑 ~/.zshrc 文件
vim ~/.zshrc

# 添加相同的环境变量（同上）

# 让配置生效
source ~/.zshrc
```

**Windows**：

```powershell
# 如果使用 .msi 安装，环境变量已自动配置
# 如果需要手动配置：
# 1. 右键"此电脑" → "属性" → "高级系统设置" → "环境变量"
# 2. 在"系统变量"中新建：
#    GOROOT = C:\Go
#    GOPATH = C:\Users\你的用户名\go
# 3. 编辑 Path 变量，添加：
#    %GOROOT%\bin
#    %GOPATH%\bin
```

#### 1.4 配置 GOPROXY（国内必须！）

由于网络原因，国内直接访问 Go 官方模块代理会很慢或失败，**必须配置国内镜像**：

```bash
# 配置七牛云提供的 goproxy（推荐）
go env -w GOPROXY=https://goproxy.cn,direct

# 或使用阿里云镜像
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# 关闭模块验证（可选，某些私有仓库需要）
go env -w GOSUMDB=off

# 开启 Go Modules（Go 1.16+ 默认开启）
go env -w GO111MODULE=on
```

**解释**：
- `GOPROXY`：指定 Go 模块代理服务器
- `,direct`：如果代理不可用，直接访问源站
- `GOSUMDB`：模块校验数据库，设为 off 可跳过校验

#### 1.5 验证安装

```bash
# 查看 Go 版本
go version
# 应该输出：go version go1.23.5 linux/amd64

# 查看 Go 环境配置
go env
# 检查 GOROOT、GOPATH、GOPROXY 是否正确

# 测试 Go Modules 下载
go mod download
```

### 2. 安装 Docker 环境

Docker 是本项目的核心依赖，必须正确安装和配置。

#### 2.1 Linux 安装 Docker

**方式一：使用国内优化脚本（强烈推荐！）**

[@tech-shrimp](https://github.com/tech-shrimp) 提供了一个针对国内环境优化的 Docker 安装脚本：

```bash
# 1. 克隆安装脚本仓库
git clone https://github.com/tech-shrimp/docker_installer.git
cd docker_installer

# 2. 给脚本执行权限
chmod +x docker_installer.sh

# 3. 运行安装脚本（会自动检测系统并安装）
sudo ./docker_installer.sh

# 4. 启动 Docker
sudo systemctl start docker
sudo systemctl enable docker

# 5. 验证安装
docker --version
sudo docker run hello-world
```

**方式二：使用官方脚本（可能较慢）**

```bash
# Ubuntu / Debian
curl -fsSL https://get.docker.com | sudo bash -s docker

# CentOS / RHEL（使用阿里云镜像）
curl -fsSL https://get.docker.com | sudo bash -s docker --mirror Aliyun

# 启动 Docker
sudo systemctl start docker
sudo systemctl enable docker
```

**方式三：手动安装（各发行版）**

<details>
<summary>Ubuntu / Debian 手动安装步骤</summary>

```bash
# 1. 更新包索引
sudo apt-get update

# 2. 安装依赖
sudo apt-get install -y \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

# 3. 添加 Docker GPG 密钥
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | \
    sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

# 4. 添加 Docker 仓库（使用阿里云镜像）
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] \
  https://mirrors.aliyun.com/docker-ce/linux/ubuntu \
  $(lsb_release -cs) stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# 5. 安装 Docker
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io \
    docker-buildx-plugin docker-compose-plugin

# 6. 启动 Docker
sudo systemctl start docker
sudo systemctl enable docker
```

</details>

<details>
<summary>CentOS / RHEL / veLinux 手动安装步骤</summary>

```bash
# 1. 安装依赖
sudo yum install -y yum-utils device-mapper-persistent-data lvm2

# 2. 添加 Docker 仓库（使用阿里云镜像）
sudo yum-config-manager --add-repo \
    https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

# 3. 安装 Docker
sudo yum install -y docker-ce docker-ce-cli containerd.io \
    docker-buildx-plugin docker-compose-plugin

# 4. 启动 Docker
sudo systemctl start docker
sudo systemctl enable docker
```

</details>

#### 2.2 配置 Docker 镜像加速（国内必须！）

Docker Hub 在国内访问很慢，**必须配置镜像加速**：

```bash
# 创建 Docker 配置目录
sudo mkdir -p /etc/docker

# 配置镜像加速器（使用多个镜像源，提高成功率）
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": [
    "https://dockerproxy.cn",
    "https://docker.1panel.live",
    "https://hub.rat.dev",
    "https://docker.m.daocloud.io",
    "https://docker.mirrors.ustc.edu.cn"
  ],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "10m",
    "max-file": "3"
  }
}
EOF

# 重启 Docker 服务
sudo systemctl daemon-reload
sudo systemctl restart docker

# 验证配置
sudo docker info | grep -A 10 "Registry Mirrors"
```

**镜像源说明**：
- `dockerproxy.cn`：Docker 代理加速
- `docker.1panel.live`：1Panel 提供的镜像
- `hub.rat.dev`：社区镜像
- `docker.m.daocloud.io`：DaoCloud 镜像
- `docker.mirrors.ustc.edu.cn`：中科大镜像

#### 2.3 配置 Docker 用户权限（可选）

避免每次使用 Docker 都要 `sudo`：

```bash
# 将当前用户添加到 docker 组
sudo usermod -aG docker $USER

# 重新登录或执行以下命令使权限生效
newgrp docker

# 验证（不需要 sudo）
docker ps
```

#### 2.4 macOS 安装 Docker

```bash
# 1. 下载 Docker Desktop for Mac
# 访问：https://www.docker.com/products/docker-desktop

# 2. 安装 .dmg 文件

# 3. 启动 Docker Desktop

# 4. 配置镜像加速
# 打开 Docker Desktop → Settings → Docker Engine
# 在 JSON 配置中添加：
{
  "registry-mirrors": [
    "https://dockerproxy.cn",
    "https://docker.mirrors.ustc.edu.cn"
  ]
}
```

#### 2.5 Windows 安装 Docker

```bash
# 1. 启用 WSL2（Windows 10/11）
# 打开 PowerShell（管理员模式）：
wsl --install

# 2. 下载 Docker Desktop for Windows
# 访问：https://www.docker.com/products/docker-desktop

# 3. 安装并重启

# 4. 配置镜像加速（同 macOS）
```

#### 2.6 验证 Docker 安装

```bash
# 查看 Docker 版本
docker --version
# 应输出：Docker version 24.0.0 或更高

# 查看 Docker Compose 版本
docker compose version
# 应输出：Docker Compose version v2.x.x

# 运行测试容器
docker run hello-world
# 应输出：Hello from Docker!

# 查看 Docker 信息
docker info
```

### 3. 安装其他工具

#### 3.1 安装 Git

**Linux**：
```bash
# Ubuntu / Debian
sudo apt-get install -y git

# CentOS / RHEL
sudo yum install -y git
```

**macOS**：
```bash
# 使用 Homebrew
brew install git

# 或使用 Xcode Command Line Tools
xcode-select --install
```

**Windows**：
- 下载：https://git-scm.com/download/win
- 双击安装即可

#### 3.2 IDE 推荐

**GoLand（商业，推荐）**：
- 下载：https://www.jetbrains.com/go/
- 功能最强大，调试方便
- 学生可免费使用

**VS Code（免费，推荐）**：
- 下载：https://code.visualstudio.com/
- 安装扩展：
  - Go（官方）
  - Docker
  - REST Client
  - GitLens

### 4. 环境验证清单

安装完成后，运行以下命令验证：

```bash
# ✅ Go 环境
go version                    # 应显示 go1.23+
go env GOPROXY                # 应显示 https://goproxy.cn,direct

# ✅ Docker 环境
docker --version              # 应显示 Docker version 20.10+
docker compose version        # 应显示 Docker Compose version v2.0+
docker run hello-world        # 应成功运行

# ✅ Git
git --version                 # 应显示 git version 2.x.x

# ✅ 网络测试
ping -c 4 github.com          # 测试 GitHub 访问
curl -I https://goproxy.cn    # 测试 Go 代理
```

### 5. 常见问题

<details>
<summary>Q: Go 下载太慢怎么办？</summary>

**A**: 使用国内镜像下载：
```bash
# 使用阿里云镜像
wget https://mirrors.aliyun.com/golang/go1.23.5.linux-amd64.tar.gz

# 或使用七牛云镜像
wget https://golang.google.cn/dl/go1.23.5.linux-amd64.tar.gz
```
</details>

<details>
<summary>Q: Docker 镜像拉取失败 403/429 错误？</summary>

**A**: 更换镜像源或使用多个镜像源：
```bash
# 配置多个镜像源（按顺序尝试）
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": [
    "https://dockerproxy.cn",
    "https://docker.1panel.live",
    "https://hub.rat.dev"
  ]
}
EOF
sudo systemctl restart docker
```
</details>

<details>
<summary>Q: go get 下载依赖失败？</summary>

**A**: 确认 GOPROXY 配置正确：
```bash
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=off  # 如果还是失败，可以关闭校验
```
</details>

<details>
<summary>Q: Docker 权限问题 "permission denied"？</summary>

**A**:
```bash
# 方式一：将用户加入 docker 组
sudo usermod -aG docker $USER
newgrp docker

# 方式二：使用 sudo
sudo docker ps
```
</details>

---

## 🚀 快速开始

> 请确保已完成 [环境准备](#-环境准备新手必读) 章节的所有步骤

### 前置要求

- ✅ **操作系统**：Linux / macOS / Windows (WSL2)
- ✅ **Docker**：20.10+
- ✅ **Docker Compose**：2.0+
- ✅ **Go**：1.23+ （本地开发）
- ✅ **Git**：任意版本

### 一键启动（推荐）

```bash
# 1. 克隆项目
git clone https://github.com/me-cs/dianping-gozero.git
cd dianping-gozero/backend

# 2. 启动所有服务（会自动下载镜像、构建、启动）
./start.sh

# 3. 等待服务启动完成（约 2-3 分钟）
# 看到 "All services started successfully!" 表示成功

# 4. 验证服务
curl http://localhost:8081/health
```

就这么简单！所有服务都已经运行起来了。

### 访问服务

| 服务 | 地址 | 说明 |
|------|------|------|
| **Nginx** | http://localhost:80 | 反向代理（自动启动） |
| **API Gateway** | http://localhost:8081 | RESTful API 入口 |
| **Grafana** | http://localhost:3000 | 监控面板 (admin/admin) |
| **Prometheus** | http://localhost:9090 | 指标数据 |
| **Jaeger UI** | http://localhost:16686 | 链路追踪可视化 |
| **MySQL** | localhost:3306 | 数据库 (root/root) |
| **Redis** | localhost:6379 | 缓存 |
| **etcd** | localhost:2379 | 服务发现 |

> **提示**：Nginx 反向代理会自动启动，你可以直接通过 http://localhost 访问 API（默认 80 端口），请求会自动转发到 API Gateway (8081)

### 测试 API

使用 Postman 或 curl 测试：

```bash
# 1. 发送验证码
curl -X POST http://localhost:8081/user/code \
  -H "Content-Type: application/json" \
  -d '{"phone":"13800138000"}'

# 2. 登录（验证码在 Redis 中，可以用任意 6 位数字测试）
curl -X POST http://localhost:8081/user/login \
  -H "Content-Type: application/json" \
  -d '{"phone":"13800138000","code":"123456"}'

# 3. 查询商铺列表（需要 token）
curl http://localhost:8081/shop/list/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## 📦 部署指南

### 开发环境部署

使用 `start.sh` 一键启动：

```bash
./start.sh                      # 默认数据目录: /var/lib/dianping
./start.sh /custom/path         # 自定义数据目录
./start.sh --reset              # 清空数据重新启动
```

### 生产环境部署

#### 1. 准备服务器

```bash
# 最低配置
- CPU: 2 核
- 内存: 4GB
- 磁盘: 20GB
- 操作系统: Ubuntu 20.04 / CentOS 7+
```

#### 2. 安装 Docker

如果服务器上还没有安装 Docker，请参考 [环境准备](#-环境准备新手必读) 章节的详细安装步骤。

**快速安装（生产环境推荐）**：

```bash
# 使用国内优化脚本
git clone https://github.com/tech-shrimp/docker_installer.git
cd docker_installer
sudo ./docker_installer.sh

# 启动 Docker
sudo systemctl start docker
sudo systemctl enable docker

# 配置镜像加速（必须！）
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": [
    "https://dockerproxy.cn",
    "https://docker.1panel.live",
    "https://hub.rat.dev"
  ]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker
```

#### 3. 克隆并部署

```bash
git clone https://github.com/me-cs/dianping-gozero.git
cd dianping-gozero/backend

# 设置数据目录
export DATA_DIR=/data/dianping

# 启动服务
./start.sh $DATA_DIR
```

#### 4. 配置 Nginx（可选）

```nginx
upstream api_backend {
    server 127.0.0.1:8081;
}

server {
    listen 80;
    server_name api.yourdomain.com;

    location / {
        proxy_pass http://api_backend;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

#### 5. 监控和日志

```bash
# 查看所有容器状态
docker-compose ps

# 查看服务日志
docker-compose logs -f user-rpc
docker-compose logs -f api-gateway

# 查看资源使用
docker stats
```

### 高可用部署（进阶）

对于生产环境，建议：

1. **多实例部署**：每个 RPC 服务部署多个实例
2. **负载均衡**：使用 Nginx 或 K8s Ingress
3. **数据库主从**：MySQL 主从复制 + 读写分离
4. **Redis 集群**：Redis Sentinel 或 Cluster 模式
5. **持久化存储**：使用外部存储卷（NFS、Ceph 等）

---

## 💻 开发指南

### 环境搭建

> 如果你还没有安装 Go、Docker 等环境，请先阅读 [环境准备（新手必读）](#-环境准备新手必读) 章节。

#### 1. 验证环境

确保以下环境已正确安装：

```bash
# 检查 Go 版本
go version        # 应为 go1.23+

# 检查 Docker
docker --version  # 应为 Docker 20.10+

# 检查 Git
git --version

# 检查 GOPROXY 配置
go env GOPROXY    # 应为 https://goproxy.cn,direct
```

#### 2. 安装 IDE

推荐使用 **GoLand**（商业）或 **VS Code**（免费）

**VS Code 扩展**：
- Go (official)
- Go Test Explorer
- REST Client
- Docker
- Protocol Buffers

#### 3. 克隆项目

```bash
git clone https://github.com/me-cs/dianping-gozero.git
cd dianping-gozero/backend

# 下载依赖
go mod download
```

### 项目结构

```
backend/
├── api/                      # API Gateway
│   ├── internal/
│   │   ├── config/          # 配置
│   │   ├── handler/         # HTTP 处理器
│   │   ├── logic/           # 业务逻辑
│   │   ├── svc/             # 服务上下文
│   │   └── types/           # 类型定义
│   ├── etc/                 # 配置文件
│   └── dianping.go          # 主入口
│
├── rpc/                      # RPC 服务
│   ├── user/
│   │   ├── internal/
│   │   │   ├── config/
│   │   │   ├── logic/       # 业务逻辑
│   │   │   ├── server/      # gRPC 服务器
│   │   │   └── svc/
│   │   ├── pb/              # Protobuf 生成代码
│   │   ├── etc/             # 配置文件
│   │   └── user.proto       # Protobuf 定义
│   ├── shop/
│   ├── voucher/
│   ├── order/
│   └── blog/
│
├── common/                   # 共享代码
│   ├── errorx/              # 错误处理
│   ├── utils/               # 工具函数
│   └── types/               # 共享类型
│
├── deploy/                   # 部署配置
│   ├── mysql/
│   │   ├── init/            # 初始化 SQL
│   │   └── conf/            # MySQL 配置
│   ├── prometheus/
│   └── grafana/
│
├── docker-compose.yml        # Docker Compose 配置
├── start.sh                  # 启动脚本
├── stop.sh                   # 停止脚本
├── build-binaries.sh         # 编译脚本
└── build-docker.sh           # 构建镜像脚本
```

### 添加新功能

#### 示例：添加新的 RPC 服务

```bash
# 1. 定义 Protobuf
cd rpc/newservice
vim newservice.proto

# 2. 生成代码
goctl rpc protoc newservice.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.

# 3. 实现业务逻辑
vim internal/logic/xxxLogic.go

# 4. 更新配置
vim etc/newservice.yaml

# 5. 更新 docker-compose.yml
# 添加新服务配置

# 6. 编译和启动
./build-binaries.sh
docker-compose up -d newservice-rpc
```

#### 示例：添加新的 API 接口

```bash
# 1. 定义 API
cd api
vim api/newapi.api

# 2. 生成代码
goctl api go -api api/newapi.api -dir ./

# 3. 实现 Logic
vim internal/logic/newapi/xxxLogic.go

# 4. 重新编译
cd ..
./build-binaries.sh
docker-compose restart api-gateway
```

### 代码规范

- **命名**：遵循 Go 官方命名规范（驼峰命名）
- **注释**：公开 API 必须有注释
- **错误处理**：使用 `errorx` 包统一错误处理
- **日志**：使用 `logx` 包，避免使用 `fmt.Println`
- **格式化**：使用 `gofmt` 或 `goimports`

### 测试

#### 单元测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./rpc/user/internal/logic

# 查看覆盖率
go test -cover ./...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

#### 集成测试

```bash
# 确保服务运行
./start.sh

# 运行集成测试
go test -tags=integration ./test/integration

# API 测试
cd test/api
go test -v
```

---

## 🐛 调试与故障排查

### 查看日志

```bash
# 实时查看所有服务日志
docker-compose logs -f

# 查看特定服务日志
docker-compose logs -f user-rpc
docker-compose logs -f api-gateway

# 查看最近 100 行
docker-compose logs --tail=100 user-rpc

# 导出日志
docker-compose logs > logs.txt
```

### 进入容器调试

```bash
# 进入服务容器
docker exec -it dianping-user-rpc /bin/sh

# 进入 MySQL
docker exec -it dianping-mysql mysql -u root -proot

# 进入 Redis
docker exec -it dianping-redis redis-cli
```

### 常见问题排查

#### 1. 服务启动失败

```bash
# 检查容器状态
docker-compose ps

# 查看失败原因
docker-compose logs <service-name>

# 常见原因：
# - 端口被占用：修改 docker-compose.yml 端口映射
# - 依赖服务未就绪：等待 MySQL/Redis 启动完成
# - 配置错误：检查 etc/*.yaml 配置文件
```

#### 2. RPC 调用失败

```bash
# 检查 etcd 服务注册
docker exec dianping-etcd etcdctl get --prefix ""

# 检查服务是否注册成功
# 应该看到类似 user.rpc/xxxx 的键

# 查看 Jaeger 追踪
# 访问 http://localhost:16686 查看调用链
```

#### 3. 数据库连接问题

```bash
# 进入 MySQL 容器
docker exec -it dianping-mysql mysql -u root -proot

# 检查数据库
SHOW DATABASES;
USE hmdp;
SHOW TABLES;

# 检查用户权限
SELECT User, Host FROM mysql.user;
```

#### 4. Redis 连接问题

```bash
# 进入 Redis 容器
docker exec -it dianping-redis redis-cli

# 测试连接
PING  # 应返回 PONG

# 查看所有键
KEYS *

# 查看特定键
GET login:code:13800138000
```

#### 5. 网络问题

```bash
# 检查 Docker 网络
docker network ls
docker network inspect backend_dianping-network

# 测试服务间连通性
docker exec dianping-user-rpc ping mysql
docker exec dianping-user-rpc ping redis
```

### 使用 Jaeger 追踪

1. 访问 http://localhost:16686
2. 选择服务（如 user.rpc）
3. 点击 "Find Traces" 查看调用链
4. 分析慢查询和错误

### 性能监控

1. 访问 Grafana：http://localhost:3000 (admin/admin)
2. 导入预定义的 Dashboard
3. 查看 CPU、内存、请求量、响应时间等指标

---

## 🎓 学习路径

### 从这个项目你可以学到什么？

这不是一个简单的 CRUD 项目，而是一个**完整的企业级微服务系统**，涵盖了从开发到部署的全流程。

#### 🔥 核心技术栈

1. **Go 语言**
   - Go 基础语法和特性
   - Goroutine 和 Channel 并发编程
   - 接口和反射
   - 测试和性能优化

2. **go-zero 微服务框架**
   - API 网关设计
   - gRPC 服务开发
   - 中间件使用
   - 服务注册与发现
   - 配置管理

3. **gRPC & Protobuf**
   - Protocol Buffers 语法
   - gRPC 服务定义
   - RPC 调用和错误处理
   - 拦截器（Interceptor）

#### 💾 数据存储

4. **MySQL**
   - 表设计和 ER 图
   - 索引优化
   - 事务处理
   - 连接池配置

5. **Redis**
   - 缓存策略（Cache-Aside、Read-Through）
   - 分布式锁
   - Session 管理
   - Lua 脚本
   - 排行榜实现
   - 发布订阅

6. **etcd**
   - 服务注册与发现
   - 配置中心
   - 分布式协调

#### 🏗️ 系统设计

7. **微服务架构**
   - 服务拆分原则
   - 服务间通信（同步/异步）
   - API 网关模式
   - 服务治理

8. **高并发设计**
   - 秒杀系统设计
   - 库存扣减方案
   - 一人一单实现
   - 幂等性保证

9. **缓存设计**
   - 缓存穿透、击穿、雪崩
   - 缓存预热
   - 缓存更新策略
   - 多级缓存

10. **分布式事务**
    - 最终一致性
    - 补偿机制
    - 幂等设计

#### 🐳 DevOps

11. **Docker**
    - Dockerfile 编写
    - Docker Compose 编排
    - 容器化最佳实践
    - 多阶段构建

12. **CI/CD**
    - Git 工作流
    - 自动化构建
    - 自动化测试
    - 容器化部署
    - Dependabot 依赖自动更新

#### 📊 可观测性

13. **监控**
    - Prometheus 指标收集
    - Grafana 可视化
    - 告警配置

14. **链路追踪**
    - Jaeger 分布式追踪
    - Span 和 Trace
    - 性能分析

15. **日志**
    - 结构化日志
    - 日志收集和分析
    - 日志级别管理

#### 🧪 测试

16. **单元测试**
    - Go testing 框架
    - Mock 和 Stub
    - 测试覆盖率

17. **集成测试**
    - API 测试
    - 数据库测试
    - Redis 测试

#### 🤖 AI 辅助开发

18. **Claude AI**
    - AI 辅助编程
    - 代码生成和重构
    - 问题诊断
    - 文档编写

### 学习建议

#### 初级阶段（1-2 周）

1. **环境搭建**：完成 Docker、Go 等工具安装
2. **运行项目**：使用 `./start.sh` 启动项目
3. **熟悉接口**：使用 Postman 测试所有 API
4. **阅读代码**：从 API Gateway 开始，理解请求流程
5. **查看日志**：学习使用 `docker-compose logs` 查看日志
6. **数据库操作**：进入 MySQL 查看表结构和数据

#### 中级阶段（2-4 周）

1. **添加功能**：尝试添加新的 API 接口
2. **修改逻辑**：修改现有业务逻辑，理解代码结构
3. **调试代码**：使用 Jaeger 追踪请求链路
4. **优化性能**：分析慢查询，添加索引
5. **编写测试**：为关键功能编写单元测试
6. **学习 Redis**：理解缓存策略和分布式锁

#### 高级阶段（4-8 周）

1. **架构重构**：尝试改进架构设计
2. **性能优化**：进行压力测试和性能优化
3. **高可用部署**：实现服务多实例部署
4. **监控告警**：配置 Prometheus 告警规则
5. **秒杀系统**：深入理解秒杀流程和优化方案
6. **前端开发**：使用 Vue/React 开发前端界面

### 推荐资源

- **Go 语言**
  - [Go 语言官方文档](https://go.dev/doc/)
  - [Go by Example](https://gobyexample.com/)
  - [Effective Go](https://go.dev/doc/effective_go)

- **go-zero**
  - [go-zero 官方文档](https://go-zero.dev/)
  - [go-zero 实战教程](https://go-zero.dev/docs/tutorials)

- **微服务**
  - [微服务设计模式](https://microservices.io/)
  - [《微服务架构设计模式》](https://www.manning.com/books/microservices-patterns)

- **Redis**
  - [Redis 官方文档](https://redis.io/docs/)
  - [《Redis 设计与实现》](http://redisbook.com/)

- **Docker**
  - [Docker 官方文档](https://docs.docker.com/)
  - [Docker 从入门到实践](https://yeasy.gitbook.io/docker_practice/)

---

## 🤝 贡献指南

我们非常欢迎各种形式的贡献！无论是报告 Bug、提出新功能、改进文档还是提交代码。

### 如何贡献

#### 1. Fork 项目

点击右上角的 "Fork" 按钮，将项目 Fork 到你的账号下。

#### 2. 克隆到本地

```bash
git clone https://github.com/YOUR_USERNAME/dianping-gozero.git
cd dianping-gozero
```

#### 3. 创建分支

```bash
git checkout -b feature/your-feature-name
# 或
git checkout -b fix/your-bug-fix
```

#### 4. 开发

- 添加你的功能或修复 Bug
- 编写测试
- 确保所有测试通过
- 更新文档

#### 5. 提交代码

```bash
git add .
git commit -m "feat: add awesome feature"

# 提交信息规范（参考 Conventional Commits）
# feat: 新功能
# fix: 修复 Bug
# docs: 文档更新
# style: 代码格式
# refactor: 重构
# test: 测试
# chore: 构建/工具
```

#### 6. 推送到 GitHub

```bash
git push origin feature/your-feature-name
```

#### 7. 创建 Pull Request

- 访问你的 Fork 仓库
- 点击 "New Pull Request"
- 填写 PR 描述，说明你的改动
- 等待 Review

### 贡献领域

#### 🎨 前端开发（急需！）

**我们非常欢迎有前端经验的同学加入，开发本项目的前端部分！**

推荐技术栈：
- **Vue 3** + Vite + TypeScript
- **React** + Next.js + TypeScript
- **组件库**：Element Plus / Ant Design / Chakra UI
- **状态管理**：Pinia / Redux / Zustand
- **地图**：高德地图 / 百度地图 API

前端功能需求：
- [ ] 用户登录/注册页面
- [ ] 首页（商铺列表、分类）
- [ ] 商铺详情页
- [ ] 优惠券列表
- [ ] 秒杀页面
- [ ] 个人中心
- [ ] 探店笔记
- [ ] 后台管理系统

**如果你想开发前端，请联系我们或直接提 Issue！**

#### 🔧 后端改进

- [ ] 添加单元测试
- [ ] 性能优化
- [ ] 添加新功能
- [ ] 代码重构
- [ ] Bug 修复

#### 📝 文档完善

- [ ] 翻译英文文档
- [ ] 添加更多示例
- [ ] 改进 API 文档
- [ ] 视频教程

#### 🎨 其他

- [ ] 设计 Logo
- [ ] 架构图优化
- [ ] Grafana Dashboard
- [ ] K8s 部署方案

### 行为准则

- 尊重他人，友好交流
- 遵循项目代码规范
- 测试你的代码
- 清晰描述你的改动

---

## ⚖️ 法律声明

### 项目性质

本项目是一个**纯学习项目**，旨在帮助开发者学习和实践微服务架构、Go 语言开发、系统设计等技术。

### 重要提示

⚠️ **请注意以下法律风险**：

1. **商业使用风险**
   - 本项目仅供学习和研究使用
   - 未经授权不得用于商业目的
   - 项目名称、界面设计等可能涉及商标权

2. **知识产权**
   - 本项目参考了黑马点评的业务模型
   - 所有商标、服务标志归其各自所有者
   - 本项目不隶属于黑马点评或美团

3. **数据安全**
   - 请勿使用真实用户数据
   - 请勿在公网暴露测试环境
   - 注意保护用户隐私

4. **责任声明**
   - 项目作者不对使用本项目造成的任何损失负责
   - 使用者需自行承担所有法律风险
   - 建议仅在学习环境中使用

### 许可证

本项目采用 MIT License，详见 [LICENSE](LICENSE) 文件。

**MIT License 摘要**：
- ✅ 可以自由使用、修改、分发
- ✅ 可以用于商业项目（需自行承担风险）
- ❗ 需保留版权声明
- ❗ 软件按"原样"提供，不含任何担保

### 免责声明

使用本项目即表示你已阅读并同意：
- 本项目仅供学习使用
- 作者不对任何后果负责
- 你将遵守当地法律法规

---

## 💝 致谢

### 🤖 特别鸣谢：Claude AI

**本项目的 90% 代码由 Claude AI 生成和优化！**

这是一次关于 AI 辅助编程的实验：
- **代码生成**：所有微服务代码、API 接口、数据库设计
- **架构设计**：系统架构、技术选型、最佳实践
- **问题诊断**：Bug 修复、性能优化、故障排查
- **文档编写**：README、API 文档、注释

**AI 辅助开发的优势**：
- ⚡ **开发效率**：10 倍提升（原本需要 2 周的工作，2 天完成）
- 🎯 **代码质量**：遵循最佳实践、一致性高
- 📚 **学习成本**：边开发边学习，实时解答疑问
- 🐛 **快速调试**：AI 能快速定位和修复问题

**如果你想学习如何使用 Claude 进行编程，这个项目是最好的示例！**

### 技术支持

感谢以下开源项目：

- [go-zero](https://github.com/zeromicro/go-zero) - 优秀的微服务框架
- [gRPC](https://grpc.io/) - 高性能 RPC 框架
- [Docker](https://www.docker.com/) - 容器化平台
- [Redis](https://redis.io/) - 高性能缓存
- [MySQL](https://www.mysql.com/) - 关系型数据库
- [Jaeger](https://www.jaegertracing.io/) - 分布式追踪
- [Prometheus](https://prometheus.io/) - 监控系统
- [Grafana](https://grafana.com/) - 可视化平台

### 学习资源

- 黑马程序员
- Go 语言中文网
- go-zero 官方文档

---

## ☕ 支持项目

如果这个项目对你有帮助，请考虑支持我们：

### ⭐ Star 项目

在 GitHub 上给我们一个 Star，让更多人看到这个项目！

### 💰 赞助

**Buy me a coffee! ☕**

你的支持将帮助我们：
- 🚀 持续维护和更新项目
- 📚 创作更多教程和文档
- 🎥 录制视频课程
- 🎨 开发前端界面

**赞助方式**：

<details>
<summary>点击查看赞助二维码</summary>

#### 微信赞赏

<img src="./docs/images/wechat-pay.png" width="200" alt="微信赞赏码">

#### 支付宝

<img src="./docs/images/alipay.png" width="200" alt="支付宝收款码">

#### Buy Me a Coffee

[![Buy Me A Coffee](https://img.shields.io/badge/Buy%20Me%20A%20Coffee-Support-FFDD00?style=for-the-badge&logo=buy-me-a-coffee)](https://www.buymeacoffee.com/yourname)

</details>

### 🎁 赞助者名单

感谢以下赞助者（持续更新）：

| 赞助者 | 金额 | 留言 |
|--------|------|------|
| **\*\*明** | ¥66 | 很棒的项目！ |
| **\*\*华** | ¥100 | 学到很多 |
| **\*\*\*\*** | $10 | Great work! |

---

## 📞 联系我们

### 交流讨论

- **GitHub Issues**: [提问题/建议](https://github.com/me-cs/dianping-gozero/issues)
- **GitHub Discussions**: [讨论区](https://github.com/me-cs/dianping-gozero/discussions)

### 社交媒体

- **技术博客**: [博客地址]
- **B站**: [B站频道]
- **知乎**: [知乎主页]

### 问题反馈

遇到问题？请按以下步骤：

1. 查看 [常见问题](#常见问题)
2. 搜索 [已有 Issues](https://github.com/me-cs/dianping-gozero/issues)
3. 如果都没有，[创建新 Issue](https://github.com/me-cs/dianping-gozero/issues/new)

**提问时请包含**：
- 问题描述
- 复现步骤
- 错误日志
- 环境信息（OS、Docker 版本等）

---

## 🗺️ 路线图

### v1.0 (当前版本)

- [x] 完整的微服务架构
- [x] 用户、商铺、优惠券、订单、博客模块
- [x] Docker Compose 一键部署
- [x] 分布式追踪和监控
- [x] 完整文档
- [x] Dependabot 自动依赖更新
- [x] GitHub Actions CI/CD 流水线

### v2.0 (计划中)

- [ ] 前端界面（Vue 3 / React）
- [ ] 管理后台
- [ ] 单元测试覆盖率 >80%
- [ ] K8s 部署方案
- [ ] 文件上传迁移到 MinIO 对象存储

### v3.0 (未来)

- [ ] 微信小程序
- [ ] 移动端 App (Flutter)
- [ ] 实时推荐系统
- [ ] 智能客服（AI）
- [ ] 大数据分析

---

## 📊 项目统计

![GitHub stars](https://img.shields.io/github/stars/me-cs/dianping-gozero?style=social)
![GitHub forks](https://img.shields.io/github/forks/me-cs/dianping-gozero?style=social)
![GitHub watchers](https://img.shields.io/github/watchers/me-cs/dianping-gozero?style=social)

![GitHub commit activity](https://img.shields.io/github/commit-activity/m/me-cs/dianping-gozero)
![GitHub last commit](https://img.shields.io/github/last-commit/me-cs/dianping-gozero)
![GitHub contributors](https://img.shields.io/github/contributors/me-cs/dianping-gozero)

### Star History

[![Star History Chart](https://api.star-history.com/svg?repos=me-cs/dianping-gozero&type=Date)](https://star-history.com/#me-cs/dianping-gozero&Date)

---

## 🏷️ 标签

`go` `golang` `microservices` `go-zero` `grpc` `docker` `kubernetes` `redis` `mysql` `etcd` `jaeger` `prometheus` `grafana` `distributed-systems` `high-concurrency` `seckill` `dianping` `learning-project` `ai-generated` `claude-ai` `backend` `rest-api` `architecture` `devops` `cloud-native`

---

## 📄 许可证

MIT License - 详见 [LICENSE](LICENSE) 文件

Copyright (c) 2024 DianPing Clone Project

---

<div align="center">

**⭐️ 如果这个项目对你有帮助，请给个 Star！⭐️**

**🎉 欢迎 PR，一起让这个项目变得更好！🎉**

**☕ 欢迎赞助，支持开源！☕**

Made with ❤️ by Claude AI & Developers

[回到顶部](#-dianping-clone---go-微服务实战项目)

</div>
