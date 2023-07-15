package domains

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

type CreateUserRequest struct {
	Username string
	Email    string
}

type UpdateUserRequest struct {
	Username string
	Email    string
}
