# Cost — Technical Plan

## Architecture

Clean architecture with 4 layers: controller → service → repository → entity.

```
HTTP Request
    ↓
Controller       internal/controller/cost_controller.go
    ↓
Service          internal/model/service/cost_service_impl.go
    ↓
Repository       internal/model/repository/cost_repository_impl.go
    ↓
Database         PostgreSQL (raw SQL via database/sql + pgx driver)
```

---

## Files

| Layer | File | Responsibility |
|-------|------|----------------|
| Entity | `internal/model/repository/entity/cost.go` | GORM struct for cost_entities table |
| Entity | `internal/model/repository/entity/cost_split.go` | GORM struct for cost_split_entities table |
| Domain | `internal/model/domains/cost_domain.go` | In-memory cost struct (with SplitDomain slice) |
| Domain | `internal/model/domains/cost_domain_interface.go` | Contract for cost data |
| Converter | `internal/model/converter/convert_domain_to_entity.go` | Add CostDomain → CostEntity |
| Converter | `internal/model/converter/convert_entity_to_domain.go` | Add CostEntity → CostDomain |
| Repository | `internal/model/repository/cost_repository_impl.go` | Raw SQL CRUD + split creation/recalculation in transaction |
| Service | `internal/model/service/cost_service_impl.go` | Business logic, ownership checks, split calculation/recalculation, returns RestErrors |
| Controller | `internal/controller/cost_controller.go` | HTTP handlers for cost endpoints |
| Request DTO | `internal/view/request/cost_request.go` | CreateCostRequest, UpdateCostRequest |
| Response DTO | `internal/view/response/cost_response.go` | CostResponse (list), CostDetailResponse (with splits) |
| View | `internal/view/convert_domain_to_response.go` | Add CostDomain → CostResponse / CostDetailResponse |
| Validation | `internal/configuration/validation/validate_cost.go` | ValidateCostError |
| Routes | `internal/routes/cost_routes.go` | Cost route registration |
| Migration | `internal/configuration/database/migrations.go` | Add CostEntity, CostSplitEntity to AutoMigrate |
| Main | `cmd/api/main.go` | Wire costRepo, costSvc, costCtrl |

---

## Database

### Table: `cost_entities`
| Column | Type | Constraints |
|--------|------|-------------|
| id | uuid | PK, default gen_random_uuid() |
| user_id | uuid | NOT NULL, INDEX, FK → users_entities(id) ON DELETE CASCADE |
| group_id | uuid | NULLABLE, FK → group_entities(id) ON DELETE SET NULL |
| cost_name | varchar(100) | NOT NULL |
| total_value | numeric(10,2) | NOT NULL |
| owner_percentage | numeric(5,2) | NOT NULL |
| category | varchar(20) | NOT NULL |
| created_at | timestamp | auto |
| updated_at | timestamp | auto |

### Table: `cost_split_entities`
| Column | Type | Constraints |
|--------|------|-------------|
| id | uuid | PK, default gen_random_uuid() |
| cost_id | uuid | NOT NULL, FK → cost_entities(id) ON DELETE CASCADE |
| contact_id | uuid | NOT NULL, FK → contact_entities(id) ON DELETE CASCADE |
| value | numeric(10,2) | NOT NULL |
| percentage | numeric(5,2) | NOT NULL |
| created_at | timestamp | auto |
| updated_at | timestamp | auto |

---

## Repository Methods

```
Create(domain CostDomainInterface, memberIDs []string) (CostDomainInterface, error)
Update(id string, domain CostDomainInterface) (CostDomainInterface, error)
FindAll(ownerID string) ([]CostDomainInterface, error)
FindByID(id string) (CostDomainInterface, error)
Delete(id string) error
GetGroupMemberIDs(groupID string) ([]string, error)
GetGroupName(groupID string) (string, error)
```

- `Create` uses a DB transaction: inserts cost_entity first, then cost_split_entities for each member
- `Update` uses a DB transaction: updates cost_entities row, then recalculates and updates all cost_split_entities rows if any exist; calls `FindByID` to return the updated detail
- `FindAll` uses a LEFT JOIN on group_entities to return groupName, and a COUNT subquery for splitCount
- `FindByID` uses LEFT JOINs on cost_split_entities + contact_entities to return splits inline

---

## Split Calculation Logic (Service Layer)

```
if groupId is provided:
    memberIDs = repo.GetGroupMemberIDs(groupId)
    memberCount = len(memberIDs)

    if ownerPercentage not provided:
        ownerPercentage = 100 / (memberCount + 1)
    
    memberPercentage = (100 - ownerPercentage) / memberCount
    memberValue = totalValue * memberPercentage / 100

else:
    ownerPercentage = 100
    memberIDs = []
```

---

## Update Recalculation Logic (Repository Layer)

```
if splitCount > 0:
    memberPercentage = round((100 - ownerPercentage) / splitCount, 2)
    memberValue = round(totalValue * memberPercentage / 100, 2)
    UPDATE cost_split_entities SET value = memberValue, percentage = memberPercentage WHERE cost_id = id
```

---

## Key Technical Decisions

**Mutable scalar fields only** — `PUT /cost/:id` allows updating `costName`, `totalValue`, `category`, and `ownerPercentage`. The `groupId` cannot be changed — changing the group would require deleting and recreating splits which is better handled by delete + recreate.

**Owner isolation** — Every query filters by `user_id` extracted from the JWT. A user can never read or modify another user's costs.

**403 vs 404** — If a cost exists but belongs to another user, return 403 (Forbidden). Avoids leaking whether a cost ID exists.

**Atomic creation** — `POST /cost` creates cost + all splits in a single DB transaction. If any insert fails, the entire transaction rolls back.

**Group FK ON DELETE SET NULL** — If a group is deleted, `group_id` on cost_entities is set to NULL. The cost is preserved but no longer linked to the group.

**Split FK ON DELETE CASCADE** — Deleting a cost automatically removes all its splits at the DB level.

**ownerValue computed at query time** — Not stored in the DB. Computed as `totalValue * ownerPercentage / 100` in the response converter.

**splitCount via COUNT subquery** — Computed in the `FindAll` query to avoid N+1 queries.

**groupName via LEFT JOIN** — `FindAll` and `FindByID` join group_entities to return groupName without a second query.

**Optional ownerPercentage** — If not provided by the user, equal split is calculated automatically. If provided, must be > 0 and < 100; remaining percentage is distributed equally among members.
