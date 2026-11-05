# BirrWise Backend

Go/Gin REST API with SQLite storage.

## Run

```bash
cd backend
cp .env.example .env
# Environment variables are read directly by the process; export them or use:
PORT=8080 DATABASE_PATH=birrwise.db GIN_MODE=debug go run ./cmd/server
```

## API

- `GET /api/health`
- `/api/transactions` — list, create, update, and delete transactions
- `/api/categories` — list, create, update, and delete categories
- `/api/budgets` — list, create, update, and delete budgets
- `/api/preferences` — get and update user preferences

The SQLite schema is created and default categories seeded automatically. PostgreSQL integration can replace the repository/database layer while preserving the REST contract.
