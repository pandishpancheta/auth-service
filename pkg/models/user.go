package models

import (
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
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Instagram string    `json:"instagram"`
	Other     string    `json:"other"`
}
