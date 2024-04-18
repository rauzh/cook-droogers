package techUI

import (
	"cookdroogers/app"
	"cookdroogers/models"
	"fmt"
)

func loginCLI(a *app.App) (*models.User, error) {

	fmt.Print(`
	Введите ваш логин (email): `)

	var email string
	_, _ = fmt.Scanf("%s", &email)

	fmt.Print(`
	Введите ваш пароль (password): `)

	var password string
	_, _ = fmt.Scanf("%s", &password)

	return a.Services.UserService.Login(email, password)
}

func registerCLI(a *app.App) (*models.User, error) {

	fmt.Print(`
	Введите ваше имя: `)

	var name string
	_, _ = fmt.Scanf("%s", &name)

	fmt.Print(`
	Введите ваш логин (email): `)

	var email string
	_, _ = fmt.Scanf("%s", &email)

	fmt.Print(`
	Введите ваш пароль (password): `)

	var password string
	_, _ = fmt.Scanf("%s", &password)

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := a.Services.UserService.Create(user)

	return user, err
}
