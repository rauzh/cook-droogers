package techUI

import (
	"cookdroogers/app"
	"cookdroogers/models"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
)

func RunMenu(a *app.App, log *slog.Logger) error {
	startPosition :=
		`
		0 -- выйти
		1 -- авторизоваться
		2 -- зарегистрироваться
		3 -- посмотреть информацию о лейбле
	Выберите пункт меню: `
	fmt.Printf("%s", startPosition)

	var action int
	_, _ = fmt.Scanf("%d", &action)

	var err error

	switch action {
	case 0:
		err = ErrEXIT
	case 1:
		user, err := loginCLI(a)
		if err != nil {
			log.Error("Login error: ", slog.Any("error", err))
			break
		}

		err = switchMenu(a, user, log)

	case 2:
		user, err := registerCLI(a)
		if err != nil {
			log.Error("Register error: ", slog.Any("error", err))
			break
		}

		fmt.Println(`
	Переводим в меню пользователя..`)
		err = userLoop(a, user, log)
		if errors.Is(err, ErrEXIT) {
			err = nil
		}

	case 3:
		printInfo()

	default:
		fmt.Printf("Неверный пункт меню")
		err = ErrCase
	}

	return err
}

func printInfo() {
	file, _ := os.Open("label_info.txt")

	defer func() { _ = file.Close() }()

	b, _ := io.ReadAll(file)
	fmt.Print(b)
}

func switchMenu(a *app.App, user *models.User, log *slog.Logger) (err error) {
	switch user.Type {
	case models.ManagerUser:
		fmt.Println(`
	Переводим в меню менеджера...`)

	case models.ArtistUser:
		fmt.Println(`
	Переводим в меню артиста..`)
		menu, err := initArtistMenu(a, user, log)
		if err != nil {
			log.Error("Can't init artist menu: ", slog.Any("error", err))
		}

		err = menu.Loop()
		if errors.Is(err, ErrEXIT) {
			err = nil
		}

	case models.NonMemberUser:
		fmt.Println(`
	Переводим в меню пользователя..`)
		err = userLoop(a, user, log)
		if errors.Is(err, ErrEXIT) {
			err = nil
		}
	}

	return
}
