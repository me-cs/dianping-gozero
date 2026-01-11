# Docker 镜像加速配置指南

由于国内网络环境的限制，直接从 Docker Hub 拉取镜像可能会非常慢或失败。本文档提供 Docker 镜像加速器的配置方法。

> ⚠️ **注意**: 本项目不会自动修改你的 Docker 配置，请根据需要手动配置。

## 🌏 为什么需要镜像加速

从 Docker Hub (`registry-1.docker.io`) 拉取镜像时，在中国大陆地区经常遇到：
- ❌ 连接超时：`context deadline exceeded`
- ❌ 速度极慢：下载几百 MB 需要几小时
- ❌ 连接中断：下载到一半失败

配置镜像加速后，可以：
- ✅ 从国内镜像源拉取
- ✅ 速度提升 10-100 倍
- ✅ 稳定性大幅提升

## 🚀 快速配置（推荐）

### 方法一：使用 DaoCloud 镜像（推荐）

```bash
# 1. 创建或编辑 Docker 配置文件
sudo mkdir -p /etc/docker

sudo tee /etc/docker/daemon.json > /dev/null << 'EOF'
{
  "registry-mirrors": [
    "https://docker.m.daocloud.io"
  ]
}
EOF

# 2. 重启 Docker 服务
sudo systemctl daemon-reload
sudo systemctl restart docker

# 3. 验证配置
docker info | grep -A 5 "Registry Mirrors"
```

### 方法二：使用多个镜像源（更稳定）

```bash
# 配置多个镜像源作为备份
sudo tee /etc/docker/daemon.json > /dev/null << 'EOF'
{
  "registry-mirrors": [
    "https://docker.m.daocloud.io",
    "https://docker.mirrors.sjtug.sjtu.edu.cn",
    "https://docker.nju.edu.cn",
    "https://mirror.ccs.tencentyun.com"
  ],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m",
    "max-file": "3"
  }
}
EOF

sudo systemctl daemon-reload
sudo systemctl restart docker
```

## 📋 可用的国内镜像源（2024 年最新）

| 镜像源 | 地址 | 速度 | 稳定性 |
|--------|------|------|--------|
| DaoCloud | `https://docker.m.daocloud.io` | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| 上海交大 | `https://docker.mirrors.sjtug.sjtu.edu.cn` | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| 南京大学 | `https://docker.nju.edu.cn` | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| 腾讯云 | `https://mirror.ccs.tencentyun.com` | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| 网易云 | `https://hub-mirror.c.163.com` | ⭐⭐⭐ | ⭐⭐⭐ |

> 💡 **提示**: 推荐配置多个镜像源，Docker 会自动选择可用的源。

## 🔍 验证配置是否生效

### 1. 查看配置

```bash
# 查看 Docker 配置信息
docker info | grep -A 10 "Registry Mirrors"

# 应该显示类似：
# Registry Mirrors:
#  https://docker.m.daocloud.io/
```

### 2. 测试拉取镜像

```bash
# 清理测试镜像（如果存在）
docker rmi hello-world 2>/dev/null

# 拉取测试镜像
time docker pull hello-world

# 如果几秒内完成，说明配置成功
```

### 3. 测试项目镜像

```bash
# 测试拉取项目所需的镜像
docker pull mysql:8.0
docker pull redis:7-alpine
docker pull bitnami/etcd:latest
```

## 📝 完整配置示例

如果你的 `/etc/docker/daemon.json` 文件已存在其他配置，需要合并配置：

```json
{
  "registry-mirrors": [
    "https://docker.m.daocloud.io",
    "https://docker.mirrors.sjtug.sjtu.edu.cn"
  ],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m",
    "max-file": "3"
  },
  "storage-driver": "overlay2",
  "insecure-registries": [],
  "debug": false
}
```

## 🛠️ 不同系统的配置方法

### CentOS / RHEL / Rocky Linux

```bash
# 1. 编辑配置文件
sudo vi /etc/docker/daemon.json

# 2. 重启服务
sudo systemctl daemon-reload
sudo systemctl restart docker

# 3. 设置开机自启
sudo systemctl enable docker
```

### Ubuntu / Debian

```bash
# 1. 编辑配置文件
sudo nano /etc/docker/daemon.json

# 2. 重启服务
sudo systemctl daemon-reload
sudo systemctl restart docker
```

### Docker Desktop (Windows / Mac)

1. 打开 Docker Desktop
2. 点击右上角设置图标 ⚙️
3. 选择 `Docker Engine`
4. 在 JSON 配置中添加：
```json
{
  "registry-mirrors": [
    "https://docker.m.daocloud.io"
  ]
}
```
5. 点击 `Apply & Restart`

## 🔧 故障排查

### 问题 1：配置后仍然很慢

**解决方案**：
```bash
# 1. 检查镜像源是否可访问
curl -I https://docker.m.daocloud.io

# 2. 尝试其他镜像源
# 修改 daemon.json，将慢的源删除或移到后面

# 3. 清理 Docker 缓存
docker system prune -a
```

### 问题 2：Docker 重启失败

**解决方案**：
```bash
# 1. 检查 JSON 格式是否正确
cat /etc/docker/daemon.json | python -m json.tool

# 2. 查看 Docker 日志
sudo journalctl -u docker -f

# 3. 如果配置有误，恢复默认配置
sudo mv /etc/docker/daemon.json /etc/docker/daemon.json.bak
sudo systemctl restart docker
```

### 问题 3：某些镜像源不可用

**解决方案**：
```bash
# 测试各个镜像源的连通性
for mirror in \
  "https://docker.m.daocloud.io" \
  "https://docker.mirrors.sjtug.sjtu.edu.cn" \
  "https://docker.nju.edu.cn"; do
  echo "Testing $mirror"
  curl -I --connect-timeout 5 $mirror
  echo ""
done

# 只保留可用的镜像源
```

## 📊 性能对比

配置前后拉取 `mysql:8.0` (约 500MB) 的时间对比：

| 方式 | 下载时间 | 速度 |
|------|---------|------|
| 直连 Docker Hub | 30-60 分钟 或超时 | < 1 MB/s |
| DaoCloud 镜像 | 2-5 分钟 | 10-20 MB/s |
| 上海交大镜像 | 2-5 分钟 | 10-20 MB/s |

## 🔒 安全说明

### 镜像源安全性

- ✅ 官方认可的镜像源（如 DaoCloud、阿里云）定期同步官方镜像
- ✅ 镜像内容通过 SHA256 校验，确保完整性
- ⚠️ 避免使用来源不明的镜像源

### 验证镜像完整性

```bash
# 拉取镜像后验证
docker pull mysql:8.0
docker inspect mysql:8.0 | grep RepoDigest
```

## 💡 最佳实践

### 1. 推荐配置

```json
{
  "registry-mirrors": [
    "https://docker.m.daocloud.io",
    "https://docker.mirrors.sjtug.sjtu.edu.cn"
  ],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m",
    "max-file": "3"
  },
  "max-concurrent-downloads": 10,
  "max-concurrent-uploads": 5
}
```

### 2. 生产环境建议

- 使用企业级镜像源（阿里云、腾讯云）
- 配置私有镜像仓库（Harbor）
- 定期更新镜像源列表

### 3. 开发环境建议

- 配置多个公共镜像源
- 预拉取常用镜像
- 使用 Docker Compose 加速

## 🆘 还是无法拉取镜像？

如果配置镜像加速后仍然无法拉取，可以尝试：

### 方案 1：使用代理

```bash
# 配置 Docker 使用 HTTP 代理
sudo mkdir -p /etc/systemd/system/docker.service.d
sudo tee /etc/systemd/system/docker.service.d/http-proxy.conf > /dev/null << 'EOF'
[Service]
Environment="HTTP_PROXY=http://proxy.example.com:80"
Environment="HTTPS_PROXY=http://proxy.example.com:80"
Environment="NO_PROXY=localhost,127.0.0.1"
EOF

sudo systemctl daemon-reload
sudo systemctl restart docker
```

### 方案 2：离线镜像

如果网络完全不可用，可以使用离线镜像：

```bash
# 在有网络的机器上导出镜像
docker save mysql:8.0 -o mysql-8.0.tar
docker save redis:7-alpine -o redis-7-alpine.tar

# 传输到目标机器后导入
docker load -i mysql-8.0.tar
docker load -i redis-7-alpine.tar
```

### 方案 3：使用项目提供的镜像包

如果我们提供了镜像包，解压后导入：

```bash
# 批量导入镜像
for image in images/*.tar; do
  docker load -i "$image"
done
```

## 📚 相关资源

- [Docker 官方文档 - Registry Mirror](https://docs.docker.com/registry/recipes/mirror/)
- [DaoCloud 镜像加速器](https://www.daocloud.io/mirror)
- [上海交大 Docker 镜像源](https://mirrors.sjtug.sjtu.edu.cn/)

## 🔄 更新日志

- **2024-01**: 添加最新可用的镜像源
- **2024-01**: 更新配置示例和最佳实践

---

**配置完成后，返回项目根目录运行 `./start.sh` 启动服务！**
