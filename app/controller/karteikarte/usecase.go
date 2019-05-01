package karteikarte

import (
	"github.com/Elbroto1993/web-ss19/app/model"
)

// Usecase Karteikarte user contract
type Usecase interface {
	Fetch() ([]model.Karteikarte, error)
	GetByID(id int64) (*model.Karteikarte, error)
	GetByKastenID(id int64) ([]model.Karteikarte, error)
	Store(u *model.Karteikarte) error
	Delete(id int64) error
	Update(u *model.Karteikarte) error
}
