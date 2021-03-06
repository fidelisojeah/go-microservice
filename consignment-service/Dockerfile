FROM golang:1.9.3 as builder

WORKDIR /go/src/github.com/fidelisojeah/go-microservice/consignment-service

COPY . .

RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep init
RUN echo "[[override]]" >> Gopkg.toml
RUN echo '  branch = "master"' >> Gopkg.toml
RUN echo '  name = "github.com/satori/go.uuid"' >> Gopkg.toml
RUN dep ensure

RUN protoc -I. --go_out=plugins=micro:. proto/consignment/consignment.proto

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


# Here we're using a second FROM statement, which is strange,
# but this tells Docker to start a new build process with this
# image.
FROM alpine:latest

# Security related package, good to have.
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

# Here, instead of copying the binary from our host machine,
# we pull the binary from the container named `builder`, within
# this build context. This reaches into our previous image, finds
# the binary we built, and pulls it into this container. Amazing!
COPY --from=builder /go/src/github.com/fidelisojeah/go-microservice/consignment-service .

CMD [ "./consignment-service" ]
