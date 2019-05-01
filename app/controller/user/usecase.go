package user

import (
	"github.com/Elbroto1993/web-ss19/app/model"
)

// Usecase user
type Usecase interface {
	Fetch() ([]model.User, error)
	GetByID(id int64) (*model.User, error)
	GetByUsername(name string) (*model.User, error)
	GetByMail(mail string) (*model.User, error)
	Store(newUser *model.User) error
	Delete(id int64) error
	Update(updateUser *model.User) error
	Login(loginUser *model.User) bool
	GetAllUsersLength() (int, error)
}
