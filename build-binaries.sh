#!/bin/bash

echo "========================================"
echo "  Building Go Binaries on Host"
echo "========================================"
echo ""

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

echo "[1/6] Building User RPC service..."
cd rpc/user
go build -ldflags="-s -w" -o user-rpc user.go
if [ $? -ne 0 ]; then
    echo "❌ Failed to build user-rpc"
    exit 1
fi
echo "✓ user-rpc binary created"
cd ../..

echo ""
echo "[2/6] Building Shop RPC service..."
cd rpc/shop
go build -ldflags="-s -w" -o shop-rpc shop.go
if [ $? -ne 0 ]; then
    echo "❌ Failed to build shop-rpc"
    exit 1
fi
echo "✓ shop-rpc binary created"
cd ../..

echo ""
echo "[3/6] Building Voucher RPC service..."
cd rpc/voucher
go build -ldflags="-s -w" -o voucher-rpc voucher.go
if [ $? -ne 0 ]; then
    echo "❌ Failed to build voucher-rpc"
    exit 1
fi
echo "✓ voucher-rpc binary created"
cd ../..

echo ""
echo "[4/6] Building Order RPC service..."
cd rpc/order
go build -ldflags="-s -w" -o order-rpc order.go
if [ $? -ne 0 ]; then
    echo "❌ Failed to build order-rpc"
    exit 1
fi
echo "✓ order-rpc binary created"
cd ../..

echo ""
echo "[5/6] Building Blog RPC service..."
cd rpc/blog
go build -ldflags="-s -w" -o blog-rpc blog.go
if [ $? -ne 0 ]; then
    echo "❌ Failed to build blog-rpc"
    exit 1
fi
echo "✓ blog-rpc binary created"
cd ../..

echo ""
echo "[6/6] Building API Gateway..."
cd api
go build -ldflags="-s -w" -o dianping-api dianping.go
if [ $? -ne 0 ]; then
    echo "❌ Failed to build api-gateway"
    exit 1
fi
echo "✓ dianping-api binary created"
cd ..

echo ""
echo "========================================"
echo "  All binaries built successfully!"
echo "========================================"
echo ""
