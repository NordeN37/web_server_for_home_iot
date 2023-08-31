ARG GO_VER=1.17.13
ARG ALPINE_VER=3.16
FROM golang:${GO_VER}-alpine${ALPINE_VER} as builder

# ARG GO_MAIN_PATH
ARG VERSION=0.1
WORKDIR /app
COPY docker .

RUN apk --no-cache update && apk --no-cache add git gcc libc-dev

RUN CGO_ENABLED=1 GOOS=linux go build -tags musl -mod=vendor -a -installsuffix cgo -o app -ldflags "-X 'main.Version=${VERSION}'" ./main.go

FROM alpine:${ALPINE_VER}

WORKDIR /root/
COPY --from=builder /app .

CMD ["./app"]