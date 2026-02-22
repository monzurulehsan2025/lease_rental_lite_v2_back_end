# Risk Assessment Microservice

A core backend service for modern rental qualification ecosystems, designed to qualify renters faster and mitigate risk for property operators using automated assessment logic.

## Features

- **Renter Qualification**: Evaluate eligibility based on income, credit score, and rental requirements.
- **Risk Scoring**: Proprietary scoring algorithm (0-100) to categorize renter risk.
- **Qualification Recommendations**: Provides suggestions on approval, co-signer requirements, or rejection.
- **Health Monitoring**: Built-in health check endpoint for container orchestration.

## Tech Stack

- **Language**: Go 1.21+
- **Architecture**: Single-responsibility microservice following clean architecture principles.
- **API**: RESTful JSON API using standard library `net/http`.

## Getting Started

### Prerequisites

- Go installed on your machine.
- A tool like `curl` or `Postman` for testing.

### Running Locally

1. Clone the repository and navigate to the project directory.
2. Initialize and download dependencies:
   ```bash
   go mod tidy
   ```
3. Start the server:
   ```bash
   go run cmd/server/main.go
   ```
   The service will start on `http://localhost:8080`.

### API Endpoints

#### 1. Assess Renter Risk
- **URL**: `/api/v1/assess`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "id": "renter_123",
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "annual_income": 85000,
    "credit_score": 720,
    "monthly_rent": 2500,
    "employment_status": "Employed"
  }
  ```
- **Response**:
  ```json
  {
    "renter_id": "renter_123",
    "risk_score": 25,
    "qualified": true,
    "max_coverage": 30000,
    "recommendation": "Approve",
    "evaluated_at": "2026-02-22T12:00:00Z"
  }
  ```

#### 2. Health Check
- **URL**: `/health`
- **Method**: `GET`
- **Response**: `{"status":"up"}`

## Design Decisions

- **Clean Architecture**: Separation of concerns between `handler` (external interface), `service` (business logic), and `model` (domain entities).
- **Standard Library**: Used `net/http` to demonstrate proficiency with Go's powerful standard library, keeping the binary small and dependencies minimal.
- **Extensibility**: The `RiskService` interface allows for easy swapping with more advanced (e.g., ML-based) scoring engines in the future without changing the API layer.
