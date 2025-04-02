# 构建阶段
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 安装必要的工具
RUN apk add --no-cache make git gcc musl-dev

# 复制 Makefile 和依赖文件
COPY Makefile scaffold.mk ./
COPY go.mod go.sum ./

# 安装依赖
RUN go mod download

# 复制源代码
COPY . .

# 运行检查和测试
RUN make -f scaffold.mk vet
RUN make -f scaffold.mk lint
RUN make -f scaffold.mk test

# 生成swagger文档
RUN make -f scaffold.mk docs

# 构建应用
RUN make -f scaffold.mk build

# 运行阶段
FROM alpine:latest

WORKDIR /app

# 复制构建产物和swagger文档
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"] 