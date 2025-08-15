# 🚀 HR Server

> **A powerful REST API server with Telegram bot integration for HR management and user tracking**

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-Web%20Framework-green.svg)](https://gin-gonic.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-blue.svg)](https://www.postgresql.org/)
[![Swagger](https://img.shields.io/badge/Swagger-API%20Docs-orange.svg)](https://swagger.io/)

## 📋 Table of Contents

- [✨ Features](#-features)
- [🏗️ Architecture](#️-architecture)
- [🚀 Quick Start](#-quick-start)
- [⚙️ Configuration](#️-configuration)
- [📚 API Documentation](#-api-documentation)
- [🤖 Telegram Bot](#-telegram-bot)
- [🗄️ Database](#️-database)
- [🔧 Development](#-development)
- [🚀 Deployment](#-deployment)

## ✨ Features

- 🔐 **RESTful API** with JWT authentication
- 🤖 **Telegram Bot** for user registration and notifications
- 📊 **Channel Management** with unique codes and tracking
- 👥 **User Management** with Telegram integration
- 📈 **Statistics & Analytics** for channels and users
- 🔗 **Deep Link Support** for easy user onboarding
- 📝 **Swagger Documentation** with interactive API testing
- 🗄️ **PostgreSQL Database** with GORM ORM
- 🛡️ **Error Handling** with comprehensive logging
- 🚀 **Docker Support** for easy deployment

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   HTTP Client   │    │  Telegram Bot   │    │   PostgreSQL    │
│                 │    │                 │    │   Database      │
└─────────┬───────┘    └─────────┬───────┘    └─────────┬───────┘
          │                      │                      │
          ▼                      ▼                      ▼
┌─────────────────────────────────────────────────────────────────┐
│                        HR Server                                │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐              │
│  │ Controllers │  │   Services  │  │ Repositories│              │
│  └─────────────┘  └─────────────┘  └─────────────┘              │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐              │
│  │   Domain    │  │ Middleware  │  │ Background  │              │
│  │   Models    │  │             │  │   Workers   │              │
│  └─────────────┘  └─────────────┘  └─────────────┘              │
└─────────────────────────────────────────────────────────────────┘
```

## 🚀 Quick Start

### Prerequisites

- **Go 1.24+** - [Download here](https://golang.org/dl/)
- **PostgreSQL 12+** - [Download here](https://www.postgresql.org/download/)
- **Git** - [Download here](https://git-scm.com/)

### 1. Clone the Repository

```bash
git clone https://github.com/smirnoffalexx/hr_server
cd hr_server
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Set Up Environment Variables

Create a `.env` file in the root directory:

```bash
# Database Configuration
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=your_username
POSTGRES_PASSWORD=your_password
POSTGRES_DB=hr_server

# Server Configuration
HTTP_PORT=8080
ENVIRONMENT=development
LOGL=debug

# Authentication
AUTH_TOKEN=your_secret_auth_token

# Telegram Bot
TG_BOT_TOKEN=your_telegram_bot_token
```

### 4. Set Up Database

```bash
# Create database
createdb hr_server

# Or using psql
psql -U postgres -c "CREATE DATABASE hr_server;"
```

### 5. Run the Application

```bash
# Development mode
go run cmd/app/main.go

# Or build and run
go build -o bin/hr_server cmd/app/main.go
./bin/hr_server
```

### 6. Verify Installation

- 🌐 **API Server**: http://localhost:8080/api/health
- 📚 **Swagger Docs**: http://localhost:8080/api/swagger/
- 🤖 **Telegram Bot**: Start your bot with `/start`

## ⚙️ Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `POSTGRES_HOST` | PostgreSQL host | localhost | ✅ |
| `POSTGRES_PORT` | PostgreSQL port | 5432 | ✅ |
| `POSTGRES_USER` | Database username | - | ✅ |
| `POSTGRES_PASSWORD` | Database password | - | ✅ |
| `POSTGRES_DB` | Database name | hr_server | ✅ |
| `HTTP_PORT` | Server port | 8080 | ✅ |
| `ENVIRONMENT` | Environment (dev/prod) | development | ✅ |
| `LOGL` | Log level (debug/info/warn/error) | debug | ✅ |
| `AUTH_TOKEN` | API authentication token | - | ✅ |
| `TG_BOT_TOKEN` | Telegram bot token | - | ✅ |

### Docker Setup

```bash
# Build and run with Docker Compose
docker-compose up -d

# Or build manually
docker build -t hr-server .
docker run -p 8080:8080 hr-server
```

## 📚 API Documentation

### Interactive Documentation

- **Swagger UI**: [http://localhost:8080/api/swagger/](http://localhost:8080/api/swagger/)
- **Full Documentation**: [http://localhost:8080/api/swagger/index.html#/](http://localhost:8080/api/swagger/index.html#/)

### API Endpoints

#### 🔐 Authentication
All API endpoints require the `X-Auth-Token` header for authentication.

#### 👥 User Management
- `GET /api/admin/users` - Get all users
- `POST /api/admin/users` - Create new user

#### 📢 Channel Management
- `POST /api/channel/generate` - Generate channel code
- `GET /api/channel/{code}` - Get channel by code
- `POST /api/admin/channel/bulk` - Generate multiple channels with different names
- `GET /api/admin/channels` - Get all channels

#### 📊 Statistics
- `GET /api/admin/stats` - Get overall statistics
- `GET /api/stats/channels` - Get channel statistics
- `GET /api/stats/channel/{code}` - Get specific channel stats

#### 🔔 Notifications
- `POST /api/admin/notifications` - Send notification to users

### 📝 API Usage Examples

#### Generate Multiple Channels with Different Names
```bash
curl -X POST "http://localhost:8080/api/admin/channel/bulk" \
  -H "X-Auth-Token: your_auth_token" \
  -H "Content-Type: application/json" \
  -d '{
    "channel_names": [
      "Marketing Team",
      "Sales Department", 
      "Engineering Team",
      "HR Department"
    ]
  }'
```

**Response:**
```json
[
  {
    "id": 1,
    "name": "Marketing Team",
    "code": "a1b2c3d4",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  },
  {
    "id": 2,
    "name": "Sales Department",
    "code": "e5f6g7h8",
    "created_at": "2024-01-15T10:30:01Z",
    "updated_at": "2024-01-15T10:30:01Z"
  }
]
```

## 🤖 Telegram Bot

### Features
- **User Registration**: `/start [channel_code]`
- **Deep Link Support**: Handle complex parameters
- **Channel Association**: Link users to specific channels
- **Error Handling**: Graceful handling of invalid codes

### Bot Commands
```
/start [channel_code] - Register user and associate with channel
```

### Deep Link Format
```
https://t.me/YourBot?start=eyJjaGFubmVsQ29kZSI6IkFCQzEyMyJ9
```

## 🗄️ Database

### Schema Overview

#### Users Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    telegram_id BIGINT UNIQUE NOT NULL,
    username VARCHAR(255),
    channel_id INTEGER REFERENCES channels(id),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
```

#### Channels Table
```sql
CREATE TABLE channels (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
```

## 🔧 Development

### Project Structure

```
hr_server/
├── cmd/
│   └── app/
│       └── main.go              # Application entry point
├── config/
│   └── config.go                # Configuration management
├── internal/
│   ├── api/
│   │   └── http/
│   │       ├── controllers/     # HTTP controllers
│   │       ├── middleware/      # HTTP middleware
│   │       ├── routing/         # Route definitions
│   │       └── docs/            # Swagger documentation
│   ├── app/
│   │   ├── app.go              # Application setup
│   │   └── logger.go           # Logging configuration
│   ├── background/
│   │   └── tgbot.go            # Telegram bot worker
│   ├── domain/
│   │   ├── user.go             # User domain model
│   │   ├── channel.go          # Channel domain model
│   │   └── notification.go     # Notification domain model
│   ├── infrastructure/
│   │   └── database.go         # Database connection
│   ├── register/
│   │   └── storage.go          # Dependency injection
│   ├── repository/
│   │   ├── user_postgres.go    # User repository
│   │   └── channel_postgres.go # Channel repository
│   └── service/
│       ├── user_service.go     # User business logic
│       ├── channel_service.go  # Channel business logic
│       └── notification_service.go # Notification logic
├── Dockerfile                   # Docker configuration
├── docker-compose.yml          # Docker Compose setup
├── go.mod                      # Go modules
├── go.sum                      # Go modules checksum
└── README.md                   # This file
```

### Development Commands

```bash
# Run tests
go test ./...

# Run with hot reload (if using air)
air

# Generate Swagger docs
swag init -g internal/api/http/routing/routing.go -o internal/api/http/docs

# Format code
go fmt ./...

# Lint code
golangci-lint run

# Build for production
go build -ldflags="-s -w" -o bin/hr_server cmd/app/main.go
```

## 🚀 Deployment

### Docker Deployment

```bash
# Build image
docker build -t hr-server:latest .

# Run containers
docker compose up
```

---

<div align="center">

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Gin-00AC47?style=for-the-badge&logo=gin&logoColor=white)](https://gin-gonic.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)

</div>
