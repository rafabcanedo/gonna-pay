# Group — Technical Plan

## Architecture

Clean architecture with 4 layers: controller → service → repository → entity.

```
HTTP Request
    ↓
Controller       internal/controller/group_controller.go
    ↓
Service          internal/model/service/group_service_impl.go
    ↓
Repository       internal/model/repository/group_repository_impl.go
    ↓
Database         PostgreSQL (raw SQL via database/sql + pgx driver)
```

---

## Files

| Layer | File | Responsibility |
|-------|------|----------------|
| Entity | `internal/model/repository/entity/group.go` | GORM struct for group_entities table |
| Entity | `internal/model/repository/entity/group_member.go` | GORM struct for group_member_entities table |
| Domain | `internal/model/domains/group_domain.go` | In-memory group struct (with MemberDomain slice) |
| Domain | `internal/model/domains/group_domain_interface.go` | Contract for group data |
| Converter | `internal/model/converter/convert_domain_to_entity.go` | Add GroupDomain → GroupEntity |
| Converter | `internal/model/converter/convert_entity_to_domain.go` | Add GroupEntity → GroupDomain |
| Repository | `internal/model/repository/group_repository_impl.go` | Raw SQL CRUD + member management |
| Service | `internal/model/service/group_service_impl.go` | Business logic, ownership checks, returns RestErrors |
| Controller | `internal/controller/group_controller.go` | HTTP handlers for group + member endpoints |
| Request DTO | `internal/view/request/group_request.go` | CreateGroupRequest, AddMemberRequest |
| Response DTO | `internal/view/response/group_response.go` | GroupResponse (list), GroupDetailResponse (with members) |
| View | `internal/view/convert_domain_to_response.go` | Add GroupDomain → GroupResponse / GroupDetailResponse |
| Routes | `internal/routes/group_routes.go` | Group + member route registration |
| Migration | `internal/configuration/database/migrations.go` | Add GroupEntity, GroupMemberEntity to AutoMigrate |

---

## Database

### Table: `group_entities`
| Column | Type | Constraints |
|--------|------|-------------|
| id | uuid | PK, default gen_random_uuid() |
| owner_id | uuid | NOT NULL, INDEX, FK → users_entities(id) ON DELETE CASCADE |
| name | varchar(100) | NOT NULL |
| category | varchar(20) | NOT NULL |
| created_at | timestamp | auto |
| updated_at | timestamp | auto |

### Table: `group_member_entities`
| Column | Type | Constraints |
|--------|------|-------------|
| id | uuid | PK, default gen_random_uuid() |
| group_id | uuid | NOT NULL, FK → group_entities(id) ON DELETE CASCADE |
| contact_id | uuid | NOT NULL, FK → contact_entities(id) ON DELETE CASCADE |
| — | UNIQUE | (group_id, contact_id) — prevents duplicate members |

---

## Repository Methods

```
// Group CRUD
CreateGroup(domain GroupDomainInterface) (GroupDomainInterface, *rest_errors.RestErr)
FindGroupByID(id, ownerID string) (GroupDomainInterface, *rest_errors.RestErr)
FindAllGroups(ownerID string) ([]GroupDomainInterface, *rest_errors.RestErr)
DeleteGroup(id, ownerID string) *rest_errors.RestErr

// Member management
AddMember(groupID, contactID string) *rest_errors.RestErr
RemoveMember(groupID, contactID, ownerID string) *rest_errors.RestErr
FindMembersByGroupID(groupID string) ([]MemberDomain, *rest_errors.RestErr)
```

`FindGroupByID` joins `group_member_entities` + `contact_entities` to return members inline.

---

## Key Technical Decisions

**Immutability** — No `PUT /group/:id` endpoint. Groups cannot be edited after creation. Scope decision to keep MVP simple.

**Owner isolation** — Every query filters by `owner_id` extracted from the JWT. A user can never read or modify another user's groups.

**403 vs 404** — If a group exists but belongs to another user, return 403 (Forbidden). Avoids leaking whether a group ID exists.

**Members embedded only on detail** — `GET /groups` (list) omits members to avoid N+1 queries. `GET /group/:id` performs a single JOIN to return members inline.

**Member ownership validation** — Before adding a contact as a member, the service checks that the contact's `owner_id` matches the authenticated user. Only the owner's own contacts can be group members.

**Duplicate member prevention** — Enforced at both DB level (UNIQUE constraint on `group_id, contact_id`) and service level (returns 400 before hitting the DB).

**Cascade delete** — `group_member_entities.group_id` has `ON DELETE CASCADE`. Deleting a group automatically removes all its members at the DB level.

**Optional members on creation** — `CreateGroupRequest` accepts an optional `memberIds []string`. The service validates each contact belongs to the owner, then inserts group + members in a transaction.
