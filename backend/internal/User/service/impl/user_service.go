package service

import (
	"cookdroogers/internal/User/repo"
	"cookdroogers/models"
	repo_errors "cookdroogers/pkg/errors/repo"
	"errors"
	"fmt"
)

type UserService struct {
	repo repo.UserRepo
}

func (us *UserService) Create(newUser *models.User) error {

	if newUser.Email == "" {
		return errors.New("can't  create user: no email provided")
	}
	if newUser.Name == "" {
		return errors.New("can't  create user: no name provided")
	}
	if newUser.Password == "" {
		return errors.New("can't  create user: no password provided")
	}

	_, err := us.repo.GetByEmail(newUser.Email)
	if err != nil && !errors.Is(err, repo_errors.ErrorNotExists) {
		return fmt.Errorf("can't create user: %w", err)
	}

	newUser.Type = models.NonMemberUser

	err = us.repo.Create(newUser)
	if err != nil {
		return fmt.Errorf("can'r create user: %w", err)
	}

	return nil
}

func (us *UserService) Login(login, password string) (*models.User, error) {
	user, err := us.repo.GetByEmail(login)
	if err != nil {
		return nil, err
	}

	if password != user.Password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (us *UserService) GetByEmail(email string) (*models.User, error) {
	user, err := us.repo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("can't get user with err %w", err)
	}
	return user, nil
}

func (us *UserService) Get(id uint64) (*models.User, error) {
	user, err := us.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("can't get user with err %w", err)
	}
	return user, nil
}

func (us *UserService) Update(user *models.User) error {
	if err := us.repo.Update(user); err != nil {
		return fmt.Errorf("can't update user with err %w", err)
	}
	return nil
}
