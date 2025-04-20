package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"_id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
}
