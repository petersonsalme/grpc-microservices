FROM golang:1.26-alpine
WORKDIR /app
# The service imports the gateway's generated protobuf package through a
# replace directive, so both modules are copied from the repository root.
COPY go-grpc-api-gateway/ ./go-grpc-api-gateway/
COPY go-grpc-product-svc/go.mod go-grpc-product-svc/go.sum ./go-grpc-product-svc/
WORKDIR /app/go-grpc-product-svc
RUN go mod download
COPY go-grpc-product-svc/ ./
RUN go build -o main cmd/main.go
CMD ["./main"]
