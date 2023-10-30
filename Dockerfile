ARG GO_VERSION

# Build for amd64
FROM golang:$GO_VERSION-alpine as builder-amd64
WORKDIR /go/src

COPY ./app .
RUN GOARCH=amd64 go build -asmflags="-trimpath=OPATH" -ldflags="-w -s" -gcflags="-trimpath=OPATH" -o jb-app ./cmd

FROM alpine:latest as amd64
COPY --from=builder-amd64 /go/src/jb-app /jb-app
RUN mkdir -p /var/log/service/ && touch /var/log/service/service.log && chmod 0777 /var/log/service/service.log
ENTRYPOINT ["/jb-app"]

# Build for arm64
FROM golang:$GO_VERSION-alpine as builder-arm64
WORKDIR /go/src

COPY ./app .
RUN  GOARCH=arm64 go build -asmflags="-trimpath=OPATH" -ldflags="-w -s" -gcflags="-trimpath=OPATH" -o jb-app ./cmd

FROM alpine:latest as arm64
COPY --from=builder-arm64 /go/src/jb-app /jb-app
RUN mkdir -p /var/log/service/ && touch /var/log/service/service.log && chmod 0777 /var/log/service/service.log
ENTRYPOINT ["/jb-app"]
