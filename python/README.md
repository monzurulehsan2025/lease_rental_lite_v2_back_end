# Risk Assessment Microservice (Python)

FastAPI implementation of the Risk Assessment microservice.

## Features
- **FastAPI**: Extremely fast and modern web framework.
- **Pydantic V2**: Robust data validation and settings management.
- **Asynchronous**: Built on ASGI for high concurrency.
- **pytest**: industry-standard testing framework.

## Usage

### Install Dependencies
```bash
pip install -r requirements.txt
```

### Run Service
```bash
uvicorn app.main:app --host 0.0.0.0 --port 8080
```
The service will start on `http://localhost:8080`.

### Run Tests
```bash
pytest
```

## API
- `POST /api/v1/assess`: Same request/response schema as the Go/Node/Java versions.
- `GET /health`: Health status.
