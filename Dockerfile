FROM golang:1.22.5-alpine3.20 AS builder

WORKDIR /app
RUN apk update && apk upgrade && \
    apk add bash git openssh gcc libc-dev
COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ ./
RUN go build -o /dist/server cmd/app/app.go

FROM alpine:3.20

RUN apk add --update ca-certificates tzdata curl pkgconfig && \
    cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
    echo "Asia/Ho_Chi_Minh" > /etc/timezone && \
    rm -rf /var/cache/apk/*

COPY --from=builder /dist/server /app/bin/server

WORKDIR /app/bin
CMD ["/app/bin/server"]