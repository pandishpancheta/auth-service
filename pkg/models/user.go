package models

import (
	"database/sql"
	"github.com/gofrs/uuid/v5"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id;"`
	Username  string    `json:"username;"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Contacts  Contacts  `json:"contacts"`
	CreatedAt time.Time `json:"created_at"`
}

type Contacts struct {
	ID        uuid.UUID      `json:"id"`
	Email     sql.NullString `json:"email"`
	Phone     sql.NullString `json:"phone"`
	Instagram sql.NullString `json:"instagram"`
	Other     sql.NullString `json:"other"`
}
