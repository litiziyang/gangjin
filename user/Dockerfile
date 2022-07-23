FROM golang:1.18 as builder

ENV APP_PATH=/gangjin
ENV APP_NAME=user

WORKDIR $APP_PATH

COPY . .

WORKDIR $APP_PATH/$APP_NAME

RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w GO111MODULE=on \
    && go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o $APP_NAME -a -ldflags '-s -w'

FROM scratch as prod
ENV APP_PATH=/gangjin
ENV APP_NAME=user

COPY --from=builder $APP_PATH/$APP_NAME .

CMD ["./user"]