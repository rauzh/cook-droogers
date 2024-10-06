package service

import (
	"context"
	"cookdroogers/internal/repo"
	userErrors "cookdroogers/internal/user/errors"
	"cookdroogers/models"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/mail"
	"strings"
)

type IUserService interface {
	Create(context.Context, *models.User) error
	Login(ctx context.Context, login, password string) (*models.User, error)
	GetByEmail(context.Context, string) (*models.User, error)
	Get(context.Context, uint64) (*models.User, error)
	Update(context.Context, *models.User) error
	UpdateType(context.Context, uint64, models.UserType) error
	GetForAdmin(context.Context) ([]models.User, error)

	SetRole(ctx context.Context, role models.UserType) error
}

type UserService struct {
	repo repo.UserRepo

	logger *slog.Logger
}

func NewUserService(repo repo.UserRepo, logger *slog.Logger) IUserService {
	return &UserService{repo: repo, logger: logger}
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

func (usrSvc *UserService) Create(ctx context.Context, newUser *models.User) error {

	err := usrSvc.validate(newUser)
	if err != nil {
		return err
	}

	usr, err := usrSvc.repo.GetByEmail(ctx, newUser.Email)

	if err != nil && !strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
		usrSvc.logger.Error("USER SVC: Create", "error", err.Error())
		return fmt.Errorf("can't create user: %w", err)
	}
	if usr != nil {
		return userErrors.ErrAlreadyTaken
	}

	newUser.Type = models.NonMemberUser

	err = usrSvc.repo.Create(ctx, newUser)
	if err != nil {
		usrSvc.logger.Error("USER SVC: Create", "error", err.Error())
		return fmt.Errorf("can'r create user: %w", err)
	}

	return nil
}


func (usrSvc *UserService) Login(ctx context.Context, login, password string) (*models.User, error) {
	user, err := usrSvc.repo.GetByEmail(ctx, login)
	if err != nil {

		return nil, err
	}

	if password != user.Password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (usrSvc *UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := usrSvc.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("can't get user with err %w", err)
	}
	return user, nil
}

func (usrSvc *UserService) Get(ctx context.Context, id uint64) (*models.User, error) {
	user, err := usrSvc.repo.Get(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("can't get user with err %w", err)
	}
	return user, nil
}

func (usrSvc *UserService) Update(ctx context.Context, user *models.User) error {
	if err := usrSvc.repo.Update(ctx, user); err != nil {
		return fmt.Errorf("can't update user with err %w", err)
	}
	return nil
}

func (usrSvc *UserService) UpdateType(ctx context.Context, userID uint64, typ models.UserType) error {
	if err := usrSvc.repo.UpdateType(ctx, userID, typ); err != nil {
		return fmt.Errorf("can't update user with err %w", err)
	}
	return nil
}

func (usrSvc *UserService) SetRole(ctx context.Context, role models.UserType) error {
	if err := usrSvc.repo.SetRole(ctx, role); err != nil {
		return fmt.Errorf("can't set user role with err %w", err)
	}
	return nil
}

func (usrSvc *UserService) GetForAdmin(ctx context.Context) ([]models.User, error) {
	users, err := usrSvc.repo.GetForAdmin(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't get users with err %w", err)
	}
	return users, nil
}
