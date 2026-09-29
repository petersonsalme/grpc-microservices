FROM golang:1.26-alpine
WORKDIR /app
# The service imports the gateway's generated protobuf package through a
# replace directive, so both modules are copied from the repository root.
COPY go-grpc-api-gateway/ ./go-grpc-api-gateway/
COPY go-grpc-auth-svc/go.mod go-grpc-auth-svc/go.sum ./go-grpc-auth-svc/
WORKDIR /app/go-grpc-auth-svc
RUN go mod download
COPY go-grpc-auth-svc/ ./
RUN go build -o main cmd/main.go
CMD ["./main"]
