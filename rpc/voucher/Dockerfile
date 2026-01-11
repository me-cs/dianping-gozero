FROM alpine:latest

# Install ca-certificates and timezone data
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app

# Copy pre-built binary and config from host
COPY rpc/voucher/voucher-rpc /app/voucher-rpc
COPY rpc/voucher/etc /app/etc

# Use sh to ensure executable permission at runtime
CMD ["/bin/sh", "-c", "chmod +x /app/voucher-rpc && exec /app/voucher-rpc -f etc/voucher.yaml"]
