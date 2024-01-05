FROM golang:latest AS build

WORKDIR /app

RUN echo "Set up environment..." \
 && mkdir -p /float32/out

COPY . /float32

RUN echo "Start build..." \
 && cd /float32/backend   \
 && go mod download       \
 && GOOS=linux GOARCH=amd64 go build -v -o /float32/out/main -ldflags "-s -w" ./exec/svr/*.go

RUN echo "Copy prompt..." \
 && cd /float32/prompt    \
 && cp *.promptc /float32/out/

FROM debian:stable-slim

RUN mkdir -p "/app"   \
 && chmod 777 "/app/"

WORKDIR /app

COPY --from=build /float32/out /app

RUN echo "Show up all files..." \
 && ls -al /app

EXPOSE 8080/tcp

ENV LISTEN_ADDR=":8080"

ENTRYPOINT ["/app/main"]
