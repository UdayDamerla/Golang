# Go API Project

A clean Go API project structure with authentication, middleware, and package-managed dependencies.

## Structure

- cmd/server: entrypoint
- internal/config: config loading
- internal/router: route registration
- internal/handlers: HTTP handlers
- internal/services: business logic (auth)
- internal/middleware: logging and JWT auth middleware
- pkg/response: shared response utility

## Features

- REST API with chi router
- Register and login endpoints
- JWT token generation and validation
- Protected endpoint with auth middleware
- Basic request logging middleware

## Package Management

Dependencies are tracked with Go modules in go.mod.
Use go mod tidy to resolve and pin transitive dependencies.

## Run

1. Copy .env.example to .env (optional)
2. Export env vars if needed:
   - export PORT=8080
   - export JWT_SECRET=super-secret
3. Start server:
   - go run ./cmd/server

## API Endpoints

- GET /health
- POST /api/v1/auth/register
- POST /api/v1/auth/login
- GET /api/v1/profile (Bearer token required)

## Sample Requests

Register:

curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"pass123"}'

Login:

curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"pass123"}'

Profile (replace TOKEN):

curl http://localhost:8080/api/v1/profile \
  -H "Authorization: Bearer TOKEN"
