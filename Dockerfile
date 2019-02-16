# BUILD
FROM golang:1.11-alpine as builder

RUN apk add --no-cache git mercurial 

ENV p $GOPATH/src/github.com/werberson/prometheus-metrics-sample

ADD ./ ${p}
WORKDIR ${p}
RUN go get -v ./...

RUN GIT_COMMIT=$(git rev-parse --short HEAD 2> /dev/null || true) \
 && BUILDTIME=$(TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ') \
 && VERSION=$(git describe --abbrev=0 --tags 2> /dev/null || true) \
 && CGO_ENABLED=0 GOOS=linux go build --ldflags "-s -w \
    -X github.com/werberson/prometheus-metrics-sample/version.Version=${VERSION:-unknow-version} \
    -X github.com/werberson/prometheus-metrics-sample/version.GitCommit=${GIT_COMMIT} \
    -X github.com/werberson/prometheus-metrics-sample/version.BuildTime=${BUILDTIME}" \
    -a -installsuffix cgo -o metrics-sample main.go

ENTRYPOINT [ "/go/src/github.com/werberson/prometheus-metrics-sample/metrics-sample" ]

CMD [ "serve" ]

## PKG
#FROM scratch
#
#COPY --from=builder /metrics-sample /go/bin/
#COPY --from=builder /go/src/github.com/werberson/prometheus-metrics-sample/web/ui/static /
#
#
#ENTRYPOINT [ "/go/bin/metrics-sample" ]

CMD [ "serve" ]
