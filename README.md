# Go gRPC Microservices

[![CI](https://github.com/petersonsalme/grpc-microservices/actions/workflows/grpc-ci.yml/badge.svg)](https://github.com/petersonsalme/grpc-microservices/actions/workflows/grpc-ci.yml)

This repository contains a set of microservices built with Go and gRPC, working together to handle authentication, product management, and order processing, fronted by an API Gateway.

## Architecture

The system is composed of four main services and a PostgreSQL database:

1. **API Gateway**: A RESTful HTTP server built using Gin (`go-grpc-api-gateway`). It acts as the single entry point for clients, routing incoming HTTP requests to the appropriate backend gRPC services.
2. **Auth Service**: Manages user registration and authentication (`go-grpc-auth-svc`). It issues JWT tokens to authenticated users.
3. **Product Service**: Handles the creation and retrieval of products (`go-grpc-product-svc`).
4. **Order Service**: Manages customer orders (`go-grpc-order-svc`). It communicates with the Product Service via gRPC to verify product details before creating an order.
5. **PostgreSQL Database**: Serves as the central data store, with individual logical databases for each service (`auth_svc`, `product_svc`, `order_svc`).

### Why gRPC?

gRPC was chosen for inter-service communication for several reasons:
- **Performance**: gRPC uses Protocol Buffers (Protobuf) as its interface definition language and message interchange format. Protobuf is highly efficient and serializes data into a compact binary format, resulting in faster transmission and lower bandwidth usage compared to JSON over HTTP/REST.
- **Strong Typing**: Protobuf enforces a strict contract between services, ensuring type safety and reducing runtime errors.
- **Code Generation**: gRPC automatically generates client and server code in multiple languages, significantly speeding up development and ensuring consistency.
- **HTTP/2**: Built on HTTP/2, gRPC supports multiplexing, bidirectional streaming, and reduced latency, making it ideal for microservices architecture.

## How to Run Locally

### Prerequisites
- Docker and Docker Compose installed.
- (Optional) Go 1.19+ installed if you want to run services outside of Docker.

### Using Docker Compose

The easiest way to run the entire stack is using Docker Compose.

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd grpc-microservices
   ```

2. Start the services:
   ```bash
   docker-compose up --build
   ```

This will start:
- PostgreSQL database on port `5432`
- Auth Service on port `50051`
- Product Service on port `50052`
- Order Service on port `50053`
- API Gateway on port `3000`

The API Gateway will be accessible at `http://localhost:3000`.

### Endpoints (API Gateway)

#### Auth
- `POST /auth/register`: Register a new user.
- `POST /auth/login`: Login and receive a JWT.

#### Product
- `POST /product`: Create a new product.
- `GET /product/:id`: Retrieve a product by ID.

#### Order
- `POST /order`: Create a new order.

*Note: Endpoints requiring authentication must include the JWT token in the `Authorization` header as a Bearer token.*

## Roadmap

This started as a study project. The items below turn it into a study of consistency and resilience between services.

### Consistency
- [ ] Fix the order flow: today the order is saved before the stock is decreased, and it is only rolled back on a `409 Conflict`. Any other `DecreaseStock` failure leaves an order with no stock reserved.
- [ ] Send the ordered quantity to `DecreaseStock` (it currently decreases a fixed amount).
- [ ] Replace the check-then-act stock validation with an atomic, idempotent reservation keyed by order ID.
- [ ] Implement the order flow as a saga (reserve stock → confirm order → compensate on failure), with a transactional outbox so events are not lost when a service restarts.

### Resilience
- [ ] Propagate the incoming request context and set deadlines on every gRPC call instead of `context.Background()`.
- [ ] Add retries with exponential backoff for idempotent calls only.
- [ ] Replace the deprecated `grpc.WithInsecure()` with `grpc.WithTransportCredentials(insecure.NewCredentials())`.
- [ ] Add gRPC health checks and use them in `docker-compose.yml`.
- [ ] Handle database errors that are ignored today (for example, the result of `DB.Create`).

### Observability
- [ ] Distributed tracing with OpenTelemetry across the gateway and the services, viewable in Jaeger.
- [ ] Structured logs with a request ID that follows the call chain.
- [ ] Prometheus metrics for request rate, errors and latency.

### Testing and docs
- [ ] Integration tests for the order flow with Testcontainers (Postgres), covering failure and compensation paths.
- [ ] A "Design decisions" section and ADRs in `docs/adr/` explaining the saga, outbox and idempotency choices, and what was left out on purpose.
