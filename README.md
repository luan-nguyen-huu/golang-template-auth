# 🏗️ Clean Architecture Golang Backend

> A production-grade **Golang backend template** following **Clean Architecture principles** — designed for **maintainability, scalability, and testability**.  
> Built-in support for **JWT Authentication** (Register / Login / Logout / Refresh Token / Get Me) with token storage in **HTTP-only Cookies** and **Bearer Tokens**, full `context.Context` propagation, domain sentinel errors, and input validation.

---

## 📁 Project Structure

```
golang-template-auth/
├── cmd/
│   ├── main/           # Application entrypoint & Graceful Shutdown
│   └── migrate/        # Database migration runner
├── configs/            # Viper configuration loader (environment variables)
├── internal/
│   ├── domain/         # Core domain models, interfaces & sentinel errors
│   ├── entities/       # GORM database schema entities & domain mappers
│   ├── handlers/       # HTTP handlers (Presentation layer) & DTOs
│   ├── initialize/     # Infrastructure bootstrappers (PostgreSQL connection pool)
│   ├── middlewares/    # HTTP middlewares (JWT Auth, CORS, Logger, Recoverer)
│   ├── repositories/   # Data access layer (GORM with context.Context)
│   ├── routers/        # Chi router definitions
│   ├── services/       # Application business logic layer
│   └── utils/          # JWT Maker, bcrypt hasher, validator & response helpers
├── .env.example        # Environment variable template
├── docker-compose.yml  # Docker infrastructure services (PostgreSQL + Redis)
├── Dockerfile          # Multi-stage production container build
├── dev.sh              # Hot-reload runner using Air
├── Makefile            # Build and development commands
└── README.md
```

---

## 🧱 Architecture Overview

This project strictly adheres to **Clean Architecture**:

```
Client (HTTP) → Chi Router → Handler (HTTP Presentation)
                                   ↓
                             Service (Business Rules & Orchestration)
                                   ↓
                             Repository (Data Access & GORM)
                                   ↓
                             PostgreSQL Database

* Core Domain (`internal/domain`) defines pure models, interfaces, and sentinel errors without external dependencies.
```

| Layer | Responsibility |
|---|---|
| **Domain** | Pure business models, interface contracts (`UserRepository`, `UserService`), and domain errors. |
| **Handler** | Parse request, validate input, invoke service, set cookies, and map domain errors to HTTP status codes. |
| **Service** | Core business logic, password hashing, JWT token issuance. |
| **Repository** | Database persistence with `context.Context` query lifecycle. |
| **Entities** | Database schema definitions & mappers (`ToDomain`, `FromDomainUser`). |
| **Middlewares** | JWT verification (Cookie / Bearer Header), CORS with credentials, logging, and panic recovery. |

---

## 🔐 Authentication API

Base URL: `{{BASE_URL}}/api/v1`

---

### 📝 Register

**`POST`** `/api/v1/users/register`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "strongPassword123",
  "name": "John Doe"
}
```

**Response:** `201 Created`
```json
{
  "response": {
    "code": 201,
    "message": "User registered successfully",
    "data": {
      "access_token": "...",
      "refresh_token": "..."
    }
  }
}
```

---

### 🔑 Login

**`POST`** `/api/v1/users/login`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "strongPassword123"
}
```

**Response:** `200 OK`
```json
{
  "response": {
    "code": 200,
    "message": "User logged in successfully",
    "data": {
      "access_token": "...",
      "refresh_token": "..."
    }
  }
}
```

---

### 👤 Get Me

**`GET`** `/api/v1/users/me`

**Authentication:** `access_token` cookie or `Authorization: Bearer <token>` header.

**Response:** `200 OK`
```json
{
  "response": {
    "code": 200,
    "message": "User profile retrieved successfully",
    "data": {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "email": "user@example.com",
      "name": "John Doe",
      "created_at": "2026-08-30T12:00:00Z",
      "updated_at": "2026-08-30T12:00:00Z"
    }
  }
}
```

---

### 🔄 Refresh Token

**`POST`** `/api/v1/users/refresh`

**Authentication:** `refresh_token` cookie or `Authorization: Bearer <refresh_token>` header.

**Response:** `200 OK`

---

### 🚪 Logout

**`POST`** `/api/v1/users/logout`

Clears auth cookies.

---

## 🛠️ Makefile Commands

```bash
make run         # Run the application directly
make dev         # Run with Air hot reload
make migrate     # Run database migrations
make build       # Compile binaries to bin/
make test        # Run unit tests with race detection
make tidy        # Download and clean Go module dependencies
make docker-up   # Start PostgreSQL and Redis via Docker Compose
make docker-down # Stop Docker Compose containers
```

---

## 🚀 Getting Started

1. **Start Infrastructure:**
   ```bash
   make docker-up
   ```

2. **Configure Environment:**
   ```bash
   cp .env.example .env
   ```

3. **Run Migrations:**
   ```bash
   make migrate
   ```

4. **Start Development Server:**
   ```bash
   make dev
   ```
