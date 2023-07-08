package port

import "github.com/bugrakocabay/airline/auth-service/data"

type IAuthRepository interface {
	GetAll() ([]*data.User, error)
	GetByEmail(email string) (*data.User, error)
	GetOne(id int) (*data.User, error)
	Update() error
	Delete() error
	DeleteByID(id int) error
	Insert(email, password, firstname, lastname string) (int, error)
	ResetPassword(password string) error
	PasswordMatches(plainText string) (bool, error)
}
