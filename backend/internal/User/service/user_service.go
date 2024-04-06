package service

import (
	"context"
	"cookdroogers/internal/repo"
	userErrors "cookdroogers/internal/user/errors"
	"cookdroogers/models"
	repoErrors "cookdroogers/pkg/errors/repo"
	"errors"
	"fmt"
	"net/mail"
)

type IUserService interface {
	Create(*models.User) error
	Login(login, password string) (*models.User, error)
	GetByEmail(string) (*models.User, error)
	Get(uint64) (*models.User, error)
	Update(*models.User) error
	UpdateType(uint64, models.UserType) error
}

type UserService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) IUserService {
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

	_, err = usrSvc.repo.GetByEmail(context.Background(), newUser.Email)
	if err != nil && !errors.Is(err, repoErrors.ErrorNotExists) {
		return fmt.Errorf("can't create user: %w", err)
	}

	newUser.Type = models.NonMemberUser

	err = usrSvc.repo.Create(context.Background(), newUser)
	if err != nil {
		return fmt.Errorf("can'r create user: %w", err)
	}

	return nil
}

func (usrSvc *UserService) Login(login, password string) (*models.User, error) {
	user, err := usrSvc.repo.GetByEmail(context.Background(), login)
	if err != nil {
		return nil, err
	}

	if password != user.Password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (usrSvc *UserService) GetByEmail(email string) (*models.User, error) {
	user, err := usrSvc.repo.GetByEmail(context.Background(), email)
	if err != nil {
		return nil, fmt.Errorf("can't get user with err %w", err)
	}
	return user, nil
}

func (usrSvc *UserService) Get(id uint64) (*models.User, error) {
	user, err := usrSvc.repo.Get(context.Background(), id)

	if err != nil {
		return nil, fmt.Errorf("can't get user with err %w", err)
	}
	return user, nil
}

func (usrSvc *UserService) Update(user *models.User) error {
	if err := usrSvc.repo.Update(context.Background(), user); err != nil {
		return fmt.Errorf("can't update user with err %w", err)
	}
	return nil
}

func (usrSvc *UserService) UpdateType(userID uint64, typ models.UserType) error {
	if err := usrSvc.repo.UpdateType(context.Background(), userID, typ); err != nil {
		return fmt.Errorf("can't update user with err %w", err)
	}
	return nil
}
