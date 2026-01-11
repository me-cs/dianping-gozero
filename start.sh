#!/bin/bash

echo "========================================"
echo "  Dianping Microservices Startup"
echo "========================================"
echo ""

# Configure data directory
# Priority: 1. Command line argument  2. Environment variable  3. Default value
if [[ -n "$2" && "$1" != "--reset" ]]; then
    # If second argument is provided and first is not --reset
    export DATA_DIR="$2"
elif [[ -n "$1" && "$1" != "--reset" ]]; then
    # If first argument is provided and it's not --reset
    export DATA_DIR="$1"
elif [[ -z "$DATA_DIR" ]]; then
    # If DATA_DIR environment variable is not set, use default
    export DATA_DIR="/var/lib/dianping"
fi

echo "📁 Data directory: $DATA_DIR"
echo ""

# Create data directories if they don't exist
echo "Ensuring data directories exist..."
mkdir -p "$DATA_DIR"/{mysql,redis,etcd,prometheus,grafana,jaeger}

# Set proper permissions (if running as root or with sudo)
if [[ $EUID -eq 0 ]]; then
    echo "Setting directory permissions..."
    chmod -R 755 "$DATA_DIR"
    # For services running as specific users
    chown -R 1001:1001 "$DATA_DIR/etcd" 2>/dev/null || true        # etcd (bitnami default user)
    chown -R 472:472 "$DATA_DIR/grafana" 2>/dev/null || true        # grafana
    chown -R 65534:65534 "$DATA_DIR/prometheus" 2>/dev/null || true # prometheus (nobody)
else
    echo "⚠️  Note: Not running as root. If you encounter permission issues:"
    echo "   For most services: sudo chown -R \$(id -u):\$(id -g) $DATA_DIR"
    echo "   For etcd: sudo chown -R 1001:1001 $DATA_DIR/etcd"
    echo "   For grafana: sudo chown -R 472:472 $DATA_DIR/grafana"
    echo "   For prometheus: sudo chown -R 65534:65534 $DATA_DIR/prometheus"
fi
echo "✓ Data directories ready"
echo ""

# Check for --reset flag
if [[ "$1" == "--reset" ]]; then
    echo "🔄 Reset mode enabled - will clean all data and start fresh"
    echo ""
    read -p "⚠️  This will DELETE all data in $DATA_DIR! Continue? (y/n): " choice
    if [[ "$choice" != "y" && "$choice" != "Y" ]]; then
        echo "Cancelled."
        exit 0
    fi

    # Auto-detect docker compose command
    if command -v docker-compose &> /dev/null; then
        DOCKER_COMPOSE="docker-compose"
    elif docker compose version &> /dev/null; then
        DOCKER_COMPOSE="docker compose"
    fi

    echo "Stopping and removing all containers..."
    $DOCKER_COMPOSE down

    echo "Removing all data..."
    rm -rf "$DATA_DIR"/*

    echo "Recreating data directories..."
    mkdir -p "$DATA_DIR"/{mysql,redis,etcd,prometheus,grafana,jaeger}

    echo "✓ Cleanup completed"
    echo ""
fi

# Auto-detect docker compose command
if command -v docker-compose &> /dev/null; then
    DOCKER_COMPOSE="docker-compose"
    echo "✓ Using: docker-compose"
elif docker compose version &> /dev/null; then
    DOCKER_COMPOSE="docker compose"
    echo "✓ Using: docker compose"
else
    echo "❌ Error: Neither 'docker-compose' nor 'docker compose' found!"
    echo "Please install Docker Compose first."
    exit 1
fi
echo ""

echo "[1/7] Starting infrastructure services..."
$DOCKER_COMPOSE up -d mysql redis etcd prometheus grafana jaeger

# 检查启动是否成功
if [ $? -ne 0 ]; then
    echo ""
    echo "❌ Error: Failed to start infrastructure services!"
    echo ""
    echo "可能的原因："
    echo "  1. Docker 镜像拉取失败（网络问题）"
    echo "  2. 端口被占用"
    echo "  3. Docker 服务未正常运行"
    echo ""
    echo "如果是网络问题导致镜像拉取失败，请配置 Docker 镜像加速器"
    echo "详见项目文档：DOCKER_MIRROR.md"
    echo ""
    echo "快速配置命令（需要 root 权限）："
    echo "  请参考项目根目录的 DOCKER_MIRROR.md 文件"
    echo ""
    exit 1
fi

echo "Waiting for infrastructure services to be healthy..."
sleep 10

echo ""
echo "[2/7] Checking etcd connection..."
ETCD_WAIT_TIME=0
ETCD_MAX_WAIT=60  # 最多等待60秒
until docker exec dianping-etcd etcdctl --endpoints=http://127.0.0.1:2379 endpoint health &> /dev/null; do
    if [ $ETCD_WAIT_TIME -ge $ETCD_MAX_WAIT ]; then
        echo "⚠️  Warning: etcd is taking too long to start. Checking logs..."
        docker logs --tail 20 dianping-etcd
        echo ""
        read -p "Continue anyway? (y/n): " choice
        if [[ "$choice" != "y" && "$choice" != "Y" ]]; then
            echo "Aborted. Please check etcd logs: docker logs dianping-etcd"
            exit 1
        fi
        break
    fi
    echo "etcd is not ready yet, waiting... (${ETCD_WAIT_TIME}s/${ETCD_MAX_WAIT}s)"
    sleep 3
    ETCD_WAIT_TIME=$((ETCD_WAIT_TIME + 3))
done
echo "etcd is ready!"

echo ""
echo "[3/7] Checking MySQL connection..."
MYSQL_WAIT_TIME=0
MYSQL_MAX_WAIT=60
until docker exec dianping-mysql mysqladmin ping -h localhost -uroot -proot &> /dev/null; do
    if [ $MYSQL_WAIT_TIME -ge $MYSQL_MAX_WAIT ]; then
        echo "⚠️  Warning: MySQL is taking too long to start."
        docker logs --tail 20 dianping-mysql
        exit 1
    fi
    echo "MySQL is not ready yet, waiting... (${MYSQL_WAIT_TIME}s/${MYSQL_MAX_WAIT}s)"
    sleep 5
    MYSQL_WAIT_TIME=$((MYSQL_WAIT_TIME + 5))
done
echo "MySQL is ready!"

echo ""
echo "[4/7] Checking Redis connection..."
REDIS_WAIT_TIME=0
REDIS_MAX_WAIT=30
until docker exec dianping-redis redis-cli ping &> /dev/null; do
    if [ $REDIS_WAIT_TIME -ge $REDIS_MAX_WAIT ]; then
        echo "⚠️  Warning: Redis is taking too long to start."
        docker logs --tail 20 dianping-redis
        exit 1
    fi
    echo "Redis is not ready yet, waiting... (${REDIS_WAIT_TIME}s/${REDIS_MAX_WAIT}s)"
    sleep 3
    REDIS_WAIT_TIME=$((REDIS_WAIT_TIME + 3))
done
echo "Redis is ready!"

echo ""
echo "[5/7] Building microservices..."
echo "Building Go binaries on host (much faster!)..."

# Go build environment variables
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

# Force IPv4 to avoid IPv6 network issues
export GODEBUG=netdns=go+2
export GOPROXY=https://goproxy.cn,https://goproxy.io,direct
export GOSUMDB=sum.golang.google.cn

# Build all Go binaries
for service in "rpc/user" "rpc/shop" "rpc/voucher" "rpc/order" "rpc/blog"; do
    svc_name=$(basename $service)
    binary_name="${svc_name}-rpc"
    if [ ! -f "$service/$binary_name" ]; then
        echo "Building $binary_name..."
        (cd $service && go build -ldflags="-s -w" -o $binary_name ${svc_name}.go)
        if [ $? -ne 0 ]; then
            echo "❌ Failed to build $binary_name"
            exit 1
        fi
        # Set executable permission immediately after building
        chmod +x "$service/$binary_name"
    else
        echo "✓ $binary_name binary already exists"
        # Ensure it has executable permission
        chmod +x "$service/$binary_name"
    fi
done

# Build API Gateway
if [ ! -f "api/dianping-api" ]; then
    echo "Building dianping-api..."
    (cd api && go build -ldflags="-s -w" -o dianping-api dianping.go)
    if [ $? -ne 0 ]; then
        echo "❌ Failed to build dianping-api"
        exit 1
    fi
    # Set executable permission
    chmod +x api/dianping-api
else
    echo "✓ dianping-api binary already exists"
    # Ensure it has executable permission
    chmod +x api/dianping-api
fi

echo "✓ All Go binaries ready"
echo ""

echo "Building Docker images..."
$DOCKER_COMPOSE build user-rpc shop-rpc voucher-rpc order-rpc blog-rpc api-gateway nginx

if [ $? -ne 0 ]; then
    echo "❌ Error: Failed to build Docker images!"
    echo "Please check the build logs above."
    exit 1
fi
echo "✓ Docker images built successfully"

echo ""
echo "[6/7] Starting RPC services..."
$DOCKER_COMPOSE up -d user-rpc shop-rpc voucher-rpc order-rpc blog-rpc

# 检查启动是否成功
if [ $? -ne 0 ]; then
    echo "❌ Error: Failed to start RPC services!"
    echo "Please check logs: $DOCKER_COMPOSE logs"
    exit 1
fi

echo "Waiting for RPC services to register..."
sleep 5

echo ""
echo "[7/8] Starting API Gateway..."
$DOCKER_COMPOSE up -d api-gateway

# 检查启动是否成功
if [ $? -ne 0 ]; then
    echo "❌ Error: Failed to start API Gateway!"
    echo "Please check logs: $DOCKER_COMPOSE logs api-gateway"
    exit 1
fi

echo ""
echo "[8/8] Starting Nginx Frontend..."
$DOCKER_COMPOSE up -d nginx

# 检查启动是否成功
if [ $? -ne 0 ]; then
    echo "❌ Error: Failed to start Nginx!"
    echo "Please check logs: $DOCKER_COMPOSE logs nginx"
    exit 1
fi

echo ""
echo "========================================"
echo "  All services started successfully!"
echo "========================================"
echo ""
echo "📁 Data directory: $DATA_DIR"
echo ""
echo "Service URLs:"
echo "  - Frontend (Nginx): http://localhost:80"
echo "  - API Gateway:      http://localhost:8081"
echo "  - Grafana:          http://localhost:3000 (admin/admin)"
echo "  - Prometheus:       http://localhost:9090"
echo "  - Jaeger UI:        http://localhost:16686"
echo "  - etcd:             http://localhost:2379"
echo ""
echo "Check service status:"
echo "  $DOCKER_COMPOSE ps"
echo ""
echo "View logs:"
echo "  $DOCKER_COMPOSE logs -f [service-name]"
echo ""
echo "Usage:"
echo "  ./start.sh                      # Normal start (default: /var/lib/dianping)"
echo "  ./start.sh /path/to/data        # Start with custom data directory"
echo "  DATA_DIR=/data ./start.sh       # Start with env variable"
echo "  ./start.sh --reset              # Clean start (delete all data)"
echo "  ./start.sh --reset /path        # Clean start with custom directory"
echo "  ./stop.sh                       # Stop all services"
echo "  ./build-binaries.sh             # Build Go binaries only (fast, for development)"
echo "  ./build-docker.sh               # Build binaries + Docker images (complete build)"
echo ""
