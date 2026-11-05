# BirrWise — Personal Expense Tracker

BirrWise is a full-stack personal finance application for Ethiopian Birr users.

## Project structure

```text
ethiopian-expense-tracker/
├── frontend/           # React, JavaScript, Tailwind CSS
│   ├── public/
│   ├── src/
│   └── package.json
├── backend/            # Go, Gin, SQLite
│   ├── cmd/server/
│   ├── internal/
│   └── go.mod
└── package.json        # Workspace convenience commands
```

## Run locally

Terminal 1 — backend:

```bash
npm run backend:dev
```

Terminal 2 — frontend:

```bash
npm run frontend:dev
```

Frontend: `http://localhost:5173`  
Backend API: `http://localhost:8080/api`

## Build

```bash
npm run frontend:build
npm run backend:build
```
