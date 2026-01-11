FROM alpine:latest

# Install ca-certificates and timezone data
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app

# Copy pre-built binary and config from host
COPY rpc/order/order-rpc /app/order-rpc
COPY rpc/order/etc /app/etc

# Use sh to ensure executable permission at runtime
CMD ["/bin/sh", "-c", "chmod +x /app/order-rpc && exec /app/order-rpc -f etc/order.yaml"]
