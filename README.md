# assesment-deals

This is a backend service for a dating mobile app written in Golang.

## Features

- User Sign Up
- User Login
- JWT Authentication

## Project Structure

```plaintext
.
├── README.md
├── go.mod
├── go.sum
├── main.go
├── .env.example
├── config
│   └── config.go
├── handlers
│   └── auth.go
├── models
│   └── user.go
├── routes
│   └── routes.go
└── tests
    └── auth_test.go
```

## Getting Started

Prerequisites :

- Go 1.16+
- PostgreSQL

Installation :

- clone the repository: `git clone https://github.com/mfauzzi/assesment-deals.git`
- cd assesment-deals
- copy the .env.example file to .env and fill in your database credentials
- install dependencies: `go mod tidy`

Running the Application :

- Start PostgreSQL and create the database
- Run the database migrations: `psql -U your_db_user -d assesment_deals -a -f db/migration/create_users_table.up.sql`
- Run the application: `go run main.go`

Deployment :

- Docker: Provide a Dockerfile and docker-compose.yml for easy setup and deployment.
- Linting: Use golangci-lint for linting the code. Install and run: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest` and `golangci-lint run`