package usecase

import (
	"github.com/Elbroto1993/web-ss19/app/controller/karteikarte"
	"github.com/Elbroto1993/web-ss19/app/model"
)

type karteikarteUsecase struct {
	karteikarteRepo karteikarte.Repository
}

// NewKarteikarteUsecase will create new an userUsecase object representation of user.Usecase interface
func NewKarteikarteUsecase(k karteikarte.Repository) karteikarte.Usecase {
	return &karteikarteUsecase{
		karteikarteRepo: k,
	}
}

func (k *karteikarteUsecase) Fetch() ([]model.Karteikarte, error) {
	listKarte, err := k.karteikarteRepo.Fetch()
	if err != nil {
		return nil, err
	}

	return listKarte, nil
}

func (k *karteikarteUsecase) GetByID(id int64) (*model.Karteikarte, error) {
	res, err := k.karteikarteRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (k *karteikarteUsecase) GetByKastenID(id int64) ([]model.Karteikarte, error) {
	res, err := k.karteikarteRepo.GetByKastenID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (k *karteikarteUsecase) Store(u *model.Karteikarte) error {
	err := k.karteikarteRepo.Store(u)
	if err != nil {
		return err
	}
	return nil
}

func (k *karteikarteUsecase) Delete(id int64) error {
	existedKasten, err := k.karteikarteRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existedKasten == nil {
		return model.ErrNotFound
	}
	return k.karteikarteRepo.Delete(id)
}

func (k *karteikarteUsecase) Update(u *model.Karteikarte) error {
	return k.karteikarteRepo.Update(u)
}
