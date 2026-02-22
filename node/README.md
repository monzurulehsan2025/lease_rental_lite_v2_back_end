# Risk Assessment Microservice (Node.js)

Mirror implementation of the Go-based Risk Assessment microservice.

## Features
- **Express-based API**: Lightweight and fast RESTful endpoints.
- **ES Modules**: Modern JavaScript syntax.
- **Jest Testing**: Robust unit testing for risk calculation logic.

## Usage

### Install Dependencies
```bash
npm install
```

### Run Service
```bash
npm start
```
The service will start on `http://localhost:8080`.

### Run Tests
```bash
npm test
```

## API
- `POST /api/v1/assess`: Same request/response schema as the Go version.
- `GET /health`: Health status.
