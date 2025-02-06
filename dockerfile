# 使用指定平台的镜像
FROM --platform=linux/amd64 golang:1.23

# 设置工作目录
WORKDIR /app

# 更新并安装必需的包
# RUN apt-get update && apt-get install -y git openssh-client ssh sshpass
#
# RUN mkdir -p -m 0700 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts
# RUN git config --global url."ssh://git@github.com/lifeng342/".insteadOf "https://github.com/lifeng342"
ENV GOPRIVATE=github.com/lifeng342/**
RUN git config --global url."https://${GIT_USER}:${GIT_ACCESS_TOKEN}@github.com/lifeng342/".insteadOf "https://github.com/lifeng342/"
ENV GOPROXY=goproxy.cn

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# 将当前目录的内容复制到 `/code` 目录下
COPY ./ /app

# 配置 Go 环境变量
# RUN go env -w GOPROXY=https://goproxy.io,direct
# RUN go env -w GO111MODULE=on

# 整理 Go 依赖
# RUN go mod tidy

# 运行 build.sh 文件
RUN chmod +x build.sh
CMD ["./build.sh"]