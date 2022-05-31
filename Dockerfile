FROM golang:1.18.2-alpine3.16 as builder

COPY . /translateBot
WORKDIR /translateBot

RUN go mod download
RUN go build -o bin/server cmd/server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /translateBot/bin/server .
COPY --from=builder /translateBot/config config/

EXPOSE 80

CMD ["./server"]