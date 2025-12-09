# Go JWT CRUD API

A secure and modular RESTful CRUD API built with Golang using Gin framework, JWT authentication, and PostgreSQL database.

---

## 🚀 Features

- User Registration & Login
- JWT-based Authentication
- Secure Password Hashing (bcrypt)
- CRUD Operations (Create, Read, Update, Delete)
- Protected Routes using JWT Middleware
- PostgreSQL Database with GORM ORM
- Modular & Clean Project Structure
- Separate files for Create, Read, Update, Delete
- Environment-based configuration (.env)

---

## 🛠 Tech Stack

- **Language:** Golang
- **Framework:** Gin
- **Authentication:** JWT (JSON Web Token)
- **Database:** PostgreSQL
- **ORM:** GORM
- **Security:** bcrypt (password hashing)
- **Config:** godotenv
- **API Style:** RESTful API

---

## 📁 Project Structure

go-jwt-crud/
│
├── main.go
├── go.mod
├── .env
│
├── database/
│ └── db.go
│
├── models/
│ ├── user.go
│ └── post.go
│
├── controllers/
│ ├── auth_controller.go
│ └── post/
│ ├── create.go
│ ├── read.go
│ ├── update.go
│ └── delete.go
│
├── middlewares/
│ └── jwt_middleware.go
│
├── routes/
│ └── routes.go
│
├── utils/
│ ├── jwt.go
│ └── hash.go
│
└── README.md

---

## 🔐 Authentication Flow

1. User registers with email and password  
2. Password is hashed using bcrypt  
3. User logs in and receives JWT token  
4. JWT token is sent in `Authorization` header  
5. Protected routes validate JWT using middleware  

PORT=8080

DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=jwtcrud
DB_PORT=5432

JWT_SECRET=your_secret_key


---

## ▶️ How to Run the Project

```bash
go mod init go-jwt-crud
go mod tidy
go run main.go
