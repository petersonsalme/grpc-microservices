# AGENTS.md — grpc-microservices

Diretrizes para agentes de IA atuando no repositório `grpc-microservices`.

## Contexto do Projeto

Arquitetura de microsserviços em Go utilizando gRPC e Protocol Buffers para comunicação entre serviços, fronted por uma API Gateway REST (Gin), com persistência em PostgreSQL.

- Para o diagrama arquitetural completo, justificativa técnica de uso de gRPC e fluxo de requisições, consulte o [README.md](README.md).

## Estrutura do Repositório

```
grpc-microservices/
├── go-grpc-api-gateway/   # Gateway HTTP (Gin) -> clientes gRPC
├── go-grpc-auth-svc/      # Serviço gRPC de autenticação (JWT) e usuários
├── go-grpc-product-svc/   # Serviço gRPC de catálogo e estoque de produtos
├── go-grpc-order-svc/     # Serviço gRPC de criação e gestão de pedidos
├── docker-compose.yml     # Orquestração com banco PostgreSQL (bancos: auth_svc, product_svc, order_svc)
├── init-db.sh             # Script de inicialização dos esquemas de banco
└── README.md              # Documentação detalhada dos serviços e portas
```

## Regras e Comandos

- Prefix shell commands with `rtk` (see `.agents/rules/rtk.md`).

1. **Estrutura por Serviço:**
   - Cada serviço possui seu próprio `go.mod`, `Dockerfile` e `Makefile` em seu respectivo diretório.
   - Ao executar comandos Go, navegue até a pasta do serviço específico (ex.: `cd go-grpc-api-gateway && go test ./...`).
2. **Contratos Protobuf (`.proto`):**
   - Alterações em contratos de mensageria devem manter compatibilidade com os clientes e servidores nos outros serviços.
   - Atualize os arquivos `.proto` e regenere o código correspondente caso contratos sejam alterados.
3. **Ambiente Local:**
   - Subida completa do ecossistema: `docker-compose up --build`
4. **Padrões de Desenvolvimento:**
   - Conventional Commits (`feat:`, `fix:`, `refactor:`, `test:`).
   - Respostas ao usuário em Português do Brasil; código e documentação técnica em Inglês.

@.agents/rules/rtk.md
