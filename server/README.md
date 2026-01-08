# Wildex Server

A minimal Go backend API server built with the Gin framework.

## Project Structure

```
server/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── handlers/                # HTTP handlers (controllers)
│   │   └── health_handler.go
│   ├── routes/                  # Route definitions
│   │   └── routes.go
│   ├── middleware/              # Custom middleware
│   ├── models/                  # DTOs and request/response models
│   │   └── response.go
│   └── services/                # Business logic layer
├── configs/                     # Configuration files
│   └── config.go
├── go.mod                       # Go module definition
├── go.sum                       # Dependency checksums
├── .gitignore                   # Git ignore rules
└── README.md                    # This file
```

## Prerequisites

- Go 1.23 or higher
- Git

## Installation

1. Navigate to the server directory:
```bash
cd server
```

2. Install dependencies:
```bash
go mod download
```

## Running the Server

### Development Mode

```bash
go run cmd/api/main.go
```

### Build and Run

```bash
# Build the binary
go build -o bin/server cmd/api/main.go

# Run the binary
./bin/server
```

## Configuration

The server can be configured using environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVER_PORT` | Port number for the server | `8080` |
| `ENVIRONMENT` | Environment mode (development/production) | `development` |

Example:
```bash
export SERVER_PORT=3000
export ENVIRONMENT=production
go run cmd/api/main.go
```

## API Endpoints

### Health Check
```bash
GET /health
```

Response:
```json
{
  "status": "success",
  "message": "Server is healthy",
  "data": {
    "status": "ok"
  }
}
```

### Hello World
```bash
GET /api/hello
```

Response:
```json
{
  "status": "success",
  "message": "Hello, World!"
}
```

## Testing the API

Using curl:
```bash
# Health check
curl http://localhost:8080/health

# Hello World
curl http://localhost:8080/api/hello
```

Using HTTPie:
```bash
# Health check
http GET http://localhost:8080/health

# Hello World
http GET http://localhost:8080/api/hello
```

## Development

### Adding New Routes

1. Create a new handler in `internal/handlers/`
2. Register the route in `internal/routes/routes.go`

### Adding Middleware

1. Create middleware functions in `internal/middleware/`
2. Apply middleware in `cmd/api/main.go` or specific route groups

### Adding Services

1. Create service files in `internal/services/`
2. Import and use them in handlers

## Project Architecture

This project follows standard Go project layout conventions:

- **cmd/**: Main applications for this project
- **internal/**: Private application and library code
- **configs/**: Configuration file templates or default configs

The architecture is inspired by clean architecture principles, with clear separation between:
- Handlers (Presentation layer)
- Services (Business logic layer)
- Models (Data structures)

## Tech Stack

- **Framework**: [Gin](https://github.com/gin-gonic/gin) v1.11.0
- **Language**: Go 1.24

## License

This project is part of the Wildex application.
