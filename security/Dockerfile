FROM golang:1.18-alpine as builder

ENV APP_PATH=/gangjin
ENV APP_NAME=security

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.sjtug.sjtu.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk upgrade \
    && apk add git \
    && apk add ca-certificates

WORKDIR $APP_PATH

COPY . .

WORKDIR $APP_PATH/$APP_NAME

RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w GO111MODULE=on \
    && go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o $APP_NAME -a -ldflags '-s -w'

FROM scratch as prod
ENV APP_PATH=/gangjin
ENV APP_NAME=security

# 从 builder 中导入证书
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder $APP_PATH/$APP_NAME .

CMD ["./security"]