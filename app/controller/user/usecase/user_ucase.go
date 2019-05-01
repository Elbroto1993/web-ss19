package usecase

import (
	"github.com/Elbroto1993/web-ss19/app/controller/user"
	"github.com/Elbroto1993/web-ss19/app/model"
	"time"
)

type userUsecase struct {
	userRepo user.Repository
}

// NewUserUsecase will create new an userUsecase object representation of user.Usecase interface
func NewUserUsecase(u user.Repository) user.Usecase {
	return &userUsecase{
		userRepo: u,
	}
}

// Fetch all users
func (u *userUsecase) Fetch() ([]model.User, error) {
	listUser, err := u.userRepo.Fetch()
	if err != nil {
		return nil, err
	}

	return listUser, nil
}

func (u *userUsecase) GetByID(id int64) (*model.User, error) {
	res, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByUsername user
func (u *userUsecase) GetByUsername(name string) (*model.User, error) {
	res, err := u.userRepo.GetByUsername(name)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByMail user
func (u *userUsecase) GetByMail(mail string) (*model.User, error) {
	res, err := u.userRepo.GetByMail(mail)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Store new user
func (u *userUsecase) Store(newUser *model.User) error {

	// @TODO - TESTFALL IST FALSCH PROGRAMMIERT
	existedUser, _ := u.GetByUsername(newUser.Username)
	if existedUser.Username != "" {
		return model.ErrUserameConflict
	}
	existedUser, _ = u.GetByMail(newUser.Email)
	if existedUser.Email != "" {
		return model.ErrEmailConflict
	}
	if newUser.Username == "" || newUser.Password == "" || newUser.Email == "" {
		return model.ErrBadParamInput
	}

	newUser.CreatedAt = time.Now()
	newUser.LoggedIn = false
	err := u.userRepo.Store(newUser)
	if err != nil {
		return err
	}
	return nil
}

// Delete user
func (u *userUsecase) Delete(id int64) error {
	existedUser, err := u.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existedUser == nil {
		return model.ErrNotFound
	}
	return u.userRepo.Delete(id)
}

// Update user
func (u *userUsecase) Update(updateUser *model.User) error {
	existedUser, _ := u.GetByUsername(updateUser.Username)
	existedUser, _ = u.GetByMail(updateUser.Email)
	if existedUser.Email != "" {
		return model.ErrEmailConflict
	}
	if updateUser.Password == "" || updateUser.Email == "" {
		return model.ErrBadParamInput
	}
	return u.userRepo.Update(updateUser)
}

// LOGIN
func (u *userUsecase) Login(loginUser *model.User) bool {
	checkUser, _ := u.GetByUsername(loginUser.Username)

	if checkUser.Password == loginUser.Password && loginUser.Password != "" {
		return true
	}
	return false
}

// HELPER FUNCTIONS
func (u *userUsecase) GetAllUsersLength() (int, error) {
	var length int
	listUsers, err := u.Fetch()
	if err != nil {
		return -1, err
	}
	length = len(listUsers)
	return length, nil
}
