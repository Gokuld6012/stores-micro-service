# choose alpine go too
FROM golang:alpine as builder

WORKDIR $GOPATH/go/src/github.com/gokuld6012/product-micro/product-service

COPY . .

RUN go get -d -v ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/product-service

######## Start product service #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/bin/product-service .

EXPOSE 50052

CMD ["./product-service"]
