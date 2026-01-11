FROM alpine:latest

# Install ca-certificates and timezone data
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app

# Copy pre-built binary and config from host
COPY api/dianping-api /app/dianping-api
COPY api/etc /app/etc

# Use sh to ensure executable permission at runtime
CMD ["/bin/sh", "-c", "chmod +x /app/dianping-api && exec /app/dianping-api -f etc/dianping-api.yaml"]
