# Ocra Backend

Small project I'm doing to learn Go.

## About

Ocra is a real estate automation platform designed to streamline the workflow of real estate agents by matching:

- **Sellers** - Clients looking to sell their properties
- **Buyers** - Clients searching for their ideal home
- **Agent Portfolios** - Personal property listings managed by the agent

## Tech Stack

- **Go**
- **Fiber**
- **GORM**
- **PostgreSQL** (Supabase)

## Project Structure

```text
ocra-be/
├── api/
│   ├── handlers/
│   │   └── *_handler.go
│   ├── presenter/
│   │   └── *.go
│   └── routes/
│       └── *.go
├── cmd/
│   └── main.go
├── database/
│   └── database.go
├── pkg/
│   ├── entities/
│   │   └── *.go
│   └── */
│       ├── repository.go
│       └── service.go
├── go.mod
├── go.sum
├── .gitignore
└── README.md
```
