FROM storezhang/alpine:3.19.1


LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="Drone持续集成飞书插件，用来做通知使用"


# 复制文件
COPY feishu /bin


RUN set -ex \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/feishu \
    \
    \
    \
    && rm -rf /var/cache/apk/*


# 执行命令
ENTRYPOINT /bin/feishu
