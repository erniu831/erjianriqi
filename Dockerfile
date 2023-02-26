# 基础镜像
FROM golang:1.16-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将源代码复制到镜像中
COPY . .

# 构建二进制文件
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# 运行时镜像
FROM alpine:latest

# 安装运行应用程序所需的库
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /app

# 复制二进制文件到镜像中
COPY --from=builder /app/app .

# 暴露端口
EXPOSE 8080

# 运行应用程序
CMD ["./app"]
