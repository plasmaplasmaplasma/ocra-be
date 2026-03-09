# Ocra Backend

Small project I'm doing to learn Go.

The goal of Ocra is to help real estate agents automate matching between:
- clients who want to sell their house
- clients who want to buy a house
- personal portfolio of the agente

## Tech Stack

- Go
- Fiber
- GORM
- PostgreSQL (Supabase)

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
