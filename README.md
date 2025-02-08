# Real-time Video Ranking Service

A high-performance microservice for managing real-time video rankings based on user interactions.

## Features

- Real-time video ranking based on multiple metrics (views, likes, comments, shares, watch time)
- Redis-backed caching for fast ranking updates and queries
- PostgreSQL for persistent storage
- RESTful API with Swagger documentation
- Automated scoring system
- Support for global and per-user rankings

## Prerequisites

- Go 1.21+
- PostgreSQL 14+
- Redis 6+
- Docker (optional)

## Setup

1. Clone the repository:
```bash
git clone https://github.com/yourusername/ranking-service
cd ranking-service
```

2. Install dependencies:
```bash
go mod tidy
```

3. Configure environment variables:
```bash
cp dev.env .env
# Edit .env with your configuration
```

4. Start required services:
```bash
# Using Docker
docker run -d -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=ranking_service postgres:14
docker run -d -p 6379:6379 redis:6

# Or start PostgreSQL and Redis manually
```

5. Run the service:
```bash
go run cmd/server.go
```

## API Documentation

Access Swagger documentation at `http://localhost:8080/swagger/index.html`

### Key Endpoints

- `POST /api/v1/videos/:id/score` - Update video score
- `GET /api/v1/videos/top` - Get top-ranked videos
- `GET /api/v1/users/:id/videos/top` - Get user's top videos

## Testing

Run tests:
```bash
go test ./... -v
```

## Architecture

The service uses a layered architecture:
- Handlers: API endpoints and request handling
- Services: Business logic and score calculation
- DAOs: Data access and caching logic
- Models: Data structures and validation

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT License
