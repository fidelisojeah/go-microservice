FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/fidelisojeah/go-microservice/consignment-cli

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep init && dep ensure

RUN go get -u github.com/satori/go.uuid

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir -p /app
WORKDIR /app

COPY consignment.json /app/consignment.json
COPY --from=builder /go/src/github.com/fidelisojeah/go-microservice/consignment-cli .

CMD ["./consignment-cli"]
