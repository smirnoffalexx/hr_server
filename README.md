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
- [🔔 Simplified Notification System](#-simplified-notification-system)
- [🚀 Deployment](#-deployment)

## ✨ Features

- 🔐 **RESTful API** with token-based authentication
- 🤖 **Telegram Bot** for user registration and notifications
- 📊 **Channel Management** with unique codes and tracking
- 👥 **User Management** with Telegram integration
- 📈 **Statistics & Analytics** for channels and users
- 📝 **Swagger Documentation** with interactive API testing
- 🗄️ **PostgreSQL Database** with GORM ORM
- 🛡️ **Error Handling** with comprehensive error wrapping
- 🚀 **Docker Support** for easy deployment
- 🔔 **Simplified Notification System** - send to ALL users, no exceptions

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
│  ┌─────────────┐  ┌─────────────┐                               │
│  │   Domain    │  │ Middleware  │                               │
│  │   Models    │  │             │                               │
│  └─────────────┘  └─────────────┘                               │
└─────────────────────────────────────────────────────────────────┘
```

### Key Design Principles
- **Domain-Driven Design**: Models separated into `domain/` package
- **DTO Pattern**: Request/Response models in controller-specific DTO folders
- **Repository Pattern**: Data access abstraction
- **Service Layer**: Business logic encapsulation
- **Error Wrapping**: Comprehensive error context preservation
- **Simplified Notifications**: Single method to send to ALL users

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
docker build -t hr-server .
docker-compose up -d
```

## 📚 API Documentation

### Interactive Documentation

- **Swagger UI**: [http://localhost:8080/api/swagger/](http://localhost:8080/api/swagger/)
- **Full Documentation**: [http://localhost:8080/api/swagger/index.html#/](http://localhost:8080/api/swagger/index.html#/)

### API Endpoints

#### 🔐 Authentication
All API endpoints require the `X-Auth-Token` header for authentication.

#### 👥 User Management
- `GET /api/users` - Get all users

#### 📢 Channel Management
- `POST /api/channel/generate` - Generate channel code
- `GET /api/channel/{code}` - Get channel by code
- `POST /api/channel/bulk` - Generate multiple channels with different names
- `GET /api/channels` - Get all channels

#### 🔔 Notifications
- `POST /api/notifications` - Send notification to ALL users (no exceptions, no filters)

### 📝 API Usage Examples

#### Send Notification to ALL Users
```bash
curl -X POST "http://localhost:8080/api/notifications" \
  -H "X-Auth-Token: your_auth_token" \
  -H "Content-Type: application/json" \
  -d '{
    "message": "🎉 Welcome to our platform!",
    "image_url": "https://example.com/welcome.jpg"
  }'
```

**Response:**
```json
{}
```

**What happens:**
1. ✅ Loads ALL users from database in batches of 20
2. ✅ Creates jobs for EVERY user (no filters, no exceptions)
3. ✅ Uses 5 workers for concurrent processing
4. ✅ Sends to absolutely everyone in the system
5. ✅ Rate limiting: 200 ms second between messages for each worker

## 🤖 Telegram Bot

### Features
- **User Registration**: `/start [channel_code]`
- **Channel Association**: Link users to specific channels
- **Error Handling**: Graceful handling of invalid codes
- **User Tracking**: Monitor user engagement and channel usage

### Bot Commands
```
/start [channel_code] - Register user and associate with channel
```

### How It Works
1. **User starts bot** with `/start` or `/start [code]` (Link format: `https://t.me/YourBot?start=eyJjaGFubmVsQ29kZSI6IkFCQzEyMyJ9`)
2. **Bot validates** channel code if provided
3. **User is saved** regardless of code validity
4. **Channel association** is created if code is valid
5. **Welcome message** is sent to user

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
│       └── main.go                    # Application entry point
├── config/
│   └── config.go                      # Configuration management
├── internal/
│   ├── api/
│   │   └── http/
│   │       ├── controllers/           # HTTP controllers with DTOs
│   │       │   ├── user/             # User controller + DTOs
│   │       │   ├── channel/          # Channel controller + DTOs
│   │       │   ├── notification/     # Notification controller + DTOs
│   │       │   └── common/           # Common response types
│   │       ├── middleware/           # HTTP middleware (auth)
│   │       └── routing/              # Route definitions
│   ├── app/
│   │   ├── app.go                    # Application setup
│   │   └── logger.go                 # Logging configuration
│   ├── domain/                       # Domain models (business entities)
│   │   ├── user.go                   # User domain model
│   │   ├── channel.go                # Channel domain model
│   │   ├── notification.go           # Notification data
│   ├── infrastructure/
│   │   └── database.go               # Database connection
│   ├── repository/                   # Data access layer
│   │   ├── user_postgres.go          # User repository
│   │   ├── channel_postgres.go       # Channel repository
│   └── service/                      # Business logic layer
│       ├── user_service.go           # User business logic
│       ├── channel_service.go        # Channel business logic
│       ├── telegram_service.go       # Telegram integration logic
│       └── notification_service.go   # Notification logic
├── Dockerfile                        # Docker configuration
├── go.mod                            # Go modules
├── go.sum                            # Go modules checksum
└── README.md                         # This file
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
