FROM alpine:latest

# Install ca-certificates and timezone data
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app

# Copy pre-built binary and config from host
COPY rpc/user/user-rpc /app/user-rpc
COPY rpc/user/etc /app/etc

# Use sh to ensure executable permission at runtime
CMD ["/bin/sh", "-c", "chmod +x /app/user-rpc && exec /app/user-rpc -f etc/user.yaml"]
