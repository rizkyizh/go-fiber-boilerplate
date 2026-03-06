# Go Fiber Boilerplate

Go Fiber Boilerplate is a production-ready starter template for building REST APIs using Go Fiber, GORM, and PostgreSQL. It includes authentication, authorization, security hardening, structured logging, and more — following clean architecture principles.

## Technologies Used

- **Golang** — programming language
- **Fiber v2** — Express-inspired web framework
- **GORM** — ORM library for Go
- **PostgreSQL** — relational database
- **JWT** — stateless authentication (`golang-jwt/jwt/v5`)
- **Zerolog** — structured JSON logging (`rs/zerolog`)
- **Swagger** — auto-generated API docs (`swaggo/swag`)
- **Air** — hot reload for development

## Features

| Feature | Description |
|---|---|
| **JWT Auth** | Register, login, refresh token, logout |
| **RBAC** | Role-based access control (`admin`, `user`) |
| **Rate Limiting** | 60 req/min per IP (global) |
| **Helmet** | Security HTTP headers |
| **Request ID** | `X-Request-ID` on every request |
| **Structured Logging** | JSON logs via zerolog |
| **Health Check** | `GET /health` — checks DB connectivity |
| **CRUD + Pagination** | Full user resource with pagination |
| **Validation** | Input validation with `go-playground/validator` |
| **Seeder** | Auto-seeds admin and user fixtures on startup |
| **Unit Tests** | Service-layer tests with manual mocks |
| **Makefile** | Developer shortcuts |
| **Swagger Docs** | Full API documentation with security support |

## Getting Started

### Prerequisites

- Go 1.23+
- PostgreSQL

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/rizkyizh/go-fiber-boilerplate.git
   cd go-fiber-boilerplate
   ```

2. **Install dependencies:**

   ```bash
   make tidy
   ```

3. **Configure environment:**

   ```bash
   cp .example.env .env
   # Edit .env with your values
   ```

4. **Run with hot reload:**

   ```bash
   make run
   ```

### Environment Variables

| Variable | Required | Default | Description |
|---|---|---|---|
| `DATABASE_URL` | ✅ | — | PostgreSQL DSN |
| `JWT_SECRET` | ⚠️ | `change-me-in-production-access` | Access token signing secret |
| `JWT_REFRESH_SECRET` | ⚠️ | `change-me-in-production-refresh` | Refresh token signing secret |
| `JWT_ACCESS_EXPIRY` | ❌ | `15m` | Access token TTL (Go duration string) |
| `JWT_REFRESH_EXPIRY` | ❌ | `168h` | Refresh token TTL |
| `APP_ENV` | ❌ | — | Set to `development` for pretty logs |

> ⚠️ Always override `JWT_SECRET` and `JWT_REFRESH_SECRET` in production.

### Seeded Users

On every startup the database is seeded with:

| Email | Password | Role |
|---|---|---|
| `admin@example.com` | `Admin1234!` | `admin` |
| `user@example.com` | `User1234!` | `user` |

## API Endpoints

### Auth (`/auth`)

| Method | Path | Auth | Description |
|---|---|---|---|
| POST | `/auth/register` | Public | Register a new account |
| POST | `/auth/login` | Public | Login, receive token pair |
| POST | `/auth/refresh` | Public | Refresh access token |
| POST | `/auth/logout` | 🔒 Bearer | Invalidate refresh token |

### Users (`/users`)

| Method | Path | Auth | Role | Description |
|---|---|---|---|---|
| GET | `/users` | 🔒 Bearer | any | List users (paginated) |
| GET | `/users/:id` | 🔒 Bearer | any | Get single user |
| POST | `/users` | 🔒 Bearer | admin | Create user |
| PATCH | `/users/:id` | 🔒 Bearer | admin | Update user |
| DELETE | `/users/:id` | 🔒 Bearer | admin | Delete user |

### Other

| Method | Path | Description |
|---|---|---|
| GET | `/health` | App and DB health status |
| GET | `/swagger/*` | Interactive API docs |

## Makefile Targets

```bash
make run            # Start with hot reload (air)
make build          # Compile binary to ./tmp/main
make test           # Run all unit tests
make test/coverage  # Run tests with HTML coverage report
make swagger        # Regenerate Swagger docs (requires swag)
make lint           # Run golangci-lint
make tidy           # Tidy go modules
```

## Middleware Chain

```
RequestID → Logger → Helmet → CORS → RateLimiter → Routes
```

Protected routes additionally pass through:

```
AuthRequired → [RequireRole] → Handler
```

## Project Structure

```
.
├── app
│   ├── controllers    # HTTP handlers (auth, users, health)
│   ├── dto            # Request/response DTOs with validation tags
│   ├── mappers        # Model ↔ DTO conversion functions
│   ├── models         # GORM models (User with Role, Password)
│   ├── repositories   # GORM data access layer
│   ├── routes         # DI wiring + route registration
│   └── services       # Business logic (auth, users)
├── config             # Environment config loader
├── database           # DB connection, migrations, seeder
├── docs               # Auto-generated Swagger files
├── middlewares        # auth, rbac, cors, helmet, logger, requestid, limiter
├── routes             # Top-level router
├── utils              # Response handler, JWT, pagination, validator
├── Makefile
└── main.go
```

## Contributing

1. Fork the project.
2. Create your branch `git checkout -b feat/your-feature`.
3. Commit `git commit -m 'feat: add your feature'`.
4. Push and open a pull request.

## License

This project is licensed under the MIT License.

