FROM golang:1.9.3 as builder

WORKDIR /go/src/github.com/fidelisojeah/go-microservice/user-service

COPY . .

RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep init
RUN echo "[[override]]" >> Gopkg.toml
RUN echo '  branch = "master"' >> Gopkg.toml
RUN echo '  name = "github.com/satori/go.uuid"' >> Gopkg.toml
RUN dep ensure

RUN protoc -I. --go_out=plugins=micro:. proto/auth/auth.proto

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/fidelisojeah/go-microservice/user-service .

CMD ["./user-service"]
