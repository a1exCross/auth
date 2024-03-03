FROM golang:1.20.3-alpine as builder

COPY . /github.com/a1exCross/auth/source/
WORKDIR /github.com/a1exCross/auth/source/

RUN go mod download
RUN go build -o ./bin/auth_server cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/a1exCross/auth/source/bin/auth_server .
COPY --from=builder /github.com/a1exCross/auth/source/config/prod/.env .

CMD ["./auth_server"]