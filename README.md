# Go Roadmap API

REST API sederhana menggunakan **Golang**, **Gin**, dan **PostgreSQL** dengan arsitektur **Handler → Service → Repository**.

## Tech Stack

* **Go**
* **Gin** – HTTP web framework
* **PostgreSQL** – Database
* **SQLX** – SQL helper library
* **ULID** – Unique ID generator
* **godotenv** – Environment variable loader

---

# Project Structure

```
go-roadmap
│
├── main.go
│
├── config
│   └── database.go
│
├── handlers
│   ├── user_handler.go
│   ├── book_handler.go
│   └── product_handler.go
│
├── services
│   ├── user_service.go
│   ├── book_service.go
│   └── product_service.go
│
├── repository
│   ├── user_repository.go
│   ├── book_repository.go
│   └── product_repository.go
│
├── models
│   ├── user.go
│   ├── book.go
│   └── product.go
│
└── .env
```

---

# Architecture

```
Request
   ↓
Handler
   ↓
Service
   ↓
Repository
   ↓
PostgreSQL
```

### Handler

Menghandle HTTP request dan response.

### Service

Berisi business logic.

### Repository

Berinteraksi langsung dengan database.

---

# Installation

Clone repository:

```bash
git clone https://github.com/yourusername/go-roadmap.git
cd go-roadmap
```

Install dependencies:

```bash
go mod tidy
```

---

# Environment Configuration

Buat file `.env` di root project:

```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=go_roadmap
DB_PORT=5432
```

---

# Run Application

```bash
go run main.go
```

Server akan berjalan di:

```
http://localhost:8080
```

---

# API Endpoints

## User

| Method | Endpoint     | Description   |
| ------ | ------------ | ------------- |
| GET    | `/api/users` | Get all users |
| POST   | `/api/users` | Create user   |

Example Request:

```json
{
  "username": "johndoe",
  "name": "John Doe",
  "password": "123456"
}
```

---

## Book

| Method | Endpoint         | Description   |
| ------ | ---------------- | ------------- |
| GET    | `/api/books`     | Get all books |
| POST   | `/api/books`     | Create book   |
| PUT    | `/api/books/:id` | Update book   |
| DELETE | `/api/books/:id` | Delete book   |

Example Request:

```json
{
  "name_book": "Atomic Habits",
  "genre": "Self Development"
}
```

---

## Product

| Method | Endpoint            | Description      |
| ------ | ------------------- | ---------------- |
| GET    | `/api/products`     | Get all products |
| POST   | `/api/products`     | Create product   |
| PUT    | `/api/products/:id` | Update product   |
| DELETE | `/api/products/:id` | Delete product   |

---

# Database Example

Create table example for books:

```sql
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    name_book TEXT,
    genre TEXT
);
```

---

# Example Curl

Create book:

```bash
curl -X POST http://localhost:8080/api/books \
-H "Content-Type: application/json" \
-d '{
"name_book":"Clean Code",
"genre":"Programming"
}'
```

---

# Future Improvements

* JWT Authentication
* Pagination
* Validation
* Logging middleware
* Docker support
* Unit testing

---

# Author

Go Backend Learning Project
