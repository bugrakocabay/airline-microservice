package ports

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"-"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IAuthRepository interface {
	Insert(email, password, firstname, lastname string) (int, error)
	GetAll() ([]*User, error)
	GetByEmail(email string) (*User, error)
	GetOne(id int) (*User, error)
	Update(email, firstname, lastname string, active, id int) error
	DeleteByID(id int) error
	ResetPassword(password string, id int) error
	PasswordMatches(plainText, savedPassword string) (bool, error)
}
