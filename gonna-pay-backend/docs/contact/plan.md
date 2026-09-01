# Contact — Technical Plan

## Architecture

Clean architecture with 4 layers: controller → service → repository → entity.

```
HTTP Request
    ↓
Controller       internal/controller/contact_controller.go
    ↓
Service          internal/model/service/contact_service_impl.go
    ↓
Repository       internal/model/repository/contact_repository_impl.go
    ↓
Database         PostgreSQL (raw SQL via database/sql + pgx driver)
```

---

## Files

| Layer | File | Responsibility |
|-------|------|----------------|
| Entity | `internal/model/repository/entity/contact.go` | GORM struct, DB schema (already done) |
| Domain | `internal/model/domains/contact_domain.go` | In-memory contact struct |
| Domain | `internal/model/domains/contact_domain_interface.go` | Contract for contact data |
| Converter | `internal/model/converter/convert_domain_to_entity.go` | Add ContactDomain → ContactEntity |
| Converter | `internal/model/converter/convert_entity_to_domain.go` | Add ContactEntity → ContactDomain |
| Repository | `internal/model/repository/contact_repository_impl.go` | Raw SQL CRUD |
| Service | `internal/model/service/contact_service_impl.go` | Business logic, returns RestErrors |
| Controller | `internal/controller/contact_controller.go` | HTTP handlers for contacts |
| Request DTO | `internal/view/request/contact_request.go` | CreateContactRequest, UpdateContactRequest |
| Response DTO | `internal/view/response/contact_response.go` | ContactResponse |
| Routes | `internal/routes/contact_routes.go` | Contact route registration |
| Migration | `internal/configuration/database/migrations.go` | Add ContactEntity to AutoMigrate |

---

## Database

### Table: `contact_entities`
| Column | Type | Constraints |
|--------|------|-------------|
| id | uuid | PK, default gen_random_uuid() |
| owner_id | uuid | NOT NULL, INDEX, FK → users_entities(id) ON DELETE CASCADE |
| name | varchar(50) | NOT NULL |
| email | varchar(50) | NOT NULL |
| phone | varchar(20) | NOT NULL |
| category | varchar(20) | NOT NULL |

---

## Key Technical Decisions

**Owner isolation** — Every query filters by `owner_id` extracted from the JWT. A user can never read or modify another user's contacts, enforced at the repository layer.

**403 vs 404** — If a contact exists but belongs to another user, return 403 (Forbidden), not 404. This avoids leaking whether a contact ID exists in the system.

**No unique constraint on email** — A contact's email is personal data, not a login credential. Multiple users may have the same person as a contact, so no unique index on `email`.

**Ownership extracted from JWT** — `ownerId` is never taken from the request body. It is always read from the authenticated JWT claims in the controller and passed down to the service.
