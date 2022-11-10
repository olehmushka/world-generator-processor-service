package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID    `json:"id"`
	Username    string       `json:"username"`
	Language    string       `json:"lang"`
	Country     string       `json:"country"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Email       string       `json:"email"`
	Birthday    sql.NullTime `json:"birthday"`
	Password    string       `json:"password"`
	DynamicSalt string       `json:"dynamic_salt"`
	Status      string       `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	ModifiedAt  time.Time    `json:"modified_at"`
}
