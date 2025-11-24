# go-api-skeleton
First go project 

# 🚀 Go API Skeleton - E-Commerce Backend

My first Go backend project! A RESTful API built with Go, MySQL, and Gorilla Mux featuring full CRUD operations.

## 📋 Table of Contents

- [About](#about)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Database Setup](#database-setup)
- [Usage](#usage)
- [What I Learned](#what-i-learned)
- [Future Improvements](#future-improvements)

## 🎯 About

This is a learning project where I built my first backend API using Go. The goal was to understand:

- RESTful API design
- Go's HTTP server and routing
- Database connections and CRUD operations
- Environment configuration
- Project structure and organization

## ✨ Features

- ✅ RESTful API architecture
- ✅ MySQL database integration
- ✅ User authentication endpoints (login/register)
- ✅ Environment-based configuration
- ✅ Clean project structure
- ✅ Full CRUD operations support
- ✅ Error handling and logging

## 🛠️ Tech Stack

- **Language:** Go 1.21+
- **Web Framework:** [Gorilla Mux](https://github.com/gorilla/mux) - HTTP router and URL matcher
- **Database:** MySQL 8.0+
- **Driver:** [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- **Config:** [godotenv](https://github.com/joho/godotenv) - Environment variable management

## 📁 Project Structure

```
backend/
├── bin/                    # Compiled binaries
│   └── ecom               # Built executable
├── cmd/                    # Application entry points
│   ├── main.go            # Main application entry
│   ├── api/               # API server setup
│   │   └── api.go         # Server configuration & routing
│   └── migrate/           # Database migrations (future)
├── config/                 # Configuration management
│   └── env.go             # Environment variables handler
├── db/                     # Database layer
│   └── db.go              # Database connection setup
├── service/                # Business logic layer
│   └── user/              # User service
│       └── routes.go      # User routes and handlers
├── .env                    # Environment variables (not in git)
├── .env.example           # Environment template
├── go.mod                 # Go module dependencies
├── go.sum                 # Dependency checksums
├── makefile               # Build automation
└── README.md              # This file
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher ([Download](https://golang.org/dl/))
- MySQL 8.0 or higher ([Download](https://dev.mysql.com/downloads/mysql/))
- Git

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/JaniduP2003/go-api-skeleton.git
   cd go-api-skeleton/backend
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Set up environment variables**

   Create a `.env` file in the `backend` directory:

   ```bash
   cp .env.example .env
   ```

   Edit `.env` with your configuration:

   ```env
   # Database Configuration
   DB_USER=root
   DB_PASSWORD=your_password
   DB_HOST=127.0.0.1
   DB_PORT=3306
   DB_NAME=ecom

   # Server Configuration
   PUBLIC_HOST=http://localhost
   PORT=8080
   ```

4. **Create the database**

   ```bash
   mysql -u root -p
   ```

   In MySQL shell:

   ```sql
   CREATE DATABASE ecom;
   exit;
   ```

5. **Build and run**

   ```bash
   make build
   make run
   ```

   Or simply:

   ```bash
   make run
   ```

## 📡 API Endpoints

### Base URL

```
http://localhost:8080/api/v1
```

### User Endpoints

| Method | Endpoint           | Description       | Status         |
| ------ | ------------------ | ----------------- | -------------- |
| POST   | `/api/v1/login`    | User login        | 🚧 In Progress |
| POST   | `/api/v1/register` | User registration | 🚧 In Progress |

### Example Request (Coming Soon)

```bash
# Register a new user
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "secure_password"
  }'
```

## 🗄️ Database Setup

### Create the Database

```sql
CREATE DATABASE ecom;
USE ecom;

-- Example users table (to be implemented)
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### Migration (Future Feature)

Migrations will be handled through `cmd/migrate/` directory.

## 💻 Usage

### Using Make Commands

```bash
# Build the project
make build

# Run the application
make run

# Run tests (to be implemented)
make test

# Clean build artifacts
make clean
```

### Manual Running

```bash
# Build
go build -o bin/ecom cmd/main.go

# Run
./bin/ecom
```

### Check if it's running

```bash
# The server should start and show:
2025/11/23 20:41:32 DB: successfully connected!!!
2025/11/23 20:41:32 Listening on :8080
```

## 📚 What I Learned

Building this project taught me:

1. **Go Fundamentals**

   - Package structure and organization
   - Pointers and receivers
   - Struct types and methods
   - Error handling patterns

2. **Web Development**

   - HTTP server setup with `net/http`
   - Routing with Gorilla Mux
   - Request/Response handling
   - Middleware concepts

3. **Database Integration**

   - MySQL connection pooling
   - SQL driver usage
   - Database initialization
   - Environment-based configuration

4. **Best Practices**
   - Project structure conventions
   - Environment variable management
   - Separation of concerns
   - Clean code principles

## 🔮 Future Improvements

- [ ] Complete CRUD operations for users
- [ ] Add JWT authentication
- [ ] Implement password hashing (bcrypt)
- [ ] Add input validation
- [ ] Create database migrations
- [ ] Add unit tests
- [ ] Implement middleware (logging, CORS, etc.)
- [ ] Add more entities (products, orders, etc.)
- [ ] API documentation with Swagger
- [ ] Docker containerization
- [ ] Add rate limiting
- [ ] Implement refresh tokens
- [ ] Add API versioning
- [ ] Create comprehensive error handling

## 🐛 Known Issues

- [ ] `/register` route currently points to wrong handler
- [ ] Duplicate `/login` route registration
- [ ] Handler implementations are empty (in progress)
- [ ] No input validation yet
- [ ] No authentication/authorization

## 🤝 Contributing

This is a personal learning project, but suggestions and feedback are welcome!

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 👤 Author

**Janidu Poornima**

- GitHub: [@JaniduP2003](https://github.com/JaniduP2003)

## 🙏 Acknowledgments

- [Go Documentation](https://go.dev/doc/)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Go MySQL Driver](https://github.com/go-sql-driver/mysql)
- The Go community for amazing resources!

---

⭐ If you found this helpful, please star the repository!

**Status:** 🚧 Active Development | First Go Project | Learning Journey
