# Notes Management API

REST API sederhana untuk mengelola catatan (Notes Management System) menggunakan Golang, Gin Framework, GORM, MySQL, dan JWT Authentication.

---

# Features

* User Registration
* User Login
* User Logout
* JWT Authentication
* Create Note
* Get All Notes (Pagination)
* Get Note Detail
* Update Note
* Delete Note (Soft Delete)
* Search Notes by Title
* Filter Notes by Status

---

# Tech Stack

* Golang
* Gin Framework
* GORM
* MySQL
* JWT Authentication

---

# Project Structure

```
notes_api
│
├── cmd
│   └── main.go
├── configs
├── database
├── docs
├── internal
│   ├── dto
│   ├── handler
│   ├── middleware
│   ├── model
│   ├── repository
│   ├── routes
│   ├── service
│   └── utils
├── .env
├── go.mod
└── README.md
```

---

# Installation

## 1. Clone Repository

```bash
git clone <repository-url>
cd notes_api
```

## 2. Install Dependencies

```bash
go mod tidy
```

## 3. Create Database

Create a MySQL database named:

```
notes_db
```

Import:

```
database/schema.sql
```

## 4. Configure Environment

Create a `.env` file.

Example:

```env
APP_PORT=8080

DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=notes_db

JWT_SECRET=your_secret_key
```

## 5. Run Application

```bash
go run cmd/main.go
```

Server:

```
http://localhost:8080
```

---

# API Endpoints

## Authentication

| Method | Endpoint      |
| ------ | ------------- |
| POST   | /api/register |
| POST   | /api/login    |
| POST   | /api/logout   |

## Notes

| Method | Endpoint       |
| ------ | -------------- |
| POST   | /api/notes     |
| GET    | /api/notes     |
| GET    | /api/notes/:id |
| PUT    | /api/notes/:id |
| DELETE | /api/notes/:id |

---

# Query Parameters

Search by title

```
GET /api/notes?search=meeting
```

Filter by status

```
GET /api/notes?status=active
```

Pagination

```
GET /api/notes?page=1&limit=10
```

---

# Authentication

Protected endpoints require JWT Token.

Example:

```
Authorization: Bearer <your_token>
```

---

# Database

Database schema is available in:

```
database/schema.sql
```

---

# Deployment

Deployment URL:

```
<deployment-url>
```

If the application has not been deployed yet, replace it after deployment.
