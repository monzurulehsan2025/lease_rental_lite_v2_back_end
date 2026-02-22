# Risk Assessment Microservice (Java)

Spring Boot implementation of the Risk Assessment microservice.

## Features
- **Spring Boot 3.2**: High-performance, production-ready Java framework.
- **REST API**: Standard endpoints for renter assessment.
- **Lombok**: Reduced boilerplate code for models.
- **JUnit 5**: Comprehensive test suite.

## Usage

### Build Project
```bash
mvn clean install
```

### Run Service
```bash
mvn spring-boot:run
```
The service will start on `http://localhost:8080`.

### Run Tests
```bash
mvn test
```

## API
- `POST /api/v1/assess`: Same request/response schema as the Go/Node versions.
- `GET /health`: Health status.
