package handlers

import (
	"bytes"
	"cookdroogers/app"
	"cookdroogers/models"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

var Users map[string]*models.User

// DecodeBasicAuthHeader декодирует базовый заголовок авторизации и возвращает имя пользователя и пароль.
func DecodeBasicAuthHeader(header string) (username, password string, err error) {
	// Проверяем, что заголовок авторизации имеет правильный формат.
	if !strings.HasPrefix(header, "Basic ") {
		return "", "", fmt.Errorf("неверный формат заголовка авторизации: %s", header)
	}

	// Декодируем закодированные в base64 учетные данные из заголовка.
	encodedCredentials := header[6:]
	decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		return "", "", fmt.Errorf("не удалось декодировать учетные данные: %w", err)
	}

	// Разделяем учетные данные на имя пользователя и пароль.
	credentials := bytes.SplitN(decodedCredentials, []byte(":"), 2)
	if len(credentials) != 2 {
		return "", "", fmt.Errorf("неверный формат учетных данных: %s", decodedCredentials)
	}

	return string(credentials[0]), string(credentials[1]), nil
}

func LoginManager(authHeader string, app *app.App) (*models.Manager, error) {
	username, _, err := DecodeBasicAuthHeader(authHeader)
	if err != nil {
		return nil, errors.New("can't authorize")
	}

	user, err := app.Services.UserService.GetByEmail(username)
	if err != nil {
		return nil, errors.New("can't authorize")
	}

	mngr, err := app.Services.ManagerService.GetByUserID(user.UserID)
	if err != nil {
		return nil, errors.New("can't find manager")
	}

	err = app.Services.UserService.SetRole(models.ManagerUser)
	if err != nil {
		return nil, errors.New("can't set manager role")
	}

	return mngr, nil
}

func LoginArtist(authHeader string, app *app.App) (*models.Artist, error) {
	username, _, err := DecodeBasicAuthHeader(authHeader)
	if err != nil {
		return nil, errors.New("can't authorize")
	}

	user, err := app.Services.UserService.GetByEmail(username)
	if err != nil {
		return nil, errors.New("can't authorize")
	}

	artist, err := app.Services.ArtistService.GetByUserID(user.UserID)
	if err != nil {
		return nil, errors.New("can't find artist")
	}

	err = app.Services.UserService.SetRole(models.ArtistUser)
	if err != nil {
		return nil, errors.New("can't set artist role")
	}

	return artist, nil
}

func LoginNonMember(authHeader string, app *app.App) (*models.User, error) {
	username, _, err := DecodeBasicAuthHeader(authHeader)
	if err != nil {
		return nil, errors.New("can't authorize")
	}

	user, err := app.Services.UserService.GetByEmail(username)
	if err != nil {
		return nil, errors.New("can't authorize")
	}

	err = app.Services.UserService.SetRole(models.NonMemberUser)
	if err != nil {
		return nil, errors.New("can't set nonmember role")
	}

	return user, nil
}

func LoginAdmin(authHeader string, app *app.App) error {
	username, password, err := DecodeBasicAuthHeader(authHeader)
	if err != nil {
		return errors.New("can't authorize")
	}

	if app.Config.Root.Username != username || app.Config.Root.Password != password {
		return errors.New("invalid root username or password")
	}

	err = app.Services.UserService.SetRole(models.AdminUser)
	if err != nil {
		return errors.New("can't set admin role")
	}

	return nil
}
