# CICD Runner Service

A Go-based microservice that manages CI/CD pipelines and handles GitHub webhook callbacks for automated pipeline execution.

## Overview

The CICD Runner service provides:
- Pipeline creation and management
- GitHub webhook integration
- Pipeline event tracking
- Real-time status updates to frontend applications

## API Endpoints

### 1. Create Pipeline
**POST** `/create_pipeline`

Creates a new CI/CD pipeline and sets up GitHub webhook integration.

**Request Body:**
```json
{
  "pipeline": {
    "name": "build-and-test",
    "description": "Automated build and test pipeline",
    "status": "running",
    "trigger_type": "commit",
    "branch_name": "main",
    "repository_path": "owner/repository",
    "labels": ["frontend", "testing"]
  },
  "access_token": "github_personal_access_token"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Pipeline created successfully",
  "data": {
    "pipeline_id": 1
  }
}
```

### 2. Get Pipelines
**GET** `/pipelines`

Retrieves all pipelines with their associated events.

**Response:**
```json
{
  "success": true,
  "message": "Pipelines fetched successfully",
  "data": {
    "pipelines": [
      {
        "id": 1,
        "name": "build-and-test",
        "description": "Automated build and test pipeline",
        "status": "running",
        "last_run": "2024-01-15T10:30:00Z",
        "trigger_type": "commit",
        "branch_name": "main",
        "repository_path": "owner/repository",
        "labels": ["frontend", "testing"],
        "events": [
          {
            "id": 1,
            "pipeline_id": 1,
            "status": "success",
            "timestamp": "2024-01-15T10:30:00Z",
            "duration": 154000000000,
            "details": "Unit tests completed successfully",
            "type": "test"
          }
        ]
      }
    ]
  }
}
```

### 3. GitHub Webhook Callback
**POST** `/hook_callback`

Handles GitHub push events and creates pipeline events. This endpoint is called by GitHub when code is pushed to the repository.

**Request Body (from GitHub):**
```json
{
  "head_commit": {
    "id": "abc123",
    "message": "Fix bug in authentication",
    "timestamp": "2024-01-15T10:30:00Z"
  },
  "repository": {
    "full_name": "owner/repository"
  }
}
```

**Response:** No response body (HTTP 200)

## Environment Variables

| Variable | Description | Example Value |
|----------|-------------|---------------|
| `PORT` | Server port | `:8082` |
| `DB_URL` | PostgreSQL connection string | `postgres://user:pass@localhost:5432/dbname?sslmode=disable` |
| `CALLBACK_URL` | Public URL for GitHub webhook callbacks | `https://your-domain.com/hook_callback` |
| `SECRET` | Webhook secret for GitHub validation | `your-webhook-secret` |
| `FRONTEND_URL` | Frontend application URL for status updates | `http://localhost:3000` |

## Database Schema

### Pipelines Table
- `id` (Primary Key)
- `name` (String)
- `description` (String)
- `status` (String)
- `last_run` (Timestamp)
- `trigger_type` (String)
- `branch_name` (String)
- `repository_path` (String)
- `labels` (String Array)

### Pipeline Events Table
- `id` (Primary Key)
- `pipeline_id` (Foreign Key)
- `status` (String)
- `timestamp` (Timestamp)
- `duration` (Int64, nanoseconds)
- `details` (String)
- `type` (String)

### Access Tokens Table
- `id` (Primary Key)
- `access_token` (String)
- `pipeline_id` (Foreign Key)

## Features

- **Graceful Shutdown**: Handles SIGTERM/SIGINT signals for clean container shutdown
- **Database Integration**: PostgreSQL with GORM ORM
- **GitHub Integration**: Automatic webhook creation and validation
- **CORS Support**: Cross-origin requests enabled
- **Real-time Updates**: Pipeline events sent to frontend applications

## Running the Service

### Using Docker Compose
```bash
cd CICD_Runner
docker-compose up -d
```

### Manual Build
```bash
go mod download
go build -o main main.go
./main
```

## Dependencies

- **Gin**: HTTP web framework
- **GORM**: ORM for database operations
- **PostgreSQL Driver**: Database connectivity
