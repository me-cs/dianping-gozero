# 点评项目 Docker 部署指南

本文档介绍如何使用 Docker 和 Docker Compose 部署点评微服务项目的开发和测试环境。

## 📋 目录

- [系统要求](#系统要求)
- [服务架构](#服务架构)
- [快速开始](#快速开始)
- [详细说明](#详细说明)
- [常见问题](#常见问题)
- [监控和调试](#监控和调试)

## 🖥️ 系统要求

- Docker 20.10+
- Docker Compose 2.0+
- 至少 4GB 可用内存
- 至少 10GB 可用磁盘空间

### Windows 用户

```bash
# 安装 Docker Desktop for Windows
# 下载地址: https://www.docker.com/products/docker-desktop/

# 确认安装成功
docker --version
docker-compose --version
```

### Linux 用户

```bash
# 安装 Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# 安装 Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

## 🏗️ 服务架构

### 基础设施服务

| 服务 | 端口 | 说明 |
|------|------|------|
| MySQL | 3306 | 数据库服务 |
| Redis | 6379 | 缓存服务 |
| etcd | 2379, 2380 | 服务发现 |

### 业务 RPC 服务

| 服务 | RPC端口 | Metrics端口 | 说明 |
|------|---------|-------------|------|
| user-rpc | 8001 | 9001 | 用户服务 |
| shop-rpc | 8002 | 9002 | 商铺服务 |
| voucher-rpc | 8003 | 9003 | 优惠券服务 |
| order-rpc | 8004 | 9004 | 订单服务 |
| blog-rpc | 8005 | 9005 | 博客服务 |

### API 网关

| 服务 | 端口 | Metrics端口 | 说明 |
|------|------|-------------|------|
| api-gateway | 8081 | 9091 | HTTP API网关 |

### 监控服务（可选）

| 服务 | 端口 | 说明 |
|------|------|------|
| Prometheus | 9090 | 指标采集 |
| Grafana | 3000 | 可视化仪表板 (admin/admin) |
| Jaeger | 16686 | 分布式追踪 UI |
| Jaeger Collector | 14268 | 追踪数据收集器 |

## 🚀 快速开始

### 方式一：使用启动脚本（推荐）

```bash
# 进入项目目录
cd C:\Users\13965\Desktop\heima\dianping\backend

# 给脚本添加执行权限（Linux/Mac）
chmod +x start.sh

# 运行启动脚本
./start.sh
```

脚本会按顺序启动：
1. 基础设施服务（MySQL, Redis, etcd）
2. 等待服务健康检查通过
3. 启动所有 RPC 服务
4. 启动 API 网关

### 方式二：手动启动

```bash
# 1. 启动所有服务
docker-compose up -d

# 2. 查看服务状态
docker-compose ps

# 3. 查看服务日志
docker-compose logs -f
```

### 方式三：分步启动

```bash
# 1. 仅启动基础设施
docker-compose up -d mysql redis etcd

# 2. 等待基础设施就绪（约30秒）
docker-compose ps

# 3. 启动 RPC 服务
docker-compose up -d user-rpc shop-rpc voucher-rpc order-rpc blog-rpc

# 4. 启动 API 网关
docker-compose up -d api-gateway

# 5. 启动监控服务（可选）
docker-compose up -d prometheus grafana jaeger
```

## 📝 详细说明

### 数据持久化

所有数据都存储在 `data/` 目录下：

```
data/
├── mysql/          # MySQL 数据文件
├── redis/          # Redis 持久化文件
├── etcd/           # etcd 数据
├── prometheus/     # Prometheus 时序数据
├── grafana/        # Grafana 配置和数据
└── jaeger/         # Jaeger 追踪数据
```

### 配置文件

服务配置文件位于各自的 `etc/` 目录：

- `rpc/user/etc/user.yaml`
- `rpc/shop/etc/shop.yaml`
- `rpc/voucher/etc/voucher.yaml`
- `rpc/order/etc/order.yaml`
- `rpc/blog/etc/blog.yaml`
- `api/etc/dianping-api.yaml`

### 网络配置

所有服务都运行在 `dianping-network` 桥接网络中，可以通过服务名互相访问：

- 容器内访问 MySQL: `mysql:3306`
- 容器内访问 Redis: `redis:6379`
- 容器内访问 etcd: `etcd:2379`
- 宿主机访问: 使用 `localhost` 加对应端口

## 🔧 常用命令

### 查看服务状态

```bash
# 查看所有服务状态
docker-compose ps

# 查看特定服务状态
docker-compose ps user-rpc
```

### 查看日志

```bash
# 查看所有服务日志
docker-compose logs -f

# 查看特定服务日志
docker-compose logs -f user-rpc

# 查看最近100行日志
docker-compose logs --tail=100 user-rpc

# 查看多个服务日志
docker-compose logs -f user-rpc shop-rpc
```

### 重启服务

```bash
# 重启所有服务
docker-compose restart

# 重启特定服务
docker-compose restart user-rpc

# 重新构建并重启
docker-compose up -d --build user-rpc
```

### 停止服务

```bash
# 停止所有服务
docker-compose stop

# 停止特定服务
docker-compose stop user-rpc

# 停止并删除容器（保留数据）
docker-compose down

# 停止并删除容器和数据卷（危险！）
docker-compose down -v
```

### 进入容器

```bash
# 进入 MySQL 容器
docker exec -it dianping-mysql mysql -uroot -proot hmdp

# 进入 Redis 容器
docker exec -it dianping-redis redis-cli

# 进入业务服务容器
docker exec -it dianping-user-rpc sh
```

### 清理资源

```bash
# 清理未使用的镜像
docker image prune -a

# 清理未使用的容器
docker container prune

# 清理未使用的卷
docker volume prune

# 清理所有未使用的资源
docker system prune -a --volumes
```

## ❓ 常见问题

### 1. 端口冲突

如果某个端口已被占用，可以修改 `docker-compose.yml` 中的端口映射：

```yaml
ports:
  - "13306:3306"  # 将宿主机端口改为 13306
```

### 2. 服务启动失败

```bash
# 查看详细错误信息
docker-compose logs [service-name]

# 检查健康状态
docker-compose ps

# 重新构建镜像
docker-compose build --no-cache [service-name]
docker-compose up -d [service-name]
```

### 3. MySQL 连接失败

```bash
# 检查 MySQL 是否就绪
docker exec dianping-mysql mysqladmin ping -h localhost -uroot -proot

# 查看 MySQL 日志
docker-compose logs mysql

# 重启 MySQL
docker-compose restart mysql
```

### 4. etcd 服务发现问题

```bash
# 检查 etcd 健康状态
docker exec dianping-etcd etcdctl endpoint health

# 查看已注册的服务
docker exec dianping-etcd etcdctl get "" --prefix --keys-only

# 删除旧的服务注册（如果需要）
docker exec dianping-etcd etcdctl del /user.rpc --prefix
```

### 5. 内存不足

如果遇到内存不足，可以：

- 关闭不必要的监控服务：`docker-compose stop prometheus grafana jaeger`
- 增加 Docker Desktop 的内存限制（Settings -> Resources）
- 分批启动服务

### 6. 构建速度慢

```bash
# 使用国内镜像加速（已在 Dockerfile 中配置）
ENV GOPROXY https://goproxy.cn,direct

# 使用 Docker 构建缓存
docker-compose build --parallel
```

## 📊 监控和调试

### Prometheus（指标监控）

访问: http://localhost:9090

查询示例：
- HTTP 请求总数: `http_request_total`
- RPC 调用延迟: `rpc_duration_seconds`
- 系统内存使用: `go_memstats_alloc_bytes`

### Grafana（可视化）

访问: http://localhost:3000
默认账号: `admin` / `admin`

步骤：
1. 添加 Prometheus 数据源（URL: http://prometheus:9090）
2. 导入 Go 应用仪表板（Dashboard ID: 6671）
3. 创建自定义面板

### Jaeger（分布式追踪）

访问: http://localhost:16686

功能：
- 查看服务调用链路
- 分析请求性能瓶颈
- 调试跨服务问题

### 健康检查

所有服务都配置了健康检查，可通过以下方式查看：

```bash
# Docker Compose 健康状态
docker-compose ps

# 单个容器健康状态
docker inspect --format='{{.State.Health.Status}}' dianping-mysql
```

## 🧪 测试 API

### 使用 curl

```bash
# 测试 API 网关健康检查
curl http://localhost:8081/health

# 测试用户登录
curl -X POST http://localhost:8081/api/user/login \
  -H "Content-Type: application/json" \
  -d '{"phone": "13800138000", "code": "123456"}'
```

### 使用 Postman

导入 API 文档并配置：
- Base URL: `http://localhost:8081`
- Headers: `Content-Type: application/json`

## 🔐 安全注意事项

⚠️ **生产环境请修改默认密码和配置！**

- MySQL root 密码: `root`（修改 `docker-compose.yml` 中的 `MYSQL_ROOT_PASSWORD`）
- Grafana 密码: `admin/admin`（首次登录后强制修改）
- Redis 密码: 无（建议在生产环境启用 `requirepass`）

## 📚 相关文档

- [go-zero 官方文档](https://go-zero.dev/)
- [Docker 官方文档](https://docs.docker.com/)
- [Docker Compose 文档](https://docs.docker.com/compose/)
- [etcd 文档](https://etcd.io/docs/)

## 🤝 支持

如遇问题，请：
1. 查看本文档的常见问题部分
2. 检查服务日志: `docker-compose logs [service]`
3. 提交 Issue 并附上错误日志

## 📄 许可证

本项目仅用于学习和测试目的。
