# Voices Backend (Go + MySQL)

This is the backend for **Voices**, a platform for citizens to anonymously share public engagements with trustees — such as teachers, police, and politicians. Trustees have public profiles, while citizen data is pseudonymous.

---

## 🚀 Features

- OAuth-ready authentication (scaffolded)
- MySQL database (via Docker)
- RESTful API in Go using `net/http`
- No external Go dependencies (only standard library + `mysql` driver)
- Extendable structure for future features

---

## 🗂️ Project Structure

```
voices-be/
├── cmd/             # Entry point
├── config/          # Environment access
├── db/              # DB connection and models
├── handlers/        # API route logic
├── routes/          # HTTP routing
├── migrations/      # SQL migrations
├── go.mod
└── run.sh           # Startup script
```

---

## 🛠️ Setup

### 1. Clone & Navigate

```bash
git clone <repo-url> && cd voices-be
```

### 2. Start MySQL via Docker

```bash
docker-compose up -d
```

### 3. Set Env Vars & Run

```bash
./run.sh
```

---

## 🧪 API Endpoints

### 📍 Trustees

#### `POST /trustees`
Create a trustee profile.
```json
{
  "user_id": 2,
  "name": "Jane Doe",
  "position": "Principal",
  "work_location": "Lincoln High School"
}
```

#### `GET /trustees`
List all trustees.

---

### 📍 Engagements

#### `POST /engagements`
Create a new citizen engagement.
```json
{
  "citizen_id": 1,
  "trustee_id": 2,
  "category": "positive",
  "comment": "Very helpful!",
  "tags": ["kind", "professional"],
  "location": "City Hall"
}
```

#### `GET /engagements?trustee_id=2`
Get all engagements for a specific trustee.

---

### 📍 Citizens

#### `GET /citizens/{id}`
Get anonymized citizen profile with all their engagements.

---

## 🧰 Dev Tools

### Run Migrations (using CLI like `golang-migrate`)

```bash
migrate -path ./migrations -database "mysql://voices_user:voices_pass@tcp(localhost:3306)/voices" up
```

---

## ✅ Next Steps

- Integrate OAuth for user auth
- Add rate limiting / moderation tools
- Add pagination and filtering to listings
- Create OpenAPI / Swagger docs

---

## 🧑‍💻 Author

Built for MVP deployment with rapid iteration in mind. No external frameworks or third-party dependencies.
