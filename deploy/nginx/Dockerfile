FROM nginx:1.18-alpine

# 复制自定义的 nginx 配置文件
COPY deploy/nginx/nginx.conf /etc/nginx/nginx.conf

# 复制前端静态文件
# 注意：需要在构建时将前端文件复制到 deploy/nginx/html 目录
COPY deploy/nginx/html /usr/share/nginx/html

# 暴露 80 端口
EXPOSE 80

# 使用默认的 nginx 启动命令
CMD ["nginx", "-g", "daemon off;"]
