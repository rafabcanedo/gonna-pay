# Users — Specification

## Overview

The Users feature manages user accounts in the system. It covers registration, profile management, and authentication.

---

## Endpoints

### Public (no authentication required)

| Method | Path | Description |
|--------|------|-------------|
| POST | `/users` | Register a new user |
| POST | `/auth/login` | Authenticate and receive tokens |
| POST | `/auth/refresh` | Refresh access token |
| POST | `/auth/logout` | Invalidate session |

### Protected (JWT cookie required)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/users` | List all users |
| GET | `/users/:id` | Get a user by ID |
| PUT | `/users/:id` | Update a user |
| DELETE | `/users/:id` | Delete a user |

---

## Business Rules

### Registration (`POST /users`)
- `name`, `email`, `password`, `phone` are required
- `email` must be a valid email format
- `email` must be unique — returns 400 if already in use
- `password` minimum 6 characters
- Password is stored hashed (bcrypt), never in plaintext

### Login (`POST /auth/login`)
- Requires `email` and `password`
- Returns 401 with generic `"invalid credentials"` for both wrong email and wrong password (security: does not reveal which field is wrong)
- On success: sets two HttpOnly cookies (`access_token`, `refresh_token`) and returns user info

### Token Refresh (`POST /auth/refresh`)
- Reads `refresh_token` cookie
- Invalidates the used token and issues a new pair (rotation)
- Returns 401 if token is missing, invalid, or expired

### Logout (`POST /auth/logout`)
- Deletes refresh token from database
- Clears both cookies
- Always returns 200 (even if no cookie was present)

### Update (`PUT /users/:id`)
- All fields are optional
- If `password` is provided, it is re-hashed before saving
- `email` must be valid format if provided

### Delete (`DELETE /users/:id`)
- Returns 200 with `{ "message": "user deleted successfully" }`

---

## Responses

### User object
```json
{
  "id": "uuid",
  "name": "string",
  "email": "string",
  "phone": "string"
}
```

### Error object
```json
{
  "message": "string",
  "error": "string",
  "code": 400,
  "causes": [{ "field": "string", "message": "string" }]
}
```
