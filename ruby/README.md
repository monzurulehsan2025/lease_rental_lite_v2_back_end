# Risk Assessment Microservice (Ruby)

Sinatra implementation of the Risk Assessment microservice.

## Features
- **Sinatra**: DSL for quickly creating web applications in Ruby with minimal effort.
- **RSpec**: Behaviour Driven Development for Ruby.
- **Clean Logic**: Shared business logic with Go, Node, Java, and Python versions.

## Usage

### Install Dependencies
```bash
bundle install
```

### Run Service
```bash
ruby app/api.rb
```
The service will start on `http://localhost:8080`.

### Run Tests
```bash
rspec
```

## API
- `POST /api/v1/assess`: Same request/response schema as all other versions.
- `GET /health`: Health status.
