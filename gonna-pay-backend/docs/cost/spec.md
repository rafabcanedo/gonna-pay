# Cost — Specification

## Overview

The Cost feature allows a user to record an expense and optionally split it among a group's contacts. A cost has an owner, a name, a total value, a category, and an optional group. When a group is provided, the backend automatically calculates equal percentage splits for all members. The owner may optionally override their own percentage — if provided, the remaining percentage is distributed equally among the group members.

When a cost is created with a group, Cost_Split records are created atomically in the same database transaction. A cost can be updated after creation — only the scalar fields (`costName`, `totalValue`, `category`, `ownerPercentage`) can be changed. The group cannot be changed via update. If splits exist, their values and percentages are recalculated automatically based on the new `totalValue` and `ownerPercentage`.

---

## Endpoints

### Protected (JWT cookie required)

| Method | Path | Description |
|--------|------|-------------|
| POST | `/cost` | Create a new cost (and splits if group is provided) |
| GET | `/costs` | List all costs owned by the authenticated user |
| GET | `/cost/:id` | Get a cost by ID (includes splits inline) |
| PUT | `/cost/:id` | Update a cost's scalar fields (recalculates splits if any) |
| DELETE | `/cost/:id` | Delete a cost and all its splits |

---

## Business Rules

### Create (`POST /cost`)
- `costName`, `totalValue`, and `category` are required
- `category` must be one of: `Dinner`, `Lunch`, `Entertainment`, `Travel`, `Others`
- `totalValue` must be greater than 0
- Cost is automatically linked to the authenticated user via `userId`
- `groupId` is optional — if not provided, `ownerPercentage = 100` and no splits are created
- If `groupId` is provided:
  - The group must belong to the authenticated user — returns 404 if not found, 403 if not owner
  - The backend fetches the number of members in the group
  - If `ownerPercentage` is not provided: equal split = `100 / (memberCount + 1)`
  - If `ownerPercentage` is provided: must be > 0 and < 100; remaining `(100 - ownerPercentage)` is split equally among members
  - One Cost_Split record is created per group member, all within the same DB transaction
  - If any insert fails, the entire transaction rolls back

### List (`GET /costs`)
- Returns only costs where the authenticated user is the **owner**
- Does not return costs owned by other users
- Splits are **not** embedded in the list response — only cost summary metadata
- `ownerValue` is computed as `totalValue * ownerPercentage / 100`
- `splitCount` is the total number of Cost_Split records for that cost
- `groupName` is returned alongside `groupId` (null if no group)

### Get (`GET /cost/:id`)
- Returns 404 if the cost does not exist
- Returns 403 if the cost does not belong to the authenticated user
- Splits are embedded inline in the response

### Update (`PUT /cost/:id`)
- `costName`, `totalValue`, and `category` are required
- `ownerPercentage` is optional — if not provided, defaults to `0` (members absorb 100%)
- `groupId` cannot be changed via update
- Returns 404 if the cost does not exist
- Returns 403 if the cost does not belong to the authenticated user
- If the cost has splits, their `value` and `percentage` are recalculated atomically in the same DB transaction based on the new `totalValue` and `ownerPercentage`
- Returns the full detail response (with splits inline)

### Delete (`DELETE /cost/:id`)
- Returns 404 if the cost does not exist
- Returns 403 if the cost does not belong to the authenticated user
- Cascade deletes all Cost_Split records (handled at DB level via `ON DELETE CASCADE`)
- Returns 200 with `{ "message": "cost deleted successfully" }`

---

## Responses

### Cost object (list — no splits)
```json
{
  "id": "uuid",
  "costName": "Pizza",
  "totalValue": 100.00,
  "ownerPercentage": 25.00,
  "ownerValue": 25.00,
  "category": "Dinner",
  "groupId": "uuid or null",
  "groupName": "Road Trip or null",
  "splitCount": 3,
  "createdAt": "2024-01-10T10:00:00Z"
}
```

### Cost object (detail — splits inline)
```json
{
  "id": "uuid",
  "costName": "Pizza",
  "totalValue": 100.00,
  "ownerPercentage": 25.00,
  "ownerValue": 25.00,
  "category": "Dinner",
  "groupId": "uuid or null",
  "groupName": "Road Trip or null",
  "splitCount": 3,
  "createdAt": "2024-01-10T10:00:00Z",
  "splits": [
    { "id": "uuid", "contactId": "uuid", "contactName": "Ana", "value": 25.00, "percentage": 25.00 },
    { "id": "uuid", "contactId": "uuid", "contactName": "Bia", "value": 25.00, "percentage": 25.00 },
    { "id": "uuid", "contactId": "uuid", "contactName": "Carlos", "value": 25.00, "percentage": 25.00 }
  ]
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
