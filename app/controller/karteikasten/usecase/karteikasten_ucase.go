package usecase

import (
	"github.com/Elbroto1993/web-ss19/app/controller/karteikasten"
	"github.com/Elbroto1993/web-ss19/app/model"
)

type karteikastenUsecase struct {
	karteikastenRepo karteikasten.Repository
}

// NewKarteikastenUsecase will create new an userUsecase object representation of user.Usecase interface
func NewKarteikastenUsecase(k karteikasten.Repository) karteikasten.Usecase {
	return &karteikastenUsecase{
		karteikastenRepo: k,
	}
}

func (k *karteikastenUsecase) Fetch() ([]model.Karteikasten, error) {
	listKasten, err := k.karteikastenRepo.Fetch()
	if err != nil {
		return nil, err
	}

	return listKasten, nil
}

func (k *karteikastenUsecase) GetByID(id int64) (*model.Karteikasten, error) {
	res, err := k.karteikastenRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (k *karteikastenUsecase) GetByUserID(id int64) ([]model.Karteikasten, error) {
	listKasten, err := k.karteikastenRepo.GetByUserID(id)
	if err != nil {
		return nil, err
	}
	return listKasten, nil
}

func (k *karteikastenUsecase) GetByCreatedUserID(id int64) ([]model.Karteikasten, error) {
	res, err := k.karteikastenRepo.Fetch()
	var listKasten []model.Karteikasten
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(res); i++ {
		if res[i].CreatedByUserID == id {
			listKasten = append(listKasten, res[i])
		}
	}
	return listKasten, nil
}

func (k *karteikastenUsecase) Store(u *model.Karteikasten) error {
	err := k.karteikastenRepo.Store(u)
	if err != nil {
		return err
	}
	return nil
}

func (k *karteikastenUsecase) Delete(id int64) error {
	existedKasten, err := k.karteikastenRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existedKasten == nil {
		return model.ErrNotFound
	}
	return k.karteikastenRepo.Delete(id)
}

func (k *karteikastenUsecase) Update(u *model.Karteikasten) error {
	return k.karteikastenRepo.Update(u)
}
