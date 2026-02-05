# RSS Aggregator

A web scraper and RSS feed aggregator built with Go.

## Tech Stack

- **Go** with [Gin](https://github.com/gin-gonic/gin) web framework
- **PostgreSQL** database
- **sqlc** for type-safe SQL queries
- **goose** for database migrations
- **Docker Compose** for local development

## Getting Started

### Prerequisites

- Go 1.25+
- Docker & Docker Compose
- [goose](https://github.com/pressly/goose) - `go install github.com/pressly/goose/v3/cmd/goose@latest`
- [sqlc](https://github.com/sqlc-dev/sqlc) - `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`

### Setup

1. Clone the repo:
   ```bash
   git clone https://github.com/KaiyueWei/rssagg.git
   cd rssagg
   ```

2. Start PostgreSQL:
   ```bash
   docker compose up -d
   ```

3. Create a `.env` file:
   ```
   PORT=8000
   DB_URL=postgres://admin:password@localhost:5432/scraper?sslmode=disable
   ```

4. Run database migrations:
   ```bash
   goose -dir sql/schema postgres "postgres://admin:password@localhost:5432/scraper?sslmode=disable" up
   ```

5. Build and run:
   ```bash
   go build && ./rssagg
   ```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/v1/healthz` | Health check |
| GET | `/v1/err` | Error test |
| POST | `/v1/user` | Create a user |

### Create a User

```bash
curl -X POST http://localhost:8000/v1/user \
  -H "Content-Type: application/json" \
  -d '{"name": "John"}'
```
