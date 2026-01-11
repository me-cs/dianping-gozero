# 快速启动指南 - 5分钟部署点评微服务

本指南帮助你在5分钟内快速启动整个点评微服务系统。

## 前置要求 ✅

确保已安装：
- ✅ Docker Desktop（Windows/Mac）或 Docker Engine（Linux）
- ✅ Docker Compose

验证安装：
```bash
docker --version
docker-compose --version
```

## 一键启动 🚀

### Windows 用户

双击运行 `start.bat` 文件，或在命令行执行：

```cmd
cd C:\Users\13965\Desktop\heima\dianping\backend
start.bat
```

### Linux/Mac 用户

```bash
cd /path/to/dianping/backend
chmod +x start.sh
./start.sh
```

### 或者使用 Docker Compose

```bash
# 进入项目目录
cd C:\Users\13965\Desktop\heima\dianping\backend

# 启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps
```

## 验证服务 ✔️

等待约2-3分钟后，所有服务应该启动完成。

### 1. 检查服务状态

```bash
docker-compose ps
```

所有服务应显示 `Up` 或 `Up (healthy)` 状态。

### 2. 测试 API

```bash
# 测试 API 网关（Windows CMD）
curl http://localhost:8081/health

# 或在浏览器访问
http://localhost:8081
```

### 3. 访问监控界面

- **Grafana**: http://localhost:3000 (admin/admin)
- **Prometheus**: http://localhost:9090
- **Jaeger**: http://localhost:16686

## 服务端口速查 📋

### 业务服务
- API Gateway: `8081`
- User RPC: `8001`
- Shop RPC: `8002`
- Voucher RPC: `8003`
- Order RPC: `8004`
- Blog RPC: `8005`

### 基础设施
- MySQL: `3306` (root/root)
- Redis: `6379`
- etcd: `2379`

## 常用操作 🛠️

### 查看日志

```bash
# 查看所有服务日志
docker-compose logs -f

# 查看特定服务
docker-compose logs -f user-rpc
docker-compose logs -f api-gateway
```

### 重启服务

```bash
# 重启所有服务
docker-compose restart

# 重启单个服务
docker-compose restart user-rpc
```

### 停止服务

```bash
# 停止所有服务（保留数据）
docker-compose stop

# 停止并删除容器（保留数据）
docker-compose down

# 停止并删除所有数据（慎用！）
docker-compose down -v
```

## 问题排查 🔍

### 端口被占用

修改 `docker-compose.yml` 中冲突的端口：
```yaml
ports:
  - "18081:8081"  # 将8081改为18081
```

### 服务启动失败

```bash
# 查看错误日志
docker-compose logs [service-name]

# 重新构建
docker-compose build --no-cache [service-name]
docker-compose up -d [service-name]
```

### MySQL/Redis 连接失败

```bash
# 检查 MySQL
docker exec dianping-mysql mysqladmin ping -h localhost -uroot -proot

# 检查 Redis
docker exec dianping-redis redis-cli ping

# 重启基础设施
docker-compose restart mysql redis etcd
```

### 内存不足

只启动核心服务：
```bash
# 停止监控服务
docker-compose stop prometheus grafana jaeger

# 只启动必要服务
docker-compose up -d mysql redis etcd user-rpc shop-rpc voucher-rpc order-rpc blog-rpc api-gateway
```

## API 测试示例 📝

### 1. 发送验证码

```bash
curl -X POST http://localhost:8081/api/user/code \
  -H "Content-Type: application/json" \
  -d '{"phone": "13800138000"}'
```

### 2. 用户登录

```bash
curl -X POST http://localhost:8081/api/user/login \
  -H "Content-Type: application/json" \
  -d '{"phone": "13800138000", "code": "123456"}'
```

### 3. 查询商铺列表

```bash
curl -X GET "http://localhost:8081/api/shop/list?typeId=1&page=1&pageSize=10"
```

### 4. 查询优惠券

```bash
curl -X GET "http://localhost:8081/api/voucher/list?shopId=1"
```

## 下一步 ⏭️

服务启动成功后，你可以：

1. 📖 查看完整文档: [DOCKER_DEPLOY.md](./DOCKER_DEPLOY.md)
2. 🔧 调整配置文件: `rpc/*/etc/*.yaml` 和 `api/etc/*.yaml`
3. 📊 配置监控: 访问 Grafana 导入仪表板
4. 🧪 运行集成测试: 编写测试脚本调用 API
5. 📝 查看 API 文档: 使用 Postman 或 Swagger

## 获取帮助 💬

- 详细文档: [DOCKER_DEPLOY.md](./DOCKER_DEPLOY.md)
- 问题反馈: 提交 Issue 并附带日志
- 服务日志: `docker-compose logs -f [service]`

---

**提示**: 首次启动可能需要下载镜像，耗时较长。后续启动会快很多。

**注意**: 本环境仅用于开发和测试，生产部署请参考专门的部署文档。
