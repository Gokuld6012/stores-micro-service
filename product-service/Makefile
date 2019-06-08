proto-build:
	protoc -I. --go_out=plugins=grpc:. \
		proto/product.proto
docker-build:
	docker build --rm=true -t micro-product-service .
docker-run:
	docker run -d -p 50052:50052 micro-product-service
