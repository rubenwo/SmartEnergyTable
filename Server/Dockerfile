FROM golang:1.14 AS builder

WORKDIR /go/server
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/server/server .

CMD [ "./server" ]