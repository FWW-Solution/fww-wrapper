package entity

import (
	"database/sql"
	"time"
)

type Partner struct {
	ID        string       `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	ApiKey    string       `db:"api_key"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

// User represents a user entity in the database.
type User struct {
	ID        int64        `db:"id"`
	FullName  string       `db:"full_name"`
	Username  string       `db:"username"`
	Email     string       `db:"email"`
	Password  string       `db:"password"` // encrypted
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
