# Users — Technical Plan

## Architecture

Clean architecture with 4 layers: controller → service → repository → entity.

```
HTTP Request
    ↓
Controller       internal/controller/user_controller.go
    ↓                                auth_controller.go
Service          internal/model/service/user_service_impl.go
    ↓
Repository       internal/model/repository/user_repository_impl.go
                                           auth_repository_impl.go
    ↓
Database         PostgreSQL (raw SQL via database/sql + pgx driver)
```

---

## Files

| Layer | File | Responsibility |
|-------|------|----------------|
| Entity | `internal/model/repository/entity/users.go` | GORM struct, DB schema |
| Entity | `internal/model/repository/entity/refresh_tokens.go` | Refresh token table |
| Domain | `internal/model/domains/user_domain.go` | In-memory user struct |
| Domain | `internal/model/domains/user_domain_interface.go` | Contract for user data |
| Domain | `internal/model/domains/user_domain_password.go` | bcrypt encrypt/compare |
| Converter | `internal/model/converter/convert_domain_to_entity.go` | Domain → Entity |
| Converter | `internal/model/converter/convert_entity_to_domain.go` | Entity → Domain |
| Repository | `internal/model/repository/user_repository_impl.go` | Raw SQL CRUD |
| Repository | `internal/model/repository/auth_repository_impl.go` | Refresh token CRUD |
| Service | `internal/model/service/user_service_impl.go` | Business logic, returns RestErrors |
| Auth | `internal/auth/token.go` | JWT generation/parsing, token hashing |
| Auth | `internal/auth/middleware.go` | Gin JWT middleware |
| Controller | `internal/controller/user_controller.go` | HTTP handlers for users |
| Controller | `internal/controller/auth_controller.go` | HTTP handlers for auth |
| Request DTO | `internal/view/request/user_request.go` | CreateUserRequest, UpdateUserRequest |
| Request DTO | `internal/view/request/auth_request.go` | LoginRequest |
| Response DTO | `internal/view/response/user_response.go` | UserResponse (no password) |
| Routes | `internal/routes/user_routes.go` | User route registration |
| Routes | `internal/routes/auth_routes.go` | Auth route registration |
| Migration | `internal/configuration/database/migrations.go` | GORM AutoMigrate |

---

## Database

### Table: `users_entities`
| Column | Type | Constraints |
|--------|------|-------------|
| id | uuid | PK, default gen_random_uuid() |
| name | varchar(50) | NOT NULL |
| email | varchar(50) | NOT NULL, UNIQUE |
| password | varchar(60) | NOT NULL |
| phone | varchar(20) | NOT NULL |

### Table: `refresh_token_entities`
| Column | Type | Constraints |
|--------|------|-------------|
| id | uuid | PK, default gen_random_uuid() |
| user_id | uuid | NOT NULL, INDEX |
| token_hash | varchar(64) | NOT NULL, UNIQUE INDEX |
| expires_at | timestamp | NOT NULL |
| created_at | timestamp | auto |

---

## Key Technical Decisions

**Two DB connections** — `database/sql` with pgx driver for all runtime queries (performance). GORM used only in migrations (convenience for schema management), then closed.

**Refresh token security** — Raw token (32 random bytes as hex) sent in cookie. Only `SHA-256(token)` stored in DB. No bcrypt needed because the token is random, not a user-chosen password. Lookup is O(1) by hash (unique index).

**Token rotation** — Each `POST /auth/refresh` deletes the old token and issues a new pair. Stolen refresh tokens can only be used once.

**HttpOnly cookies** — Both `access_token` and `refresh_token` set as HttpOnly. JavaScript cannot read them, preventing XSS token theft.

**Password hashing** — bcrypt with `DefaultCost`. Password is hashed in the service layer before reaching the repository. The `UserResponse` DTO never includes the password field.

**Generic auth errors** — Login returns `"invalid credentials"` for both "user not found" and "wrong password". Prevents user enumeration attacks.

**CORS** — Configured to allow only `http://localhost:3000` with `AllowCredentials: true` (required for cookies to work cross-origin).

---

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| JWT_SECRET | — | Secret for signing JWT tokens |
| DB_HOST | localhost | PostgreSQL host |
| DB_PORT | 5432 | PostgreSQL port |
| DB_USER | postgres | PostgreSQL user |
| DB_PASSWORD | postgres | PostgreSQL password |
| DB_NAME | internal | Database name |
| DB_SSLMODE | disable | SSL mode |
