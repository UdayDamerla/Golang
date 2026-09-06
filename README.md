# Golang Learning Repository

This is my practical Golang learning workspace.
I use it to combine concept tutorials with a basic API project implementation.

## Repository Goal

- Learn Go fundamentals through small tutorials.
- Practice concurrency, channels, and generics with runnable examples.
- Build a basic API project with authentication and middleware.
- Follow clean project structure and module-based package management.

## Folder Structure

- `tutorials/`: Hands-on topic-wise programs (small, focused examples).
- `api/`: A basic production-style API project with clean architecture.

## Suggested Learning Path

1. I start with basics in `tutorials/`:
   - Variables, functions, structs, interfaces
   - Goroutines and synchronization
   - Channels (buffered/unbuffered, close, range, select)
   - Generics
2. I move to `api/` after fundamentals.
3. I understand how packages are split by responsibility.
4. I run and test endpoints with curl or Postman.

## Basic Project: Go API (inside `api/`)

In this project, I practice:

- Standard Go project layout (`cmd`, `internal`, `pkg`)
- Routing using Chi
- Authentication using JWT
- Password hashing using bcrypt
- Middleware (logging and auth)
- Environment-based configuration

### API Structure

- `cmd/server/main.go`: App entrypoint
- `internal/config`: Environment configuration
- `internal/router`: Route registration
- `internal/handlers`: HTTP request handlers
- `internal/services`: Business logic
- `internal/middleware`: Logging and JWT guard
- `pkg/response`: Shared JSON response helper

## How To Run

From the `api` directory, I run:

```bash
go mod tidy
go run ./cmd/server
```

Server default:

- `http://localhost:8080`

## Basic Endpoint Test Flow

1. Health check

```bash
curl http://localhost:8080/health
```

2. Register user

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"pass123"}'
```

3. Login user

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"pass123"}'
```

4. Access protected profile endpoint with token

```bash
curl http://localhost:8080/api/v1/profile \
  -H "Authorization: Bearer <TOKEN_FROM_LOGIN_RESPONSE>"
```

## Next Improvements

- Add refresh token flow
- Add role-based authorization
- Add database integration (PostgreSQL)
- Add unit and integration tests
- Add Dockerfile and CI pipeline
