######## build stage ########
FROM golang:1.18 AS builder

COPY . /src
WORKDIR /src

# build，并且把所有程序都统一命名为app 方便下面的启动
RUN cd cmd/timeline && GOPROXY=https://goproxy.cn mkdir -p bin/ && go build -o ./bin/ ./... \
    && cd bin && mv timeline app

FROM debian:stable-20220912-slim

# 为了加快速度，不在进行额外的依赖安装
#RUN apt-get update && apt-get install -y --no-install-recommends \
#    ca-certificates  \
#    netbase \
#    && rm -rf /var/lib/apt/lists/ \
#    && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/cmd/timeline/bin /app

WORKDIR /app

# http & grpc
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

# ${app} 是环境变量，dockerfile 中 args 只在build time生效，CMD是 runtime
# see: https://github.com/moby/moby/issues/34772
CMD exec ./app -conf /data/conf/config.timeline.yaml