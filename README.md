# 🚀 URL Shortener - Go + Fiber + PostgreSQL

## 🛠 Tech Stack

- **Golang** - High-performance backend
- **Fiber** - Minimalist web framework
- **PostgreSQL** - Robust and scalable database
- **Docker** - Containerized for easy deployment
- **Golang-Migrate** - Database version control

---

## 🔥 Features

✅ Shorten long URLs into concise and shareable links
✅ Auto-expiring URLs after **7 days**
✅ Persistent storage using **PostgreSQL**
✅ Auto database migration using **golang-migrate**
✅ Lightweight and optimized with **Fiber**
✅ Fully containerized with **Docker**

---

## 📌 API Endpoints

### 🔹 1. Shorten a URL

**Request:**

```bash
curl -X POST http://localhost:3000/shorten \
     -H "Content-Type: application/json" \
     -d '{"url":"https://example.com"}'
```

**Response:**

```json
{
  "short_url": "http://localhost:3000/abc123"
}
```

### 🔹 2. Redirect to Original URL

**Request:**

```bash
curl -v http://localhost:3000/abc123
```

**Response:**

```http
HTTP/1.1 302 Found
Location: https://example.com
```

---

## 🚀 Getting Started

### 🔹 1. Clone the repository

```sh
git clone https://github.com/yourusername/url-shortener.git
cd url-shortener
```

### 🔹 2. Setup `.env` file

Create a `.env` file and add the following:

```
DATABASE_URL=postgres://user:password@localhost:5432/url_shortener?sslmode=disable
REDIS_URL=redis://localhost:6379
```

### 🔹 3. Run with Docker (PostgreSQL & Redis)

```sh
docker-compose up -d db redis
```

### 🔹 4. Run the application

```sh
go run src/cmd/main.go
```

### 🔹 5. Test API

```sh
curl -X POST http://localhost:3000/shorten -H "Content-Type: application/json" -d '{"url":"https://example.com"}'
```

---

## 📜 Project Structure

```
url-shortener/
├── src/
│   ├── cmd/
│   │   ├── main.go       # Entry point
│   ├── domain/
│   │   ├── url.go        # URL entity
│   ├── handler/
│   │   ├── url_handler.go  # HTTP Handlers
│   ├── infrastructure/
│   │   ├── db.go          # Database connection & migration
│   ├── repository/
│   │   ├── url_repository.go  # Database interaction
│   ├── usecase/
│   │   ├── url_usecase.go  # Business logic
├── migrations/
│   ├── 000001_create_urls_table.up.sql  # DB Schema
│   ├── 000001_create_urls_table.down.sql
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

---

## 🛠 Development Notes

- **Database Migration:** Automatically applied when running the application.
- **Containerization:** PostgreSQL & Redis run inside Docker.
- **Performance:** Optimized with Fiber for high-speed handling.
