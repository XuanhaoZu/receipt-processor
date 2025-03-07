# Receipt Processor

A simple web service for processing receipts and calculating reward points.

## Requirements
- Go 1.20+
- Docker (optional)

## Installation and Usage

### Run Locally

```sh
go run main.go
```
### Run with Docker
```sh
docker build -t receipt-processor .
docker run -p 8080:8080 receipt-processor
```

## API Endpoints
### Process a Receipt
- **POST** `/receipts/process`  
- Request Body: JSON receipt
- Response: { "id": "receipt-uuid" }
### Get Points
- **GET** `/receipts/{id}/points`  
- Response: { "points": 32 }

## License
MIT
