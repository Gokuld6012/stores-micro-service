proto-build:
	protoc -I. --go_out=plugins=grpc:. \
		proto/auth.proto
docker-build:
	docker build -t micro-auth-service .
docker-run:
	docker run -d -p 50051:50051 micro-auth-service