# Book API with Fiber and Go

![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go)
![Fiber](https://img.shields.io/badge/Fiber-2.52-00ADD8)
![SQLite](https://img.shields.io/badge/SQLite-3-003B57?logo=sqlite)
![Docker](https://img.shields.io/badge/Docker-✓-2496ED?logo=docker)

A high-performance REST API for book management built with Go, Fiber, and SQLite. Supports full CRUD operations with pagination, search, and Docker deployment.

## Features

- **Full CRUD Operations**
- **Advanced Search** by title/author
- **Pagination** with page size control
- **Docker** containerization
- **Unit Tests** with 85%+ coverage
- **Swagger Documentation**

## 🚀 Quick Start

### Prerequisites
- Go 1.23+
- SQLite3
- Docker (optional)

### Installation
```bash
git clone https://github.com/yourusername/book-api.git
cd book-api
go mod download

## Tech Stack

- **Framework**: [Fiber](https://gofiber.io/)
- **Database**: SQLite (with GORM ORM)
- **Containerization**: Docker
- **Testing**: Testify

## API Endpoints

| Method | Endpoint                | Description                     |
|--------|-------------------------|---------------------------------|
| POST   | `/api/books`            | Create a new book               |
| GET    | `/api/books`            | Get all books                   |
| GET    | `/api/books/:id`        | Get single book by ID           |
| PUT    | `/api/books/:id`        | Update a book                   |
| DELETE | `/api/books/:id`        | Delete a book                   |
| GET    | `/api/books/paginated`  | Get paginated results           |
| GET    | `/api/books/search`     | Search books by title/author    |
| DELETE | `/api/books`            | Delete all books (Admin only)   |
