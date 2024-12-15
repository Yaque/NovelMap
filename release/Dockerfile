# 基础镜像
FROM alpine:3.12
RUN mkdir -p "/data/app/data"
# docker build 时执行命令 - 创建目录
RUN mkdir -p "/data/app/ui" \
&& ln -sf /dev/stdout /data/app/service.log
# 工作目录
WORKDIR "/data/app"
# 拷贝
COPY main /data/app/main
COPY ui* /data/app/ui/
# docker run 时执行命令
ENTRYPOINT ["./main"]

# https://segmentfault.com/a/1190000042229902

# docker run -d -p 9091:80 -v /mnt/data_sda1/DockerData:/data/app/data novelmap:v1.0.1

# docker build -t novelmap:v1.0.1 .