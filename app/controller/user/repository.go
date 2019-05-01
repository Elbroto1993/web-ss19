package user

import (
	"github.com/Elbroto1993/web-ss19/app/model"
)

// Repository user contract
type Repository interface {
	Fetch() ([]model.User, error)
	GetByID(id int64) (*model.User, error)
	GetByUsername(name string) (*model.User, error)
	GetByMail(mail string) (*model.User, error)
	Store(u *model.User) error
	Delete(id int64) error
	Update(u *model.User) error
}
