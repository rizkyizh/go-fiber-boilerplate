# Copilot Instructions

## Stack
Go + Fiber v2 + GORM + PostgreSQL. Swagger docs via `swaggo/swag`. Validation via `go-playground/validator`. Hot reload via `air`.

## Commands

```bash
# Run (development, with hot reload)
air

# Build
go build -o ./tmp/main .

# Lint
golangci-lint run

# Generate/update Swagger docs
swag init

# Install dependencies
go mod tidy
```

No test suite exists yet. When adding tests, use `go test ./...` to run all, or `go test ./app/services/...` to run a single package.

## Architecture

Strict layered clean architecture. The dependency flow is:

```
Route file → Controller → Service → Repository → database.DB (GORM global)
```

All dependency injection is **manual** and wired in each `app/routes/<entity>.route.go` file:

```go
repo := repositories.NewUserRepository()
svc  := services.NewUserService(repo)
ctrl := controllers.NewUserController(svc)
```

Each layer is defined as an **interface + unexported struct** (e.g., `UserService` interface and `userService` struct). Always define the interface in the same file as the implementation.

- **`app/models/`** — GORM models (embed `gorm.Model`). These are the only structs that touch the DB.
- **`app/dto/`** — Request/response payloads with `validate` struct tags.
- **`app/mappers/`** — Pure conversion functions between models and DTOs; no logic.
- **`app/repositories/`** — Direct GORM calls only; no business logic.
- **`app/services/`** — All business logic; no direct DB access.
- **`app/controllers/`** — Fiber handlers; parse input, call service, return response via `utils.ResponseHandler`.
- **`app/routes/`** — Wire up DI and register routes under a `fiber.Router` group.
- **`routes/index.go`** — Top-level router; mounts route groups and registers Swagger + 404.

## Key Conventions

### File naming
All files follow `<entity>.<layer>.go` (e.g., `template.controller.go`, `template.service.go`).

### Mapper naming
Mapper functions use the pattern `SourceType_ToTargetType`:
```go
UserModel_ToUserDTO(user *models.User) *dto.UserDTO
CreateUserDTO_ToUserModel(dto *dto.CreateUserDTO) *models.User
```

### Response handling
All handlers use `utils.ResponseHandler{}`. Never write `c.JSON(...)` directly:
```go
h := &utils.ResponseHandler{}
return h.Ok(c, data, "message", &meta)       // 200
return h.Created(c, nil, "message")           // 201
return h.BadRequest(c, []string{"msg"})       // 400
return h.NotFound(c, []string{"msg"})         // 404
return h.InternalServerError(c, []string{})   // 500
```

### Request validation
Routes that accept a body **must** use the `ValidateRequest` middleware, which rejects unknown fields and runs struct tag validation. The validated body is stored in `c.Locals("validatedReqBody")`.

```go
route.Post("/", middlewares.ValidateRequest(&dto.CreateUserDTO{}), ctrl.CreateUser)
```

### Pagination
Services use `utils.GetPaginationParams(query.Page, query.PerPage)` and return a `utils.Meta` built by `utils.MetaPagination(page, perPage, totalCurrentPage, total)`. Controllers pass `&meta` as the fourth argument to `h.Ok`.

### Swagger annotations
Every controller method needs Swagger godoc comments. Follow the existing `template.controller.go` patterns — `@Router`, `@Param`, `@Success`, `@Failure` with `utils.ResponseData` / `utils.ErrorResponse` as the response objects.

### Database migrations
`database.ConnectDB()` **drops and recreates** the table on every startup (development convenience). When adding a new model, register it in `database.go`'s `AutoMigrate` call.

### Environment
Only one env var is required: `DATABASE_URL` (Postgres DSN). See `.example.env`.

### Commit style
Use conventional commits: `feat:`, `fix:`, `chore:`, etc.
