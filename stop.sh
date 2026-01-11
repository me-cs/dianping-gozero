#!/bin/bash

echo "========================================"
echo "  Dianping Microservices Shutdown"
echo "========================================"
echo ""

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

echo "Stopping all services..."
$DOCKER_COMPOSE stop

if [ $? -ne 0 ]; then
    echo "❌ Error: Failed to stop services!"
    exit 1
fi

echo ""
echo "========================================"
echo "  All services stopped successfully!"
echo "========================================"
echo ""
echo "Note: Data volumes are preserved."
echo ""
echo "To restart services:"
echo "  ./start.sh"
echo ""
echo "To remove all data and start fresh:"
echo "  ./start.sh --reset"
echo ""
