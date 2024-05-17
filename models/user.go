package models

import (
    "database/sql"
    "time"
)

type User struct {
    ID           int
    Username     string
    PasswordHash string
    Email        string
    IsPremium    bool
    CreatedAt    time.Time
}

func CreateUser(db *sql.DB, username, passwordHash, email string) error {
    query := `INSERT INTO users (username, password_hash, email) VALUES ($1, $2, $3)`
    _, err := db.Exec(query, username, passwordHash, email)
    return err
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
    user := &User{}
    query := `SELECT id, username, password_hash, email, is_premium, created_at FROM users WHERE email = $1`
    err := db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Email, &user.IsPremium, &user.CreatedAt)
    if err != nil {
        return nil, err
    }
    return user, nil
}