# choose alpine go too
FROM golang:alpine as builder

WORKDIR /go/src/github.com/gokuld6012/product-micro/auth-service

COPY . .

RUN go get -d -v ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/auth-service

######## Start auth service #######

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/


COPY --from=builder /go/bin/auth-service .

EXPOSE 50051

CMD ["./auth-service"]
