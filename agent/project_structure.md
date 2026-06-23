# Project Structure

## Overview
This document outlines the current structure of the Minecraft project and provides guidance for potential future expansion.

## Current Directory Structure
```
minecraft/
├── cmd/
│   └── mine/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── db/
│   │   └── postgres.go
│   └── server/
│       ├── controllers/
│       │   └── book-controller.go
│       ├── routers/
│       │   └── book-router.go
│       └── services/
│           └── book-service.go
├── agent/
└── CLAUDE.md
```

## Components
- **cmd/mine/**: Application entry point, initializes the HTTP server and manages graceful shutdown.
  - `main.go`: Sets up configuration, database connection, service/controller/router initialization, and handles signal-based graceful shutdown.

- **internal/config/**: Contains configuration management code.
  - `config.go`: Defines the Config struct with database and server settings. Provides Load() function that reads environment variables (with godotenv support) and fallback defaults.

- **internal/db/**: Contains database interaction code.
  - `postgres.go`: Sets up PostgreSQL connection pool using lib/pq driver. Configures max open/idle connections, connection lifetime, and provides Ping() validation.

- **internal/server/routers/**: Handles HTTP routing logic, mapping URL paths to controllers.
  - `book-router.go`: Defines routes for book-related operations (GET /books, POST /books/create, GET /books/get, etc.).

- **internal/server/controllers/**: Contains business logic and request handlers for API endpoints.
  - `book-controller.go`: Implements CRUD operations by delegating to the service layer. Handles JSON encoding/decoding, HTTP status codes, and error responses.

- **internal/server/services/**: Core business logic layer.
  - `book-service.go`: Defines Book struct (ID, Title, Author), implements in-memory storage with Create, GetAll, GetByID, Update, Delete operations. Includes custom errors (ErrBookNotFound).

- **agent/**: Placeholder directory for additional agent-related files.
- **CLAUDE.md**: Likely contains project-specific instructions or guidelines.

## Architecture Pattern
The application follows a layered architecture:
```
Router → Controller → Service → Database
```
- **Router**: Maps HTTP endpoints to controller methods
- **Controller**: Handles HTTP requests/responses, validates input, delegates business logic
- **Service**: Contains core business rules and data operations
- **Database**: Provides persistence layer (currently in-memory, PostgreSQL configured but not yet integrated)

## Future Expansion
- Integrate the PostgreSQL database with the BookService (replace in-memory storage)
- Add migration scripts for database schema initialization
- Implement middleware for authentication, logging, and error handling
- Add more resource routers and controllers (e.g., users, inventory)
- Consider introducing a repository pattern to separate data access from business logic
- Add input validation using packages like `github.com/go-playground/validator`
- Implement proper API versioning strategy

## Conclusion
This structure provides a solid foundation for the Minecrafrt project, allowing for easy scalability and maintainability. Future enhancements will focus on expanding routing capabilities, server management, and database interactions.

