## Description ##
Simple microservice with api-gateway built on node js, services built on go. All the internal service calls through gRPC.

```
project
│   README.md 
│
└───api-gateway (node js)
│   │   .dockerignore
│   │   .gitignore
│   │   docker-compose.yml
│   │   Dockerfile
│   │   package-lock.json
│   │   package.json
│   │   server.js
│   └───controllers
│   │    │____product.controller.js
│   │    │____auth.controller.js
│   │
│   └───proto
│   │    │____auth.proto
│   │    │____product.proto
│   └───routes
│   │    │____auth.routes.js
│   │    │____product.routes.js
│   └───services
│        │____index.js
│        │____service-list.js
│   
└───auth-service (golang)
│   │   Dockerfile
│   │   go.mod
│   │   go.sum
│   │   main.go
│   │   Makefile
│   └───proto
│        │____auth.pb.go
│        │____auth.proto
│   
└───product-service (golang)
│   │   Dockerfile
│   │   go.mod
│   │   go.sum
│   │   main.go
│   │   Makefile
│   └───proto
│        │____product.pb.go
│        │____product.proto
```
