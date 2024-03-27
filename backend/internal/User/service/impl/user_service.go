package service

import (
	userErrors "cookdroogers/internal/User/errors"
	"cookdroogers/internal/User/repo"
	userService "cookdroogers/internal/User/service"
	"cookdroogers/models"
	repoErrors "cookdroogers/pkg/errors/repo"
	"errors"
	"fmt"
	"net/mail"
)

type UserService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) userService.IUserService {
	return &UserService{repo: repo}
}

func (usrSvc *UserService) validate(usr *models.User) error {

	_, err := mail.ParseAddress(usr.Email)
	if err != nil {
		return userErrors.ErrInvalidEmail
	}

	if usr.Name == "" {
		return userErrors.ErrInvalidName
	}

	if usr.Password == "" {
		return userErrors.ErrInvalidPassword
	}

	return nil
}

func (usrSvc *UserService) Create(newUser *models.User) error {

	err := usrSvc.validate(newUser)
	if err != nil {
		return err
	}

	_, err = usrSvc.repo.GetByEmail(newUser.Email)
	if err != nil && !errors.Is(err, repoErrors.ErrorNotExists) {
		return fmt.Errorf("can't create user: %w", err)
	}

	newUser.Type = models.NonMemberUser

	err = usrSvc.repo.Create(newUser)
	if err != nil {
		return fmt.Errorf("can'r create user: %w", err)
	}

	return nil
}

func (usrSvc *UserService) Login(login, password string) (*models.User, error) {
	user, err := usrSvc.repo.GetByEmail(login)
	if err != nil {
		return nil, err
	}

	if password != user.Password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (usrSvc *UserService) GetByEmail(email string) (*models.User, error) {
	user, err := usrSvc.repo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("can't get user with err %w", err)
	}
	return user, nil
}

func (usrSvc *UserService) Get(id uint64) (*models.User, error) {
	user, err := usrSvc.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("can't get user with err %w", err)
	}
	return user, nil
}

func (usrSvc *UserService) Update(user *models.User) error {
	if err := usrSvc.repo.Update(user); err != nil {
		return fmt.Errorf("can't update user with err %w", err)
	}
	return nil
}

func (usrSvc *UserService) UpdateType(userID uint64, typ models.UserType) error {
	if err := usrSvc.repo.UpdateType(userID, typ); err != nil {
		return fmt.Errorf("can't update user with err %w", err)
	}
	return nil
}
