# Contact — Specification

## Overview

The Contact feature allows a user to manage their personal contact list. Contacts are people the user interacts with — they can be added to groups and included in cost splits.

---

## Endpoints

### Protected (JWT cookie required)

| Method | Path | Description |
|--------|------|-------------|
| POST | `/contacts` | Create a new contact |
| GET | `/contacts` | List all contacts of the authenticated user |
| GET | `/contacts/:id` | Get a contact by ID |
| PUT | `/contacts/:id` | Update a contact |
| DELETE | `/contacts/:id` | Delete a contact |

---

## Business Rules

### Create (`POST /contacts`)
- `name` and `category` are required
- `email` and `phone` are required
- `email` must be a valid format if provided
- Contact is automatically linked to the authenticated user via `ownerId`

### List (`GET /contacts`)
- Returns only contacts belonging to the authenticated user
- Does not return contacts from other users

### Get (`GET /contacts/:id`)
- Returns 404 if the contact does not exist
- Returns 403 if the contact does not belong to the authenticated user

### Update (`PUT /contacts/:id`)
- All fields are optional
- `email` must be valid format if provided
- Returns 404 if the contact does not exist
- Returns 403 if the contact does not belong to the authenticated user

### Delete (`DELETE /contacts/:id`)
- Returns 404 if the contact does not exist
- Returns 403 if the contact does not belong to the authenticated user
- Returns 200 with `{ "message": "contact deleted successfully" }`

---

## Responses

### Contact object
```json
{
  "id": "uuid",
  "name": "string",
  "email": "string",
  "phone": "string",
  "category": "Family | Friend | Work"
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
