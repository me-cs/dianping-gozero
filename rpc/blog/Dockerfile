FROM alpine:latest

# Install ca-certificates and timezone data
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app

# Copy pre-built binary and config from host
COPY rpc/blog/blog-rpc /app/blog-rpc
COPY rpc/blog/etc /app/etc

# Use sh to ensure executable permission at runtime
CMD ["/bin/sh", "-c", "chmod +x /app/blog-rpc && exec /app/blog-rpc -f etc/blog.yaml"]
