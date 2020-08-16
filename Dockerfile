FROM golang:1.14-alpine AS build-dist

ENV GOPROXY='https://goproxy.cn,direct'
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download
WORKDIR /go/release
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -tags netgo -installsuffix cgo -o /bin/dashboard main.go


FROM alpine as prod
COPY --from=build-dist /bin/dashboard /bin/dashboard
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache -U  tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    chmod +x /bin/dashboard

ENV PG_HOST='postgres'
ENV PG_PORT=5432
ENV PG_USER='dev'
ENV PG_PASSWORD='dev'
ENV PG_NAME='dev'
ENV GIN_MODE='release'

CMD ["/bin/dashboard"]
EXPOSE 8080