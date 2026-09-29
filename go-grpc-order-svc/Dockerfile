FROM golang:1.26-alpine
WORKDIR /app
# The service imports the gateway's generated protobuf package through a
# replace directive, so both modules are copied from the repository root.
COPY go-grpc-api-gateway/ ./go-grpc-api-gateway/
COPY go-grpc-order-svc/go.mod go-grpc-order-svc/go.sum ./go-grpc-order-svc/
WORKDIR /app/go-grpc-order-svc
RUN go mod download
COPY go-grpc-order-svc/ ./
RUN go build -o main cmd/main.go
CMD ["./main"]
