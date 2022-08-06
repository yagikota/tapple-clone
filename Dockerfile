# ===== build go binary =====
FROM golang:1.18.3-alpine as go-builder

WORKDIR /go/src/github.com/CyberAgentHack/2208-ace-a-server/pkg

COPY pkg/ .

RUN go mod download

RUN go build -o server main.go

# ==== build docker image ====
FROM alpine

RUN apk --no-cache add tzdata

COPY --from=go-builder /go/src/github.com/CyberAgentHack/2208-ace-a-server/pkg/server server

ENTRYPOINT ["/server"]
