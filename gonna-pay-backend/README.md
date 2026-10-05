<div align="center">
  <img src="../gonna-pay-frontend/src/assets/gonna-logo.svg" alt="Gonna Pay" width="200" />
</div>

<br />

<div align="center">
  <p>REST API — Gonna Pay Backend</p>
</div>

---

## About

The core REST API of Gonna Pay. Handles authentication, contacts, groups, costs and automatic expense splitting.

## Tech Stack

- **Go** — language
- **Gin** — HTTP framework
- **PostgreSQL** — database (raw SQL via `pgx`)
- **GORM** — migrations only
- **JWT** — authentication via HTTP-only cookie
- **Resend** — transactional emails (verification, password reset)
- **Swagger** — API documentation (swaggo)
- **Zap** — structured logging
- **Docker** — containerization

## Getting Started

### Prerequisites

- [Go 1.21+](https://go.dev/dl/)
- [Docker](https://www.docker.com/)

### Environment Variables

```bash
cp .env.example .env
```

Fill in the required values in `.env`.

### Running Locally

**Full stack with Docker Compose** (app + database):

```bash
docker-compose up
```

**Database only** (run the API with Go):

```bash
docker-compose up db
go run ./cmd/api
```

The API will be available at `http://localhost:3333`.

## API Documentation

Swagger UI is available at:

```
http://localhost:3333/swagger/index.html
```

## Running Tests

```bash
go test ./...
```
