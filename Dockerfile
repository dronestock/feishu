FROM storezhang/alpine:3.19.1


LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="Drone持续集成飞书插件，用来做通知使用"


# 复制文件
# 复制执行程序
ARG TARGETPLATFORM
COPY dist/${TARGETPLATFORM}/feishu /usr/local/bin/feishu


RUN set -ex \
    \
    \
    \
    # 增加执行权限
    && chmod +x /usr/local/bin/feishu \
    \
    \
    \
    && rm -rf /var/cache/apk/*


# 执行命令
ENTRYPOINT /usr/local/bin/feishu
