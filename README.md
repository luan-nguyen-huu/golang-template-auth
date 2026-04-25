# 🏗️ Clean Architecture Golang Backend

> A **Golang backend template** following **Clean Architecture principles** — designed for **maintainability, scalability, and testability**.  
> Built-in support for **JWT Authentication** (Register / Login / Refresh Token / Get Me) with token storage in **HTTP-only Cookies**.

---

## 📁 Project Structure

```
golang-template-auth/
├── cmd/
│   ├── main/           # Application entrypoint
│   └── migrate/        # Database migration entrypoint
├── configs/            # Configuration loaders
├── internal/
│   ├── constants/      # App-wide constants
│   ├── entities/       # Domain entities / models
│   ├── exceptions/     # Custom error types
│   ├── handlers/       # HTTP handlers (Controllers)
│   ├── initialize/     # App bootstrapper (DB, Redis, Router...)
│   ├── middlewares/    # HTTP middlewares (auth, logging, etc.)
│   ├── repositories/  # Data access layer (DB queries)
│   ├── routers/        # Route definitions
│   ├── services/       # Business logic layer
│   └── utils/          # Utility helpers
├── .env                # Environment variables
├── docker-compose.yml  # Docker services (PostgreSQL + Redis)
├── Dockerfile          # App Docker image
├── go.mod
├── go.sum
├── MakeFile
└── README.md
```

---

## 🧱 Architecture Overview

This project follows **Clean Architecture** with clear separation of concerns:

```
Handler (HTTP) → Service (Business Logic) → Repository (Database)
                        ↕
                    Entities / Domain
```

| Layer | Responsibility |
|---|---|
| **Handler** | Parse request, call service, return response |
| **Service** | Business rules, orchestration |
| **Repository** | Database queries, data persistence |
| **Entity** | Domain models shared across layers |
| **Middleware** | Auth, CORS, logging, error recovery |

---

## 🔐 Authentication API

Base URL: `{{BASE_URL}}/api/v1`

> All tokens are stored in **HTTP-only Cookies** (`access_token`, `refresh_token`) — not in response body — for enhanced security against XSS attacks.

---

### 📝 Register

**`POST`** `/api/v1/users/register`

Create a new user account.

**Request Body:**
```json
{
  "email": "nguyenhuuluanit69@gmail.com",
  "password": "luanadam108300996",
  "name": "Nguyen Huu Luan"
}
```

**Response:** `201 Created`
```json
{
  "message": "User registered successfully"
}
```

**Cookies set:**
| Cookie | Description |
|---|---|
| `access_token` | Short-lived JWT (default: 15 minutes) |
| `refresh_token` | Long-lived JWT (default: 24 hours) |

---

### 🔑 Login

**`POST`** `/api/v1/users/login`

Authenticate and receive JWT tokens via cookies.

**Request Body:**
```json
{
  "email": "nguyenhuuluanit69@gmail.com",
  "password": "luanadam108300996"
}
```

**Response:** `200 OK`
```json
{
  "message": "Login successful"
}
```

**Cookies set:**
| Cookie | Description |
|---|---|
| `access_token` | Short-lived JWT (default: 15 minutes) |
| `refresh_token` | Long-lived JWT (default: 24 hours) |

---

### 🔄 Refresh Token

**`POST`** `/api/v1/users/refresh`

Use the `refresh_token` cookie to obtain a new `access_token`.

**Headers / Cookies required:** `refresh_token` (auto-sent via browser cookie)

**Response:** `200 OK`
```json
{
  "message": "Token refreshed successfully"
}
```

**Cookies updated:**
| Cookie | Description |
|---|---|
| `access_token` | New short-lived JWT |

---

### 👤 Get Me

**`GET`** `/api/v1/users/me`

Retrieve the currently authenticated user's profile.

**Headers / Cookies required:** `access_token` (auto-sent via browser cookie)

**Response:** `200 OK`
```json
{
  "id": "uuid",
  "email": "nguyenhuuluanit69@gmail.com",
  "name": "Nguyen Huu Luan",
  "created_at": "2024-01-01T00:00:00Z"
}
```

---

## ⚙️ Environment Variables

Create a `.env` file in the project root with the following configuration:

```env
# ───────────────────────────────
# Database (PostgreSQL)
# ───────────────────────────────
DB_USER=admin
DB_PASSWORD=123456
DB_HOST=localhost
DB_PORT=5432
DB_NAME=

# ───────────────────────────────
# Application
# ───────────────────────────────
APP_NAME=github.com
APP_ENV=development
APP_PORT=8080
APP_HOST=0.0.0.0

# ───────────────────────────────
# JWT Configuration
# ───────────────────────────────
JWT_SECRET_ACCESS=
JWT_SECRET_REFRESH=
JWT_ACCESS_TOKEN_EXPIRE=15m
JWT_REFRESH_TOKEN_EXPIRE=24h
```

| Variable | Description | Example |
|---|---|---|
| `DB_USER` | PostgreSQL username | `admin` |
| `DB_PASSWORD` | PostgreSQL password | `123456` |
| `DB_HOST` | PostgreSQL host | `localhost` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_NAME` | PostgreSQL database name | `london` |
| `APP_NAME` | Go module name | `github.com/luan-nguyen-huu/` |
| `APP_ENV` | Environment (`development` / `production`) | `development` |
| `APP_PORT` | HTTP server port | `8080` |
| `APP_HOST` | HTTP server host | `0.0.0.0` |
| `JWT_SECRET_ACCESS` | Secret key for access token signing | `` |
| `JWT_SECRET_REFRESH` | Secret key for refresh token signing | `` |
| `JWT_ACCESS_TOKEN_EXPIRE` | Access token TTL | `15m` |
| `JWT_REFRESH_TOKEN_EXPIRE` | Refresh token TTL | `24h` |

---

## 🚀 Getting Started

### Prerequisites

- [Go 1.21+](https://golang.org/dl/)
- [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)

---

### Step 1 — Start Infrastructure (PostgreSQL + Redis)

Spin up the required services using Docker Compose:

```bash
docker compose up -d
```

This starts:
- **PostgreSQL** — primary database
- **Redis** — caching / session store

Verify services are running:

```bash
docker compose ps
```

---

### Step 2 — Configure Environment

Copy and edit your environment file:

```bash
cp .env.example .env
# Edit .env with your values
```

---

### Step 3 — Run Database Migrations

Before starting the application, apply all database migrations:

```bash
go run cmd/migrate/main.go
```

---

### Step 4 — Start the Application

```bash
go run cmd/main/main.go
```

The server will be available at:

```
http://localhost:8080
```

---

## 🛠️ Makefile Commands

Common shortcuts available via `make`:

```bash
make run        # Run the application
make migrate    # Run database migrations
make build      # Build the binary
make docker-up  # Start Docker services
make docker-down # Stop Docker services
```

---

## 🧪 Tech Stack

| Technology | Purpose |
|---|---|
| **Go** | Primary language |
| **Gin / Fiber** | HTTP framework |
| **PostgreSQL** | Relational database |
| **Redis** | Caching & token store |
| **JWT** | Authentication tokens |
| **Docker Compose** | Local infrastructure |
| **Clean Architecture** | Structural pattern |

---

## 📌 Notes

- **Access Token** expires in `15m` — short-lived for security.
- **Refresh Token** expires in `24h` — used to silently renew access tokens.
- All tokens are stored in **HTTP-only cookies** and are not accessible via JavaScript.
- Run **migrations before** starting the app — the app expects the database schema to exist.

---

## 👤 Author

**Nguyen Huu Luan**  
GitHub: [@luan-nguyen-huu](https://github.com/luan-nguyen-huu)

---

## 📄 License

This project is open-sourced under the [MIT License](LICENSE).
