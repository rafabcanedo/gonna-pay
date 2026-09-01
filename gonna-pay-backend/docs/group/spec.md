# Group — Specification

## Overview

The Group feature allows a user to create named cost aggregators (e.g., "Beach Trip"). A group has an owner, a category, and a list of members (contacts belonging to the owner). Groups are **immutable** — they cannot be edited after creation. If a mistake occurs, the user deletes and recreates the group.

Group members are managed via sub-resource endpoints under `/group/:id/member`.

---

## Endpoints

### Protected (JWT cookie required)

| Method | Path | Description |
|--------|------|-------------|
| POST | `/group` | Create a new group |
| GET | `/groups` | List all groups owned by the authenticated user |
| GET | `/group/:id` | Get a group by ID (includes members inline) |
| DELETE | `/group/:id` | Delete a group and all its members |
| POST | `/group/:id/member` | Add a contact as a member of the group |
| DELETE | `/group/:id/member/:contactId` | Remove a contact from the group |

---

## Business Rules

### Create (`POST /group`)
- `name` and `category` are required
- `category` must be one of: `Dinner`, `Lunch`, `Entertainment`, `Travel`, `Others`
- Group is automatically linked to the authenticated user via `ownerId`
- Members can be optionally included in the creation payload as a list of `contactId`s
- Each `contactId` provided must belong to the authenticated user — returns 400 if any contact does not

### List (`GET /groups`)
- Returns only groups where the authenticated user is the **owner**
- Does not return groups owned by other users
- Members are **not** embedded in the list response (performance) — only group metadata

### Get (`GET /group/:id`)
- Returns 404 if the group does not exist
- Returns 403 if the group does not belong to the authenticated user
- Members are embedded inline in the response (see response format below)

### Delete (`DELETE /group/:id`)
- Returns 404 if the group does not exist
- Returns 403 if the group does not belong to the authenticated user
- Cascade deletes all `group_member` records for that group (handled at DB level via `ON DELETE CASCADE`)
- Returns 200 with `{ "message": "group deleted successfully" }`

### Add Member (`POST /group/:id/member`)
- Returns 404 if the group does not exist
- Returns 403 if the group does not belong to the authenticated user
- `contactId` is required in the request body
- The contact must belong to the authenticated user — returns 400 if not
- Returns 400 if the contact is already a member of the group (duplicate prevention)

### Remove Member (`DELETE /group/:id/member/:contactId`)
- Returns 404 if the group does not exist
- Returns 403 if the group does not belong to the authenticated user
- Returns 404 if the contact is not a member of the group
- Returns 200 with `{ "message": "member removed successfully" }`

---

## Responses

### Group object (list — no members)
```json
{
  "id": "uuid",
  "name": "string",
  "category": "Travel",
  "createdAt": "2024-01-10T10:00:00Z",
  "updatedAt": "2024-01-10T10:00:00Z"
}
```

### Group object (detail — members inline)
```json
{
  "id": "uuid",
  "name": "Beach Trip",
  "category": "Travel",
  "createdAt": "2024-01-10T10:00:00Z",
  "updatedAt": "2024-01-10T10:00:00Z",
  "members": [
    { "id": "uuid", "name": "Ana", "email": "ana@email.com" },
    { "id": "uuid", "name": "Bia", "email": "bia@email.com" }
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
