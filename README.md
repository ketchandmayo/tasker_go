# tasker_go

A Go REST API for task management with JWT authentication and task auto-evaluation (`judge`) via local scoring + Gemini-generated comments.

## Features

- User registration and login
- JWT authentication
- Task CRUD
- Task evaluation (`score 1..10`) with text comment
- Evaluation caching by task hash with background regeneration
- PostgreSQL + GORM + automatic migrations on startup

## Stack

- Go `1.25`
- HTTP Router: `gorilla/mux`
- ORM: `gorm` + `postgres`
- Validation: `go-playground/validator`
- JWT: `golang-jwt/jwt/v5`
- LLM: `google.golang.org/genai` (Gemini)

## Environment Variables

Copy `.env.example` to `.env` and fill in values:

```env
JWT_SECRET=
PORT=8080

GEMINI_API_KEY=
GEMINI_MODEL=gemma-3-12b-it
GEMINI_SOCKS5_PROXY=

DB_HOST=
DB_USER=
DB_PASS=
DB_NAME=
DB_PORT=5432
```

Required for local run:

- `JWT_SECRET`
- `DB_HOST`
- `DB_USER`
- `DB_PASS`
- `DB_NAME`
- `DB_PORT`

Notes:

- The app loads `.env` via `godotenv` on startup.
- If `.env` is missing, system environment variables are used.
- `tasks`, `users`, and `judges` tables are created/updated automatically (`AutoMigrate`).

## Run

1. Start PostgreSQL.
2. Fill `.env`.
3. Install dependencies and start the server:

```bash
go mod download
go run .
```

By default, the server listens on `http://localhost:8080`.

## API

Error format:

```json
{ "error": "message" }
```

### 1. Register

`POST /register`

Request:

```json
{
  "email": "user@example.com",
  "password": "secret"
}
```

Response `201`:

```json
{ "status": "done" }
```

### 2. Login

`POST /login`

Request:

```json
{
  "email": "user@example.com",
  "password": "secret"
}
```

Response `200`:

```json
{ "token": "<jwt>" }
```

### 3. Create Task

`POST /tasks` (Bearer token)

Request:

```json
{
  "title": "Buy milk",
  "description": "After work"
}
```

Response `201`:

```json
{ "id": 1 }
```

### 4. List Tasks

`GET /tasks` (Bearer token)

Response `200`:

```json
[
  {
    "id": 1,
    "title": "Buy milk",
    "status": "new"
  }
]
```

### 5. Get Task

`GET /tasks/{id}` (Bearer token)

Response `200`:

```json
{
  "id": 1,
  "title": "Buy milk",
  "description": "After work",
  "status": "new",
  "created_at": "2026-02-16T10:00:00Z"
}
```

### 6. Update Task

`PATCH /tasks/{id}` (Bearer token)

Request (all fields are optional):

```json
{
  "title": "Buy milk and bread",
  "description": "After work",
  "status": "done"
}
```

Allowed `status` values: `new`, `done`, `archived`.

Response `200`: updated task object.

### 7. Delete Task

`DELETE /tasks/{id}` (Bearer token)

Response `200`:

```json
{ "status": "deleted" }
```

### 8. Task Evaluation (judge)

`GET /tasks/{id}/judge` (Bearer token)

Response `200`:

```json
{
  "task_id": 1,
  "score": 7,
  "text": "Good task, the path is clear.",
  "preliminary": true
}
```

- `preliminary: true` means a quick preliminary evaluation (returned immediately).
- A full evaluation with Gemini text is generated in the background.
- The next request for the same task version usually returns `preliminary: false`.

## Quick curl Example

```bash
# 1) Register
curl -X POST http://localhost:8080/register \
  -H 'Content-Type: application/json' \
  -d '{"email":"user@example.com","password":"secret"}'

# 2) Login
TOKEN=$(curl -s -X POST http://localhost:8080/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"user@example.com","password":"secret"}' | jq -r '.token')

# 3) Create task
curl -X POST http://localhost:8080/tasks \
  -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"title":"Read RFC","description":"Write a short summary"}'

# 4) Get evaluation
curl -X GET http://localhost:8080/tasks/1/judge \
  -H "Authorization: Bearer $TOKEN"
```

## Project Structure

```text
internal/
  app/                # application bootstrap
  config/             # DB/JWT config
  auth/               # JWT and passwords
  models/             # GORM models
  repository/         # interfaces and gorm implementations
  service/            # business logic
  analysis/           # scoring and judge prompt creation
  llm/gemini/         # Gemini client
  transport/http/     # router, middleware, dto, handlers, responder
```

## Limitations and Notes

- CORS is allowed for `http://localhost:*`.
- JWT lifetime is `24h`.
- `judge` has regeneration throttling (cooldown ~30 seconds).
- `title` validation: required, max 100 characters.
