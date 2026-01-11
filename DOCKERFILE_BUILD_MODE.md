# Dockerfile 构建模式说明

## 当前模式：宿主机构建（开发/测试阶段）

当前所有 Dockerfile 使用**宿主机预构建模式**，即在宿主机上先编译 Go 二进制文件，然后复制到 Docker 镜像中。

### 优点
- 构建速度快，尤其是在配置较差的服务器上
- 可以利用宿主机的 Go 模块缓存
- 调试方便

### 使用方法
```bash
./build-binaries.sh  # 只编译 Go 二进制文件（快速，用于开发测试）
./build-docker.sh    # 编译二进制 + 构建 Docker 镜像（完整构建）
./start.sh           # 启动时也会自动检查并构建
```

---

## 未来模式：多阶段构建（生产/开源阶段）

在开源或生产部署时，应该改回**多阶段构建模式**，让 Docker 自己处理编译过程。

### 原始 Dockerfile 模板（多阶段构建）

```dockerfile
FROM golang:1.25-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

# Copy go mod files
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy source code
COPY . .

# Build (根据不同服务修改路径和输出名称)
WORKDIR /build/rpc/user
RUN go build -ldflags="-s -w" -o /app/user user.go

FROM alpine:latest

# Install ca-certificates and timezone data
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/user /app/user
COPY --from=builder /build/rpc/user/etc /app/etc

CMD ["./user", "-f", "etc/user.yaml"]
```

### 各服务的构建路径

| 服务 | 构建路径 | 二进制名 | 主文件 | 配置文件 |
|-----|---------|---------|--------|---------|
| user-rpc | /build/rpc/user | user-rpc | user.go | etc/user.yaml |
| shop-rpc | /build/rpc/shop | shop-rpc | shop.go | etc/shop.yaml |
| voucher-rpc | /build/rpc/voucher | voucher-rpc | voucher.go | etc/voucher.yaml |
| order-rpc | /build/rpc/order | order-rpc | order.go | etc/order.yaml |
| blog-rpc | /build/rpc/blog | blog-rpc | blog.go | etc/blog.yaml |
| api-gateway | /build/api | dianping-api | dianping.go | etc/dianping-api.yaml |

### 切换回多阶段构建的步骤

1. 将各服务的 Dockerfile 改回上述多阶段构建模板
2. 修改 `build-docker.sh`，移除宿主机构建部分，直接调用 docker compose build
3. 修改 `start.sh`，移除宿主机构建部分
4. 更新 `.gitignore`（可选，但建议保留二进制文件的忽略规则）

### 何时切换

- ✅ **开发测试阶段**：使用当前的宿主机构建模式（快速）
- ✅ **开源发布时**：切换到多阶段构建（标准、可移植）
- ✅ **生产部署时**：切换到多阶段构建（安全、隔离）
- ✅ **CI/CD 流水线**：使用多阶段构建（自包含）

---

## 注意事项

1. 宿主机构建需要确保：
   - 已安装 Go 1.21+
   - 设置了 `GOOS=linux GOARCH=amd64`（如果宿主机不是 Linux）
   - 有网络访问权限下载 Go 依赖

2. 多阶段构建需要确保：
   - Docker 版本支持多阶段构建（17.05+）
   - 有足够的磁盘空间和内存
   - 网络可以访问 Go 代理

## 当前文件状态

- ✅ 所有 Dockerfile 已改为宿主机构建模式
- ✅ build-docker.sh 编译二进制 + 构建镜像（完整构建流程）
- ✅ build-binaries.sh 只编译二进制（快速开发测试）
- ✅ start.sh 已更新为宿主机构建
- ✅ .gitignore 已添加二进制文件忽略规则

## 构建脚本说明

### build-binaries.sh
- **用途**：快速编译 Go 二进制文件
- **适用场景**：本地开发、测试二进制是否能编译通过
- **输出**：生成 `*-rpc` 和 `dianping-api` 二进制文件
- **速度**：快（只编译，不构建镜像）

### build-docker.sh
- **用途**：完整构建流程（编译 + Docker 镜像）
- **适用场景**：准备部署、更新镜像
- **输出**：二进制文件 + Docker 镜像
- **速度**：较慢（包含 Docker 镜像构建）

### start.sh
- **用途**：一键启动所有服务
- **特点**：自动检查并构建缺失的二进制和镜像
- **推荐**：日常开发使用
