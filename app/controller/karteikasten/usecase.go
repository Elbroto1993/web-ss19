package karteikasten

import (
	"github.com/Elbroto1993/web-ss19/app/model"
)

// Usecase karteikasten contract
type Usecase interface {
	Fetch() ([]model.Karteikasten, error)
	GetByID(id int64) (*model.Karteikasten, error)
	GetByUserID(id int64) ([]model.Karteikasten, error)
	GetByCreatedUserID(id int64) ([]model.Karteikasten, error)
	Store(u *model.Karteikasten) error
	Delete(id int64) error
	Update(u *model.Karteikasten) error
}
