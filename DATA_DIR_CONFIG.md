# 数据目录配置说明

## 概述

从此版本开始，所有持久化数据的存储位置都是可配置的，数据不再默认存储在源代码目录下。

## 配置方式

有三种方式配置数据目录，优先级从高到低：

### 1. 命令行参数（最高优先级）

```bash
# 正常启动，指定数据目录
./start.sh /data/dianping

# 重置模式，指定数据目录
./start.sh --reset /data/dianping
```

### 2. 环境变量

```bash
# 方式A：直接设置环境变量
export DATA_DIR=/data/dianping
./start.sh

# 方式B：临时设置
DATA_DIR=/data/dianping ./start.sh

# 方式C：使用 .env 文件
cp .env.example .env
# 编辑 .env 文件，修改 DATA_DIR 值
vim .env
./start.sh
```

### 3. 默认值（最低优先级）

如果没有指定任何配置，默认使用 `/var/lib/dianping`

## 数据目录结构

无论使用哪种方式配置，start.sh 会自动创建以下子目录：

```
${DATA_DIR}/
├── mysql/         # MySQL 数据库文件
├── redis/         # Redis 持久化数据
├── etcd/          # etcd 配置和数据
├── prometheus/    # Prometheus 时序数据
├── grafana/       # Grafana 配置和仪表盘
└── jaeger/        # Jaeger 追踪数据
```

## 推荐配置

### 开发环境

```bash
# 使用家目录，方便访问
DATA_DIR=$HOME/dianping-data ./start.sh
```

### 测试环境

```bash
# 使用临时目录，方便清理
DATA_DIR=/tmp/dianping ./start.sh
```

### 生产环境

```bash
# 使用独立数据分区，确保性能和安全
DATA_DIR=/data/dianping ./start.sh
```

或者配置 .env 文件：

```bash
# 复制示例文件
cp .env.example .env

# 编辑配置
vim .env
# 修改: DATA_DIR=/data/dianping

# 启动
./start.sh
```

## 权限问题

### 非 root 用户

如果使用非 root 用户运行，确保对数据目录有写权限：

```bash
# 创建数据目录
sudo mkdir -p /data/dianping

# 设置权限
sudo chown -R $(id -u):$(id -g) /data/dianping

# 启动服务
DATA_DIR=/data/dianping ./start.sh
```

### root 用户

使用 root 用户启动时，脚本会自动设置正确的权限：
- etcd: 1001:1001 (bitnami 默认用户)
- Grafana: 472:472
- Prometheus: 65534:65534
- 其他: 755


## 数据迁移

### 从旧版本迁移

如果你之前使用的是源代码目录下的 `data/` 目录：

```bash
# 1. 停止服务
./stop.sh

# 2. 复制数据到新位置
sudo mkdir -p /var/lib/dianping
sudo cp -r ./data/* /var/lib/dianping/

# 3. 设置权限
sudo chown -R $(id -u):$(id -g) /var/lib/dianping

# 4. 启动服务（使用新位置）
DATA_DIR=/var/lib/dianping ./start.sh

# 5. 验证服务正常后，删除旧数据（可选）
rm -rf ./data
```

### 更改数据目录位置

```bash
# 1. 停止服务
./stop.sh

# 2. 移动数据
sudo mv /var/lib/dianping /data/dianping

# 3. 使用新位置启动
DATA_DIR=/data/dianping ./start.sh
```

## 清理数据

### 仅清理数据，不删除目录

```bash
./start.sh --reset
```

这会：
1. 停止所有容器
2. 删除 `$DATA_DIR` 下的所有数据
3. 重新创建空的数据目录
4. 启动服务

### 完全删除

```bash
# 停止服务
./stop.sh

# 删除数据目录
sudo rm -rf /var/lib/dianping
```

## 注意事项

1. **数据目录必须有足够的磁盘空间**
   - MySQL 数据库会随着使用增长
   - Prometheus 时序数据可能占用大量空间
   - 建议至少预留 10GB 空间

2. **备份重要数据**
   - 在执行 `--reset` 操作前，务必备份重要数据
   - 定期备份 MySQL 数据库

3. **权限问题**
   - 确保数据目录有正确的读写权限
   - 某些服务需要特定的用户权限

4. **路径要求**
   - 使用绝对路径，不要使用相对路径
   - 避免使用含有空格的路径

## 示例场景

### 场景1：开发人员本地测试

```bash
# 使用家目录，方便查看和清理
./start.sh ~/dianping-test
```

### 场景2：CI/CD 环境

```bash
# 使用临时目录，每次构建后自动清理
DATA_DIR=/tmp/dianping-ci-${BUILD_ID} ./start.sh
```

### 场景3：多环境部署

```bash
# 开发环境
DATA_DIR=/data/dianping-dev ./start.sh

# 测试环境
DATA_DIR=/data/dianping-test ./start.sh

# 生产环境
DATA_DIR=/data/dianping-prod ./start.sh
```

### 场景4：数据隔离

```bash
# 为不同的测试分支使用不同的数据目录
DATA_DIR=/data/dianping-feature-1 ./start.sh
DATA_DIR=/data/dianping-feature-2 ./start.sh
```
