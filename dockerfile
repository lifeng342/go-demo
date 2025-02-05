# 使用指定平台的镜像
FROM --platform=linux/amd64 golang:1.23

# 更新并安装必需的包
RUN apt-get update && apt-get install -y git

# 设置工作目录
WORKDIR /code

# 将当前目录的内容复制到 `/code` 目录下
COPY . .

# 配置 Go 环境变量
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go env -w GO111MODULE=on

# 整理 Go 依赖
RUN go mod tidy

# 运行 build.sh 文件
RUN chmod +x build.sh
CMD ["./build.sh"]